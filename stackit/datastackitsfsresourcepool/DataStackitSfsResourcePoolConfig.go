package datastackitsfsresourcepool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitSfsResourcePoolConfig struct {
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
	// STACKIT project ID to which the resource pool is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.77.0/docs/data-sources/sfs_resource_pool#project_id DataStackitSfsResourcePool#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Resourcepool ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.77.0/docs/data-sources/sfs_resource_pool#resource_pool_id DataStackitSfsResourcePool#resource_pool_id}
	ResourcePoolId *string `field:"required" json:"resourcePoolId" yaml:"resourcePoolId"`
	// The resource region. Read-only attribute that reflects the provider region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.77.0/docs/data-sources/sfs_resource_pool#region DataStackitSfsResourcePool#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

