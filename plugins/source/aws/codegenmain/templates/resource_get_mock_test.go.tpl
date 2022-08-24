// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

{{range .Imports}}	"{{.}}"
{{end}}
)

{{if .Parent}}
func build{{.AWSService | ToCamel}}{{.Parent.AWSSubService | ToCamel}}{{.AWSSubService | ToCamel}}(t *testing.T, ctrl *gomock.Controller) client.Services {
{{else}}
func build{{.AWSService | ToCamel}}{{.AWSSubService | ToCamel}}(t *testing.T, ctrl *gomock.Controller) client.Services {
{{end}}
	mock := mocks.NewMock{{.AWSService | ToCamel}}Client(ctrl)

	item := types.{{.ItemName}}{}
	err := faker.FakeData(&item)
	if err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().Get{{.AWSSubService | ToCamel}}(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&{{.AWSService | ToLower}}.Get{{.AWSSubService | ToCamel}}Output{
			Items: []types.{{.ItemName}}{item},
		}, nil)
	return client.Services{
		{{.AWSService | ToCamel}}: mock,
	}
}

{{if .Parent}}
func Test{{.AWSService | ToCamel}}{{.Parent.AWSSubService | ToCamel}}{{.AWSSubService | ToCamel}}(t *testing.T) {
	client.AwsMockTestHelper(t, {{.AWSService | ToCamel}}{{.Parent.AWSSubService | ToCamel}}{{.AWSSubService | ToCamel}}(), build{{.AWSService | ToCamel}}{{.Parent.AWSSubService | ToCamel}}{{.AWSSubService | ToCamel}}, client.TestOptions{})
}
{{else}}
func Test{{.AWSService | ToCamel}}{{.AWSSubService | ToCamel}}(t *testing.T) {
	client.AwsMockTestHelper(t, {{.AWSService | ToCamel}}{{.AWSSubService | ToCamel}}(), build{{.AWSService | ToCamel}}{{.AWSSubService | ToCamel}}, client.TestOptions{})
}
{{end}}
