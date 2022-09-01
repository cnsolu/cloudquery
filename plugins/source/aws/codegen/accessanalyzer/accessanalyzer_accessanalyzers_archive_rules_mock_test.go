// Code generated by codegen; DO NOT EDIT.

package accessanalyzer

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
)

func buildAccessAnalyzerAccessanalyzersArchiveRules(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockAccessAnalyzerClient(ctrl)

	item := types.ArchiveRuleSummary{}

	err := faker.FakeData(&item)
	if err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListArchiveRules(gomock.Any(), gomock.Any(), gomock.Any()).Return(

		&accessanalyzer.ListArchiveRulesOutput{
			ArchiveRules: []types.ArchiveRuleSummary{item},
		}, nil)

	return client.Services{
		AccessAnalyzer: mock,
	}
}

func TestAccessAnalyzerAccessanalyzersArchiveRules(t *testing.T) {
	client.AwsMockTestHelper(t, AccessAnalyzerAccessanalyzersArchiveRules(), buildAccessAnalyzerAccessanalyzersArchiveRules, client.TestOptions{})
}
