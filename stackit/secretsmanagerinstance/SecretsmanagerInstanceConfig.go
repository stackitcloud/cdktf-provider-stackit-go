package secretsmanagerinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecretsmanagerInstanceConfig struct {
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
	// Instance name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.1/docs/resources/secretsmanager_instance#name SecretsmanagerInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.1/docs/resources/secretsmanager_instance#project_id SecretsmanagerInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The access control list for this instance.
	//
	// Each entry is an IP or IP range that is permitted to access, in CIDR notation
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.1/docs/resources/secretsmanager_instance#acls SecretsmanagerInstance#acls}
	Acls *[]*string `field:"optional" json:"acls" yaml:"acls"`
}

