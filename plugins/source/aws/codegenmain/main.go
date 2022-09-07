package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"log"
	"os"
	"path"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/helpers"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/recipes"
	sdkgen "github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"golang.org/x/exp/slices"
)

const useFullStruct = "." // special case to use full struct instead of a member

//go:embed templates/*.go.tpl
var awsTemplatesFS embed.FS

func main() {
	templatesWithMocks := map[string]bool{
		"resource_get":           true,
		"resource_list_describe": true,
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	dir := path.Dir(filename)

	resources := recipes.AllResources

	for _, r := range resources {
		initResource(r)
		if r.Parent != nil && r.TableFuncName != "" {
			r.Parent.Table.Relations = append(r.Parent.Table.Relations, r.TableFuncName+"(),\n")
		}
	}
	for _, r := range resources {
		if r.Parent != nil {
			handleParentReference(r)
		}

		if l := len(r.Table.Relations); l > 0 {
			r.Table.Relations[l-1] = strings.TrimSuffix(r.Table.Relations[l-1], ",\n")
		}

		generateResource(dir, r, false)

		if templatesWithMocks[r.Template] {
			generateResource(dir, r, true)
		}
	}

	generatePlugin(resources)
}

func inferFromRecipe(r *recipes.Resource) {
	var (
		res, items, pag, pget *helpers.InferResult
	)

	if r.ItemsStruct != nil {
		items = helpers.InferFromStructOutput(r.ItemsStruct)
		r.GetMethod = items.Method
		if helpers.BareType(reflect.TypeOf(r.AWSStruct)) == helpers.BareType(reflect.TypeOf(r.ItemsStruct)) {
			// Same type, so we need to save the output struct completely
			r.ResponseItemsName = useFullStruct
		} else {
			needSingular := r.PaginatorStruct != nil
			if !needSingular && len(items.ItemsFieldCandidates(false)) == 0 {
				// we have wrapper? look for a single item
				singleItemMaybeWrapper := items.ItemsFieldCandidates(true)
				if len(singleItemMaybeWrapper) != 1 {
					log.Fatal("Could not determine possible ItemsName wrapper for ", items.Method, ":", len(singleItemMaybeWrapper), " candidates")
				}
				r.ResponseItemsWrapper = singleItemMaybeWrapper[0].Name
				again := helpers.InferFromType(helpers.BareType(singleItemMaybeWrapper[0].Type), "")
				if again.PaginatorTokenField != nil {
					r.WrappedNextTokenName = again.PaginatorTokenField.Name
				}

			} else {
				r.ResponseItemsName = items.ItemsField(needSingular, helpers.BareType(reflect.TypeOf(r.AWSStruct)).Name()).Name
			}
		}

		if items.PaginatorTokenField != nil {
			r.NextTokenName = items.PaginatorTokenField.Name
		}

		res = items
	}

	var pagSingleItem reflect.Type // This is the single item from the paginator, which is used later to match fields.

	if r.PaginatorStruct != nil {
		pag = helpers.InferFromStructOutput(r.PaginatorStruct)
		var f reflect.StructField
		// we need a slice field, but what if it's wrapped?
		if len(pag.ItemsFieldCandidates(false)) == 0 {
			// we have wrapper? look for a single item
			singleItemMaybeWrapper := pag.ItemsFieldCandidates(true)
			if len(singleItemMaybeWrapper) != 1 {
				log.Fatal("Could not determine possible PaginatorStruct wrapper for ", pag.Method, ":", len(singleItemMaybeWrapper), " candidates")
			}
			r.PaginatorListWrapper = singleItemMaybeWrapper[0].Name

			if singleItemMaybeWrapper[0].Type.Kind() == reflect.Ptr {
				r.PaginatorListWrapperType = "&" + path.Base(singleItemMaybeWrapper[0].Type.Elem().PkgPath()) + "." + singleItemMaybeWrapper[0].Type.Elem().Name()
			} else {
				r.PaginatorListWrapperType = path.Base(singleItemMaybeWrapper[0].Type.PkgPath()) + "." + singleItemMaybeWrapper[0].Type.Name()
			}
			again := helpers.InferFromType(helpers.BareType(singleItemMaybeWrapper[0].Type), "")
			if again.PaginatorTokenField != nil {
				r.WrappedNextTokenName = again.PaginatorTokenField.Name
			}
			f = again.ItemsField(false, "")
		} else {
			f = pag.ItemsField(false, "")
		}

		pagSingleItem = f.Type.Elem() // get type of slice

		r.PaginatorListName = f.Name
		r.PaginatorListType = f.Type.Elem().Name() // single type from a slice
		if f.Type.Elem().Kind() == reflect.Struct {
			r.PaginatorListType = "types." + r.PaginatorListType
		}

		r.ListMethod = pag.Method

		if res == nil {
			res = pag
		}
	}

	if r.PaginatorGetStruct != nil {
		if r.ItemsStruct == nil {
			log.Fatal("PaginatorGetStruct requires ItemsStruct on resource ", r.AWSService)
		}

		pget = helpers.InferFromStructInput(r.PaginatorGetStruct)
		if pget.Method != items.Method {
			log.Fatal("PaginatorGetStruct method ", pget.Method, " does not match ItemsStruct method ", items.Method)
		}

		if pag != nil {
			// figure out which fields match to what

			r.AutoCalculated.GetAndListOrder = nil
			r.AutoCalculated.MatchedGetAndListFields = make(map[string]string)

			fields := make(map[string]reflect.Type)
			//log.Println("PROCESSING", pagSingleItem.Name(), pagSingleItem.Kind().String())
			if k := pagSingleItem.Kind(); k == reflect.String {
				// special case for string
				fields[""] = pagSingleItem
			} else {
				for i := 0; i < pagSingleItem.NumField(); i++ {
					f := pagSingleItem.Field(i)
					if f.Name == "noSmithyDocumentSerde" || f.Type.String() == "document.NoSerde" {
						continue
					}
					fields[f.Name] = f.Type
				}
			}

			if len(fields) == 1 && fields[""] != nil {
				// special case for string (not struct)
				found := false
				for _, f := range pget.FieldOrder {
					ff := pget.Fields[f].Type
					if helpers.BareType(ff).Kind() == helpers.BareType(fields[""]).Kind() {
						found = true
						r.AutoCalculated.GetAndListOrder = append(r.AutoCalculated.GetAndListOrder, f)
						r.AutoCalculated.MatchedGetAndListFields[f] = "&item"
						break
					}
				}
				if !found {
					log.Println("PaginatorGetStruct field of type", fields[""].Kind().String(), "not matched in PaginatorStruct in", pagSingleItem.Name())
				}
			} else {
				for _, f := range pget.FieldOrder {
					found := false
					nameMatchFn := func(a, b string) bool { return strings.ToLower(a) == strings.ToLower(b) }

					for attempts := 0; attempts < 2; attempts++ {
						for n, t := range fields {
							if nameMatchFn(n, f) && helpers.BareType(t) == helpers.BareType(pget.Fields[f].Type) {
								found = true
								r.AutoCalculated.GetAndListOrder = append(r.AutoCalculated.GetAndListOrder, f)
								r.AutoCalculated.MatchedGetAndListFields[f] = "item." + n
								break
							}
						}
						if found {
							break
						}
						if !found {
							if attempts == 0 {
								// Either suffix or single field
								nameMatchFn = func(a, b string) bool {
									return (len(pget.FieldOrder) == 1 && a == "Name") || strings.HasSuffix(a, b)
								}

								log.Println("PaginatorGetStruct field", f, "not matched in PaginatorStruct in", pagSingleItem.Name(), "doing heuristic match")
							} else {
								log.Println("PaginatorGetStruct field", f, "not matched in PaginatorStruct in", pagSingleItem.Name(), "even after heuristic match")
							}
						}
					}
				}
			}

			if len(r.AutoCalculated.GetAndListOrder) > 0 {
				log.Println("GetAndListOrder for", pagSingleItem.Name()+":", r.AutoCalculated.GetAndListOrder)
			}
		}

	}

	if items != nil && pag != nil {
		r.AWSSubService = pag.SubService
	} else {
		r.AWSSubService = res.SubService
	}

	if items != nil && pag != nil && strings.TrimSuffix(pag.SubService, "s") != strings.TrimSuffix(items.SubService, "s") { // Certificate vs Certificates
		log.Println("Mismatching subservices between ItemsStruct and PaginatorStruct for resource ", r.AWSService, ": ", items.SubService, " vs ", pag.SubService)
	}

}

func initResource(r *recipes.Resource) {
	if r.ItemsStruct != nil {
		inferFromRecipe(r)
	}

	tableNameFromSubService, fetcherNameFromSubService := helpers.TableAndFetcherNames(r)

	var err error
	r.Table, err = sdkgen.NewTableFromStruct(
		fmt.Sprintf("aws_%s_%s", strings.ToLower(r.AWSService), tableNameFromSubService),
		r.AWSStruct,
		sdkgen.WithSkipFields(append(r.SkipFields, "noSmithyDocumentSerde")),
	)
	if err != nil {
		log.Fatal(err)
	}
	r.Table.Resolver = helpers.Coalesce(r.RawResolver, "fetch"+r.AWSService+fetcherNameFromSubService)
	r.TableFuncName = r.AWSService + fetcherNameFromSubService

	if r.TrimPrefix != "" {
		for i := range r.Table.Columns {
			r.Table.Columns[i].Name = strings.TrimPrefix(r.Table.Columns[i].Name, r.TrimPrefix)
		}
	}

	if r.Parent != nil && len(r.DefaultColumns) > 0 {
		for _, c := range r.DefaultColumns {
			if c == recipes.AccountIdColumn || c == recipes.RegionColumn {
				log.Fatal("Error: A sub-resource of ", r.Parent.AWSStructName, " should not have account_id or region columns")
			}
		}
	}

	r.Table.Columns = append(r.DefaultColumns, r.Table.Columns...)
	if r.ColumnOverrides != nil {
		for i, c := range r.Table.Columns {
			override, ok := r.ColumnOverrides[c.Name]
			if !ok {
				continue
			}
			r.Table.Columns[i].Name = helpers.Coalesce(override.Name, r.Table.Columns[i].Name)
			r.Table.Columns[i].Resolver = helpers.Coalesce(override.Resolver, r.Table.Columns[i].Resolver)
			r.Table.Columns[i].Description = helpers.Coalesce(override.Description, r.Table.Columns[i].Description)

			delete(r.ColumnOverrides, c.Name)
		}
		coSlice := make([]string, 0, len(r.ColumnOverrides))
		for k := range r.ColumnOverrides {
			coSlice = append(coSlice, k)
		}
		sort.Strings(coSlice)
		// remaining, unmatched columns are added to the end of the table. Difference from DefaultColumns? none for now
		for _, k := range coSlice {
			c := r.ColumnOverrides[k]
			if c.Type == schema.TypeInvalid {
				fmt.Println("Not adding unmatched column with unspecified type", k, c)
				continue
			}
			c.Name = helpers.Coalesce(c.Name, k)
			r.Table.Columns = append(r.Table.Columns, c)
		}
	}

	if len(r.PrimaryKeys) > 0 {
		// Move PKs in the specified order just after the first PK.
		// This way we preserve DefaultColumns (account_id etc.) at the start of the table, even with arn as PK.
		newCols := make(sdkgen.ColumnDefinitions, 0, len(r.Table.Columns))
		firstPKIndex := -1
		for _, k := range r.PrimaryKeys {
			for i, c := range r.Table.Columns {
				if c.Name == k {
					if firstPKIndex == -1 {
						firstPKIndex = i
					}
					r.Table.Columns[i].Options = schema.ColumnCreationOptions{PrimaryKey: true}
					newCols = append(newCols, r.Table.Columns[i])
				}
			}
		}
		if firstPKIndex > -1 {
			newCols = append(r.Table.Columns[:firstPKIndex], newCols...) // all data before the first PK can be safely copied
		}
		for i, c := range r.Table.Columns {
			if c.Options.PrimaryKey || i < firstPKIndex {
				continue
			}
			newCols = append(newCols, r.Table.Columns[i]) // add remaining columns
		}
		r.Table.Columns = newCols
	}

	hasReferenceToResolvers := false

	for i := range r.Table.Columns {
		if r.Table.Columns[i].Name == "tags" {
			r.HasTags = true

			if r.Table.Columns[i].Resolver == recipes.ResolverAuto {
				r.Table.Columns[i].Resolver = "resolve" + r.AWSService + r.AWSSubService + "Tags"
			}
		}
		if strings.HasPrefix(r.Table.Columns[i].Resolver, "resolvers.") {
			hasReferenceToResolvers = true
		}
	}

	if r.RawMultiplexerOverride != "" && r.MultiplexerServiceOverride != "" {
		log.Fatal("Cannot specify both RawMultiplexerOverride and MultiplexerServiceOverride")
	}

	if r.RawMultiplexerOverride != "" {
		r.Table.Multiplex = r.RawMultiplexerOverride
	} else {
		r.Table.Multiplex = `client.ServiceAccountRegionMultiplexer("` + helpers.Coalesce(r.MultiplexerServiceOverride, strings.ToLower(r.AWSService)) + `")`
	}

	if strings.HasPrefix(r.RawResolver, "resolvers.") {
		hasReferenceToResolvers = true
	}

	r.MockFuncName = "build" + r.TableFuncName
	r.TestFuncName = "Test" + r.TableFuncName

	t := reflect.TypeOf(r.AWSStruct).Elem()
	r.AWSStructName = path.Base(t.PkgPath()) + "." + t.Name() // types.Something or sometimes service.Something
	r.ItemName = helpers.Coalesce(r.ItemName, t.Name())

	r.Imports = quoteImports(r.Imports)
	r.MockImports = quoteImports(r.MockImports)

	sp := t.PkgPath()
	var (
		mainImport string
	)
	if strings.HasSuffix(sp, "/types") {
		r.TypesImport = strconv.Quote(sp)
		mainImport = strings.TrimSuffix(sp, "/types")
	} else if strings.HasSuffix(sp, "/aws-sdk-go-v2/service/"+strings.ToLower(r.AWSService)) { // main struct lives in main pkg
		mainImport = sp
	}

	if r.RawResolver == "" {
		// auto import main pkg
		r.Imports = append(r.Imports, strconv.Quote(mainImport))
		r.MockImports = append(r.MockImports, strconv.Quote(mainImport))
	}

	if hasReferenceToResolvers {
		res := "resolvers " + strconv.Quote(`github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/resolvers/`+strings.ToLower(r.AWSService))
		if !slices.Contains(r.Imports, res) {
			r.Imports = append(r.Imports, res)
		}
	}
}

func handleParentReference(r *recipes.Resource) {
	// Add one level up parent key
	var pItemName string
	if r.Parent.CQSubserviceOverride != "" {
		pItemName = r.Parent.CQSubserviceOverride
	} else {
		pItemName = strings.TrimSuffix(r.Parent.ItemName, "Summary")
	}
	pItemName = inflection.Singular(pItemName)

	r.Table.Columns = append([]sdkgen.ColumnDefinition{
		{
			Name:     strings.ToLower(pItemName + "_cq_id"),
			Type:     schema.TypeUUID,
			Resolver: "schema.ParentIdResolver",
		},
	}, r.Table.Columns...)
}

func generateResource(dir string, r *recipes.Resource, mock bool) {
	r.TemplateFilename = r.Template + helpers.StringSwitch(mock, "_mock_test", "") + ".go.tpl"
	tpl, err := template.New(r.TemplateFilename).Funcs(template.FuncMap{
		"ToCamel":  strcase.ToCamel,
		"ToLower":  strings.ToLower,
		"ToSnake":  strcase.ToSnake,
		"Coalesce": func(a1, a2 string) string { return helpers.Coalesce(a2, a1) }, // go templates argument order is backwards
	}).ParseFS(awsTemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse aws templates: %w", err))
	}
	tpl, err = tpl.ParseFS(sdkgen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse codegen template: %w", err))
	}

	var buff bytes.Buffer

	// we try two times to try and detect if we're using any `types.` references
	for i := 0; i < 2; i++ {
		if err := tpl.Execute(&buff, r); err != nil {
			log.Fatal(fmt.Errorf("failed to execute template: %w", err))
		}
		if i == 1 || r.SkipTypesImport || r.TypesImport == "" || !strings.Contains(buff.String(), "types.") {
			break
		}

		if !mock {
			r.Imports = append(r.Imports, r.TypesImport)
		} else {
			r.MockImports = append(r.MockImports, r.TypesImport)
		}
		buff.Reset()
	}

	filePath := path.Join(dir, "../codegen", strings.ToLower(r.AWSService))
	if err := os.MkdirAll(filePath, 0755); err != nil {
		log.Fatal(fmt.Errorf("failed to create directory: %w", err))
	}

	fileSuffix := helpers.StringSwitch(mock, "_mock_test.go", ".go")
	filePath = path.Join(filePath, strings.TrimPrefix(r.Table.Name, "aws_")+fileSuffix)
	content, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Println(buff.String())
		log.Fatal(fmt.Errorf("failed to format code for %s: %w", filePath, err))
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}

func generatePlugin(rr []*recipes.Resource) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	dir := path.Dir(filename)
	tpl, err := template.New("tables.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
		"ToSnake": strcase.ToSnake,
	}).ParseFS(awsTemplatesFS, "templates/tables.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse tables.go.tpl: %w", err))
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, rr); err != nil {
		log.Fatal(fmt.Errorf("failed to execute template: %w", err))
	}

	filePath := path.Join(dir, "../plugin/autogen_tables.go")
	content, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Println(buff.String())
		log.Fatal(fmt.Errorf("failed to format code for %s: %w", filePath, err))
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}

func quoteImports(imports []string) []string {
	for i := range imports {
		if !strings.HasSuffix(imports[i], `"`) {
			imports[i] = strconv.Quote(imports[i])
		}
	}
	return imports
}
