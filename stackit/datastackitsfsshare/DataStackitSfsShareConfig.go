package datastackitsfsshare

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitSfsShareConfig struct {
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
	// STACKIT project ID to which the share is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/data-sources/sfs_share#project_id DataStackitSfsShare#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The ID of the resource pool for the SFS share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/data-sources/sfs_share#resource_pool_id DataStackitSfsShare#resource_pool_id}
	ResourcePoolId *string `field:"required" json:"resourcePoolId" yaml:"resourcePoolId"`
	// share ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/data-sources/sfs_share#share_id DataStackitSfsShare#share_id}
	ShareId *string `field:"required" json:"shareId" yaml:"shareId"`
	// The resource region. Read-only attribute that reflects the provider region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/data-sources/sfs_share#region DataStackitSfsShare#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

