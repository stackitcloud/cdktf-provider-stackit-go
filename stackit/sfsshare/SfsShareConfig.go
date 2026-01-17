package sfsshare

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SfsShareConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Name of the Share Export Policy to use in the Share.
	//
	// Note that if this is set to an empty string, the Share can only be mounted in read only by
	// clients with IPs matching the IP ACL of the Resource Pool hosting this Share.
	// You can also assign a Share Export Policy after creating the Share
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/resources/sfs_share#export_policy SfsShare#export_policy}
	ExportPolicy *string `field:"required" json:"exportPolicy" yaml:"exportPolicy"`
	// Name of the share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/resources/sfs_share#name SfsShare#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the share is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/resources/sfs_share#project_id SfsShare#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The ID of the resource pool for the SFS share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/resources/sfs_share#resource_pool_id SfsShare#resource_pool_id}
	ResourcePoolId *string `field:"required" json:"resourcePoolId" yaml:"resourcePoolId"`
	// Space hard limit for the Share.
	//
	// If zero, the Share will have access to the full space of the Resource Pool it lives in.
	// 				(unit: gigabytes)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/resources/sfs_share#space_hard_limit_gigabytes SfsShare#space_hard_limit_gigabytes}
	SpaceHardLimitGigabytes *float64 `field:"required" json:"spaceHardLimitGigabytes" yaml:"spaceHardLimitGigabytes"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/resources/sfs_share#region SfsShare#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

