// Code generated by codegen using template resource_get.go.tpl; DO NOT EDIT.

package backup

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go-v2/service/backup"
)

func BackupVaults() *schema.Table {
	return &schema.Table{
		Name:      "aws_backup_vaults",
		Resolver:  fetchBackupVaults,
		Multiplex: client.ServiceAccountRegionMultiplexer("backup"),
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
				Name:     "backup_vault_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BackupVaultArn"),
			},
			{
				Name:     "backup_vault_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BackupVaultName"),
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationDate"),
			},
			{
				Name:     "creator_request_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatorRequestId"),
			},
			{
				Name:     "encryption_key_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EncryptionKeyArn"),
			},
			{
				Name:     "lock_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LockDate"),
			},
			{
				Name:     "locked",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Locked"),
			},
			{
				Name:     "max_retention_days",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxRetentionDays"),
			},
			{
				Name:     "min_retention_days",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinRetentionDays"),
			},
			{
				Name:     "number_of_recovery_points",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumberOfRecoveryPoints"),
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
		},

		Relations: []*schema.Table{
			BackupVaultsRecoveryPoints(),
		},
	}
}

func fetchBackupVaults(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Backup

	input := backup.ListBackupVaultsInput{
		MaxResults: aws.Int32(1000),
	}

	for {
		response, err := svc.ListBackupVaults(ctx, &input)
		if err != nil {

			return errors.WithStack(err)
		}

		res <- response.BackupVaultList

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
