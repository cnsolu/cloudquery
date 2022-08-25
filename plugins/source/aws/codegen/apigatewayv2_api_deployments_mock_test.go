// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
)

func buildApigatewayv2ApiDeployments(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockApigatewayv2Client(ctrl)

	item := types.Deployment{}
	err := faker.FakeData(&item)
	if err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetDeployments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetDeploymentsOutput{
			Items: []types.Deployment{item},
		}, nil)
	return client.Services{
		Apigatewayv2: mock,
	}
}

func TestApigatewayv2ApiDeployments(t *testing.T) {
	client.AwsMockTestHelper(t, Apigatewayv2ApiDeployments(), buildApigatewayv2ApiDeployments, client.TestOptions{})
}
