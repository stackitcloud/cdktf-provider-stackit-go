package kmskeyring

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KmsKeyringConfig struct {
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
	// The display name to distinguish multiple keyrings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/kms_keyring#display_name KmsKeyring#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// STACKIT project ID to which the keyring is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/kms_keyring#project_id KmsKeyring#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// A user chosen description to distinguish multiple keyrings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/kms_keyring#description KmsKeyring#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/kms_keyring#region KmsKeyring#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

