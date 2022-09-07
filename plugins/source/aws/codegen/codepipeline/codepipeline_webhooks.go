// Code generated by codegen using template resource_get.go.tpl; DO NOT EDIT.

package codepipeline

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
)

func CodePipelineWebhooks() *schema.Table {
	return &schema.Table{
		Name:      "aws_codepipeline_webhooks",
		Resolver:  fetchCodePipelineWebhooks,
		Multiplex: client.ServiceAccountRegionMultiplexer("codepipeline"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
				Description: `The AWS Account ID of the resource.`,
			},
			{
				Name:        "region",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
				Description: `The AWS Region of the resource.`,
			},
			{
				Name:     "definition",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Definition"),
			},
			{
				Name:     "url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Url"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "error_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ErrorCode"),
			},
			{
				Name:     "error_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ErrorMessage"),
			},
			{
				Name:     "last_triggered",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastTriggered"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},
	}
}

func fetchCodePipelineWebhooks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().CodePipeline

	input := codepipeline.ListWebhooksInput{}

	for {
		response, err := svc.ListWebhooks(ctx, &input)
		if err != nil {

			return errors.WithStack(err)
		}

		res <- response.Webhooks

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
