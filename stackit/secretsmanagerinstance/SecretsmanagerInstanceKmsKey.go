package secretsmanagerinstance


type SecretsmanagerInstanceKmsKey struct {
	// UUID of the key within the STACKIT-KMS to use for the encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/secretsmanager_instance#key_id SecretsmanagerInstance#key_id}
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
	// UUID of the keyring where the key is located within the STACKTI-KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/secretsmanager_instance#key_ring_id SecretsmanagerInstance#key_ring_id}
	KeyRingId *string `field:"required" json:"keyRingId" yaml:"keyRingId"`
	// Version of the key within the STACKIT-KMS to use for the encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/secretsmanager_instance#key_version SecretsmanagerInstance#key_version}
	KeyVersion *float64 `field:"required" json:"keyVersion" yaml:"keyVersion"`
	// Service-Account linked to the Key within the STACKIT-KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/secretsmanager_instance#service_account_email SecretsmanagerInstance#service_account_email}
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
}

