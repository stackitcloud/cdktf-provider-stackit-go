package secretsmanagerinstancerolebindingv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecretsmanagerInstanceRoleBindingV1Config struct {
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
	// The identifier of the resource to apply this role binding to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.95.0/docs/resources/secretsmanager_instance_role_binding_v1#resource_id SecretsmanagerInstanceRoleBindingV1#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// A valid role defined for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.95.0/docs/resources/secretsmanager_instance_role_binding_v1#role SecretsmanagerInstanceRoleBindingV1#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Identifier of user, service account or client. Usually email address or name in case of clients.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.95.0/docs/resources/secretsmanager_instance_role_binding_v1#subject SecretsmanagerInstanceRoleBindingV1#subject}
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.95.0/docs/resources/secretsmanager_instance_role_binding_v1#region SecretsmanagerInstanceRoleBindingV1#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

