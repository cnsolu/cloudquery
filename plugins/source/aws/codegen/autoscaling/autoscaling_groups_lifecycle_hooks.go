// Code generated by codegen using template resource_get.go.tpl; DO NOT EDIT.

package autoscaling

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	resolvers "github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/resolvers/autoscaling"
)

func AutoscalingGroupsLifecycleHooks() *schema.Table {
	return &schema.Table{
		Name:      "aws_autoscaling_groups_lifecycle_hooks",
		Resolver:  fetchAutoscalingGroupsLifecycleHooks,
		Multiplex: client.ServiceAccountRegionMultiplexer("autoscaling"),
		Columns: []schema.Column{
			{
				Name:     "group_cq_id",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
			},
			{
				Name:     "auto_scaling_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoScalingGroupName"),
			},
			{
				Name:     "default_result",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultResult"),
			},
			{
				Name:     "global_timeout",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("GlobalTimeout"),
			},
			{
				Name:     "heartbeat_timeout",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("HeartbeatTimeout"),
			},
			{
				Name:     "lifecycle_hook_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LifecycleHookName"),
			},
			{
				Name:     "lifecycle_transition",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LifecycleTransition"),
			},
			{
				Name:     "notification_metadata",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NotificationMetadata"),
			},
			{
				Name:     "notification_target_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NotificationTargetARN"),
			},
			{
				Name:     "role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleARN"),
			},
		},
	}
}

func fetchAutoscalingGroupsLifecycleHooks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Autoscaling

	r1 := parent.Item.(types.AutoScalingGroup)

	input := autoscaling.DescribeLifecycleHooksInput{
		AutoScalingGroupName: r1.AutoScalingGroupName,
	}

	{
		response, err := svc.DescribeLifecycleHooks(ctx, &input)
		if err != nil {

			if resolvers.IsGroupNotExistsError(err) {
				return nil
			}
			return errors.WithStack(err)
		}

		res <- response.LifecycleHooks

	}
	return nil
}
