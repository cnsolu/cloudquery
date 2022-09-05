// Code generated by codegen using template resource_get.go.tpl; DO NOT EDIT.

package cloudtrail

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
)

func CloudtrailTrailEventSelectors() *schema.Table {
	return &schema.Table{
		Name:      "aws_cloudtrail_trail_event_selectors",
		Resolver:  fetchCloudtrailTrailEventSelectors,
		Multiplex: client.ServiceAccountRegionMultiplexer("cloudtrail"),
		Columns: []schema.Column{
			{
				Name:     "trail_cq_id",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
			},
			{
				Name:     "data_resources",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DataResources"),
			},
			{
				Name:     "exclude_management_event_sources",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ExcludeManagementEventSources"),
			},
			{
				Name:     "include_management_events",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IncludeManagementEvents"),
			},
			{
				Name:     "read_write_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReadWriteType"),
			},
		},
	}
}

func fetchCloudtrailTrailEventSelectors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudtrail

	r1 := parent.Item.(types.Trail)

	input := cloudtrail.GetEventSelectorsInput{
		TrailName: r1.TrailARN,
	}

	{
		response, err := svc.GetEventSelectors(ctx, &input)
		if err != nil {

			return diag.WrapError(err)
		}

		res <- response.EventSelectors

	}
	return nil
}
