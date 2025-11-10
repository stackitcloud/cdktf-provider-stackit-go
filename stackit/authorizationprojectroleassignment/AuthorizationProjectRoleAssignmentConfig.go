package authorizationprojectroleassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthorizationProjectRoleAssignmentConfig struct {
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
	// project Resource to assign the role to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.48.0/docs/resources/authorization_project_role_assignment#resource_id AuthorizationProjectRoleAssignment#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Role to be assigned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.48.0/docs/resources/authorization_project_role_assignment#role AuthorizationProjectRoleAssignment#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Identifier of user, service account or client. Usually email address or name in case of clients.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.48.0/docs/resources/authorization_project_role_assignment#subject AuthorizationProjectRoleAssignment#subject}
	Subject *string `field:"required" json:"subject" yaml:"subject"`
}

