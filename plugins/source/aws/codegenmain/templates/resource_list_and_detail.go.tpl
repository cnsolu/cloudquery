// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
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
  return diag.WrapError(client.ListAndDetailResolver(ctx, meta, res, list{{.AWSSubService}}, list{{.AWSSubService}}Detail))
}

func list{{.AWSSubService}}(ctx context.Context, meta schema.ClientMeta, detailChan chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().{{.AWSService | ToCamel}}

{{template "resolve_parent_defs.go.tpl" .}}
	input := {{.AWSService | ToLower}}.{{.ListVerb | Coalesce "List"}}{{.AWSSubService}}Input{
{{range .CustomInputs}}{{.}}
{{end}}
{{template "resolve_parent_vars.go.tpl" .}}
	}

	for {
		response, err := svc.{{.ListVerb | Coalesce "List"}}{{.AWSSubService}}(ctx, &input)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, item := range response.{{.ListFieldName}} {
			detailChan <- item
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func list{{.AWSSubService}}Detail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, listInfo interface{}) {
	cl := meta.(*client.Client)
	itemSummary := listInfo.(types.{{.ResponseItemsType}})
	svc := cl.Services().{{.AWSService | ToCamel}}
	response, err := svc.{{.Verb | Coalesce "Get" }}{{.ItemName}}(ctx, &{{.AWSService | ToLower}}.{{.Verb | Coalesce "Get" }}{{.ItemName}}Input{
		{{.DetailInputFieldName}}: itemSummary.{{.ResponseItemsName}},
	})
	if err != nil {
		{{.CustomErrorBlock}}
		if cl.IsNotFoundError(err) {
			return
		}
		errorChan <- diag.WrapError(err)
		return
	}
	resultsChan <- *response.{{.ItemName}}
}

{{if .HasTags}}
func resolve{{.AWSService}}{{.AWSSubService}}Tags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().{{.AWSService | ToCamel}}
	item := resource.Item.(types.{{.ItemName}})
	params := {{.AWSService | ToLower}}.ListTagsForResourceInput{
		ResourceARN: {{.CustomTagField | Coalesce "item.ARN"}},
	}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTagsForResource(ctx, &params)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return diag.WrapError(err)
		}
		client.TagsIntoMap(result.Tags, tags)
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return diag.WrapError(resource.Set(c.Name, tags))
}
{{end}}

{{range .CustomResolvers}}{{.}}
{{end}}