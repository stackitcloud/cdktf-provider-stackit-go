package datastackitdremiouser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitDremioUserConfig struct {
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
	// The Dremio instance ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/data-sources/dremio_user#instance_id DataStackitDremioUser#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// STACKIT Project ID to which the resource is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/data-sources/dremio_user#project_id DataStackitDremioUser#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The Dremio user ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/data-sources/dremio_user#user_id DataStackitDremioUser#user_id}
	UserId *string `field:"required" json:"userId" yaml:"userId"`
	// The description of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/data-sources/dremio_user#description DataStackitDremioUser#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The STACKIT region name the resource is located in. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/data-sources/dremio_user#region DataStackitDremioUser#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

