// Code generated by codegen using template resource_list_describe.go.tpl; DO NOT EDIT.

package cognito

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
)

func CognitoIdentityPools() *schema.Table {
	return &schema.Table{
		Name:      "aws_cognito_identity_pools",
		Resolver:  fetchCognitoIdentityPools,
		Multiplex: client.ServiceAccountRegionMultiplexer("cognito-identity"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
				Description: `The AWS Account ID of the resource.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "region",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
				Description: `The AWS Region of the resource.`,
			},
			{
				Name:     "allow_unauthenticated_identities",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AllowUnauthenticatedIdentities"),
			},
			{
				Name:     "identity_pool_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IdentityPoolId"),
			},
			{
				Name:     "identity_pool_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IdentityPoolName"),
			},
			{
				Name:     "allow_classic_flow",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AllowClassicFlow"),
			},
			{
				Name:     "cognito_identity_providers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CognitoIdentityProviders"),
			},
			{
				Name:     "developer_provider_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeveloperProviderName"),
			},
			{
				Name:     "identity_pool_tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IdentityPoolTags"),
			},
			{
				Name:     "open_id_connect_provider_ar_ns",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("OpenIdConnectProviderARNs"),
			},
			{
				Name:     "saml_provider_ar_ns",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SamlProviderARNs"),
			},
			{
				Name:     "supported_login_providers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SupportedLoginProviders"),
			},
		},
	}
}

func fetchCognitoIdentityPools(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().CognitoIdentityPools

	input := cognito.ListIdentityPoolsInput{
		MaxResults: 60, // we want max results to reduce List calls as much as possible, services limited to less than or equal to 60

	}
	paginator := cognito.NewListIdentityPoolsPaginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {

			return errors.WithStack(err)
		}

		for _, item := range output.IdentityPools {

			do, err := svc.DescribeIdentityPool(ctx, &cognito.DescribeIdentityPoolInput{

				IdentityPoolId: item.IdentityPoolId,
			})
			if err != nil {

				if cl.IsNotFoundError(err) {
					continue
				}
				return errors.WithStack(err)
			}
			res <- do
		}
	}
	return nil
}
