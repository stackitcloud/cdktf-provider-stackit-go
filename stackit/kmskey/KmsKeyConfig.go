package kmskey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KmsKeyConfig struct {
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
	// The encryption algorithm that the key will use to encrypt data.
	//
	// Possible values are: `aes_256_gcm`, `rsa_2048_oaep_sha256`, `rsa_3072_oaep_sha256`, `rsa_4096_oaep_sha256`, `rsa_4096_oaep_sha512`, `hmac_sha256`, `hmac_sha384`, `hmac_sha512`, `ecdsa_p256_sha256`, `ecdsa_p384_sha384`, `ecdsa_p521_sha512`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.75.0/docs/resources/kms_key#algorithm KmsKey#algorithm}
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
	// The display name to distinguish multiple keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.75.0/docs/resources/kms_key#display_name KmsKey#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The ID of the associated keyring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.75.0/docs/resources/kms_key#keyring_id KmsKey#keyring_id}
	KeyringId *string `field:"required" json:"keyringId" yaml:"keyringId"`
	// STACKIT project ID to which the key is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.75.0/docs/resources/kms_key#project_id KmsKey#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The underlying system that is responsible for protecting the key material. Possible values are: `software`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.75.0/docs/resources/kms_key#protection KmsKey#protection}
	Protection *string `field:"required" json:"protection" yaml:"protection"`
	// The purpose for which the key will be used. Possible values are: `symmetric_encrypt_decrypt`, `asymmetric_encrypt_decrypt`, `message_authentication_code`, `asymmetric_sign_verify`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.75.0/docs/resources/kms_key#purpose KmsKey#purpose}
	Purpose *string `field:"required" json:"purpose" yaml:"purpose"`
	// The access scope of the key. Default is `PUBLIC`. Possible values are: `PUBLIC`, `SNA`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.75.0/docs/resources/kms_key#access_scope KmsKey#access_scope}
	AccessScope *string `field:"optional" json:"accessScope" yaml:"accessScope"`
	// A user chosen description to distinguish multiple keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.75.0/docs/resources/kms_key#description KmsKey#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// States whether versions can be created or only imported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.75.0/docs/resources/kms_key#import_only KmsKey#import_only}
	ImportOnly interface{} `field:"optional" json:"importOnly" yaml:"importOnly"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.75.0/docs/resources/kms_key#region KmsKey#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

