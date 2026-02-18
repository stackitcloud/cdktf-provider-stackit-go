package kmswrappingkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KmsWrappingKeyConfig struct {
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
	// The wrapping algorithm used to wrap the key to import.
	//
	// Possible values are: `rsa_2048_oaep_sha256`, `rsa_3072_oaep_sha256`, `rsa_4096_oaep_sha256`, `rsa_4096_oaep_sha512`, `rsa_2048_oaep_sha256_aes_256_key_wrap`, `rsa_3072_oaep_sha256_aes_256_key_wrap`, `rsa_4096_oaep_sha256_aes_256_key_wrap`, `rsa_4096_oaep_sha512_aes_256_key_wrap`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/kms_wrapping_key#algorithm KmsWrappingKey#algorithm}
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
	// The display name to distinguish multiple wrapping keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/kms_wrapping_key#display_name KmsWrappingKey#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The ID of the associated keyring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/kms_wrapping_key#keyring_id KmsWrappingKey#keyring_id}
	KeyringId *string `field:"required" json:"keyringId" yaml:"keyringId"`
	// STACKIT project ID to which the keyring is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/kms_wrapping_key#project_id KmsWrappingKey#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The underlying system that is responsible for protecting the key material. Possible values are: `software`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/kms_wrapping_key#protection KmsWrappingKey#protection}
	Protection *string `field:"required" json:"protection" yaml:"protection"`
	// The purpose for which the key will be used. Possible values are: `wrap_symmetric_key`, `wrap_asymmetric_key`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/kms_wrapping_key#purpose KmsWrappingKey#purpose}
	Purpose *string `field:"required" json:"purpose" yaml:"purpose"`
	// The access scope of the key. Default is `PUBLIC`. Possible values are: `PUBLIC`, `SNA`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/kms_wrapping_key#access_scope KmsWrappingKey#access_scope}
	AccessScope *string `field:"optional" json:"accessScope" yaml:"accessScope"`
	// A user chosen description to distinguish multiple wrapping keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/kms_wrapping_key#description KmsWrappingKey#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/kms_wrapping_key#region KmsWrappingKey#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

