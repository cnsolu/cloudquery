// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
)

func AthenaDataCatalogs() *schema.Table {
	return &schema.Table{
		Name:      "aws_athena_data_catalogs",
		Resolver:  fetchAthenaDataCatalogs,
		Multiplex: client.ServiceAccountRegionMultiplexer("athena"),
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
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "parameters",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Parameters"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveAthenaDataCatalogArn,
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveAthenaDataCatalogsTags,
				Description: `Tags associated with the Athena data catalog.`,
			},
		},
	}
}

func fetchAthenaDataCatalogs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return diag.WrapError(client.ListAndDetailResolver(ctx, meta, res, listDataCatalogs, listDataCatalogsDetail))
}

func listDataCatalogs(ctx context.Context, meta schema.ClientMeta, detailChan chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Athena

	input := athena.ListDataCatalogsInput{}

	for {
		response, err := svc.ListDataCatalogs(ctx, &input)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, item := range response.DataCatalogsSummary {
			detailChan <- item
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func listDataCatalogsDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, listInfo interface{}) {
	cl := meta.(*client.Client)
	itemSummary := listInfo.(types.DataCatalogSummary)
	svc := cl.Services().Athena
	response, err := svc.GetDataCatalog(ctx, &athena.GetDataCatalogInput{
		Name: itemSummary.CatalogName,
	})
	if err != nil {

		// retrieving of default data catalog (AwsDataCatalog) returns "not found error" but it exists and its
		// relations can be fetched by its name
		if *itemSummary.CatalogName == "AwsDataCatalog" {
			resultsChan <- types.DataCatalog{Name: itemSummary.CatalogName, Type: itemSummary.Type}
			return
		}

		if cl.IsNotFoundError(err) {
			return
		}
		errorChan <- diag.WrapError(err)
		return
	}
	resultsChan <- *response.DataCatalog
}

func resolveAthenaDataCatalogsTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Athena
	item := resource.Item.(types.DataCatalog)
	params := athena.ListTagsForResourceInput{
		ResourceARN: aws.String(createDataCatalogArn(cl, *item.Name)),
	}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTagsForResource(ctx, &params)
		if err != nil {

			// retrieving of default data catalog (AwsDataCatalog) returns "not found error" but it exists and its
			// relations can be fetched by its name
			if *itemSummary.CatalogName == "AwsDataCatalog" {
				resultsChan <- types.DataCatalog{Name: itemSummary.CatalogName, Type: itemSummary.Type}
				return
			}

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

func resolveAthenaDataCatalogArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	dc := resource.Item.(types.DataCatalog)
	return diag.WrapError(resource.Set(c.Name, createDataCatalogArn(cl, *dc.Name)))
}

func createDataCatalogArn(cl *client.Client, catalogName string) string {
	return cl.ARN(client.Athena, "datacatalog", catalogName)
}
