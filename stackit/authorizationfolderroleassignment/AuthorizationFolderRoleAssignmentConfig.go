package authorizationfolderroleassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthorizationFolderRoleAssignmentConfig struct {
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
	// folder Resource to assign the role to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/authorization_folder_role_assignment#resource_id AuthorizationFolderRoleAssignment#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Role to be assigned. Available roles can be queried using stackit-cli: `stackit curl https://authorization.api.stackit.cloud/v2/permissions`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/authorization_folder_role_assignment#role AuthorizationFolderRoleAssignment#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Identifier of user, service account or client. Usually email address or name in case of clients.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/authorization_folder_role_assignment#subject AuthorizationFolderRoleAssignment#subject}
	Subject *string `field:"required" json:"subject" yaml:"subject"`
}

