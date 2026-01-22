package datastackitkmswrappingkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitKmsWrappingKeyConfig struct {
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
	// The ID of the associated keyring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/data-sources/kms_wrapping_key#keyring_id DataStackitKmsWrappingKey#keyring_id}
	KeyringId *string `field:"required" json:"keyringId" yaml:"keyringId"`
	// STACKIT project ID to which the keyring is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/data-sources/kms_wrapping_key#project_id DataStackitKmsWrappingKey#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The ID of the wrapping key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/data-sources/kms_wrapping_key#wrapping_key_id DataStackitKmsWrappingKey#wrapping_key_id}
	WrappingKeyId *string `field:"required" json:"wrappingKeyId" yaml:"wrappingKeyId"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/data-sources/kms_wrapping_key#region DataStackitKmsWrappingKey#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

