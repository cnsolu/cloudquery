// Code generated by codegen using template resource_get_mock_test.go.tpl; DO NOT EDIT.

package cloudformation

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

func buildCloudformationStackResources(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCloudformationClient(ctrl)

	item := types.StackResourceSummary{}

	err := faker.FakeData(&item)
	if err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListStackResources(gomock.Any(), gomock.Any(), gomock.Any()).Return(

		&cloudformation.ListStackResourcesOutput{
			StackResourceSummaries: []types.StackResourceSummary{item},
		}, nil)

	return client.Services{
		Cloudformation: mock,
	}
}

func TestCloudformationStackResources(t *testing.T) {
	client.MockTestHelper(t, CloudformationStackResources(), buildCloudformationStackResources, client.TestOptions{})
}
