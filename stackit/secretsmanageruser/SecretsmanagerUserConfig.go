package secretsmanageruser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecretsmanagerUserConfig struct {
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
	// A user chosen description to differentiate between multiple users. Can't be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/secretsmanager_user#description SecretsmanagerUser#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// ID of the Secrets Manager instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/secretsmanager_user#instance_id SecretsmanagerUser#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// STACKIT Project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/secretsmanager_user#project_id SecretsmanagerUser#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// If true, the user has writeaccess to the secrets engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/secretsmanager_user#write_enabled SecretsmanagerUser#write_enabled}
	WriteEnabled interface{} `field:"required" json:"writeEnabled" yaml:"writeEnabled"`
}

