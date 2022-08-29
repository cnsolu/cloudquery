// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"

	"{{.TypesImport}}"
{{range .Imports}}	"{{.}}"
{{end}}
)

func {{.TableFuncName}}() *schema.Table {
	return &schema.Table{{template "table.go.tpl" .Table}}
}

func {{.Table.Resolver}}(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().{{.AWSService}}

{{template "resolve_parent_defs.go.tpl" .}}
	input := {{.AWSService | ToLower}}.{{.ListVerb | Coalesce "List"}}{{.AWSSubService}}Input{
{{range .CustomInputs}}{{.}}
{{end}}{{template "resolve_parent_vars.go.tpl" .}}
	}
	paginator := {{.AWSService | ToLower}}.New{{.ListVerb | Coalesce "List"}}{{.AWSSubService}}Paginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, item := range output.{{.PaginatorListName}} {
			do, err := svc.{{.Verb | Coalesce "Describe"}}{{.ItemName}}(ctx, &{{.AWSService | ToLower}}.{{.Verb | Coalesce "Describe"}}{{.ItemName}}Input{
{{range .CustomInputs}}{{.}}
{{end}}{{if not .SkipDescribeParentInputs}}{{template "resolve_parent_vars.go.tpl" .}}{{end}}
			  {{.ListFieldName}}: {{if .RawDescribeFieldValue}}{{.RawDescribeFieldValue}}{{else}}item.{{.ListFieldName}}{{end}},
			})
			if err != nil {
				if cl.IsNotFoundError(err) {
					continue
				}
				return diag.WrapError(err)
			}
			res <- do.{{.ItemName}}
		}
	}
	return nil
}

{{if .HasTags}}
func resolve{{.AWSService | ToCamel}}{{.AWSSubService | ToCamel}}Tags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(*types.{{.AWSStructName}})
	cl := meta.(*client.Client)
	svc := cl.Services().{{.AWSService}}
	out, err := svc.ListTagsFor{{.ItemName}}(ctx, &{{.AWSService | ToLower}}.ListTagsFor{{.ItemName}}Input{
	  {{.ListFieldName}}: item.{{.ListFieldName}},
  })
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, client.TagsToMap(out.Tags)))
}
{{end}}

{{range .CustomResolvers}}{{.}}
{{end}}