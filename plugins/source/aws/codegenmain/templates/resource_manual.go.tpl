// Code generated by codegen using template {{.TemplateFilename}}; DO NOT EDIT.

package {{.AWSService | ToLower}}

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"

{{range .Imports}}	{{.}}
{{end}}
)

func {{.TableFuncName}}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}
