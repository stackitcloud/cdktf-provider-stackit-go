package datastackitsecretsmanagerinstancerolebindingsv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitSecretsmanagerInstanceRoleBindingsV1Config struct {
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
	// The identifier of the resource to get the role bindings for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/data-sources/secretsmanager_instance_role_bindings_v1#resource_id DataStackitSecretsmanagerInstanceRoleBindingsV1#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/data-sources/secretsmanager_instance_role_bindings_v1#region DataStackitSecretsmanagerInstanceRoleBindingsV1#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

