package datastackitauthorizationfoldercustomrole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitAuthorizationFolderCustomRoleConfig struct {
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
	// Resource to add the custom role to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/data-sources/authorization_folder_custom_role#resource_id DataStackitAuthorizationFolderCustomRole#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The ID of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/data-sources/authorization_folder_custom_role#role_id DataStackitAuthorizationFolderCustomRole#role_id}
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
}

