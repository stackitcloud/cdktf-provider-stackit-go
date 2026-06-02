package authorizationfoldercustomrole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthorizationFolderCustomRoleConfig struct {
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
	// A human readable description of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/authorization_folder_custom_role#description AuthorizationFolderCustomRole#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/authorization_folder_custom_role#name AuthorizationFolderCustomRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Permissions for the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/authorization_folder_custom_role#permissions AuthorizationFolderCustomRole#permissions}
	Permissions *[]*string `field:"required" json:"permissions" yaml:"permissions"`
	// Resource to add the custom role to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/authorization_folder_custom_role#resource_id AuthorizationFolderCustomRole#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
}

