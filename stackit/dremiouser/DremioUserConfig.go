package dremiouser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DremioUserConfig struct {
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
	// The email address of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/dremio_user#email DremioUser#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// The first name of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/dremio_user#first_name DremioUser#first_name}
	FirstName *string `field:"required" json:"firstName" yaml:"firstName"`
	// The Dremio instance ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/dremio_user#instance_id DremioUser#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The last name of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/dremio_user#last_name DremioUser#last_name}
	LastName *string `field:"required" json:"lastName" yaml:"lastName"`
	// The username of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/dremio_user#name DremioUser#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The password of the user.
	//
	// Only used for creation and updates. Must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one number and one special character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/dremio_user#password DremioUser#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// STACKIT Project ID to which the resource is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/dremio_user#project_id DremioUser#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The description of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/dremio_user#description DremioUser#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The STACKIT region name the resource is located in. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/dremio_user#region DremioUser#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/dremio_user#timeouts DremioUser#timeouts}.
	Timeouts *DremioUserTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

