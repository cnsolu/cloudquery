// Code generated by codegen using template resource_get_mock_test.go.tpl; DO NOT EDIT.

package autoscaling

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

func buildAutoscalingGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockAutoscalingClient(ctrl)

	item := types.AutoScalingGroup{}

	err := faker.FakeData(&item)
	if err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeAutoScalingGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(

		&autoscaling.DescribeAutoScalingGroupsOutput{
			AutoScalingGroups: []types.AutoScalingGroup{item},
		}, nil)

	return client.Services{
		Autoscaling: mock,
	}
}

func TestAutoscalingGroups(t *testing.T) {
	client.MockTestHelper(t, AutoscalingGroups(), buildAutoscalingGroups, client.TestOptions{})
}
