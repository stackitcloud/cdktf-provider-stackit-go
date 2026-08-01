package sqlserverflexinstance


type SqlserverflexInstanceEncryption struct {
	// UUID of the key within the STACKIT-KMS to use for the encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.107.1/docs/resources/sqlserverflex_instance#kek_key_id SqlserverflexInstance#kek_key_id}
	KekKeyId *string `field:"required" json:"kekKeyId" yaml:"kekKeyId"`
	// UUID of the keyring where the key is located within the STACKTI-KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.107.1/docs/resources/sqlserverflex_instance#kek_keyring_id SqlserverflexInstance#kek_keyring_id}
	KekKeyringId *string `field:"required" json:"kekKeyringId" yaml:"kekKeyringId"`
	// Version of the key within the STACKIT-KMS to use for the encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.107.1/docs/resources/sqlserverflex_instance#kek_key_version SqlserverflexInstance#kek_key_version}
	KekKeyVersion *string `field:"required" json:"kekKeyVersion" yaml:"kekKeyVersion"`
	// Service-Account linked to the Key within the STACKIT-KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.107.1/docs/resources/sqlserverflex_instance#service_account SqlserverflexInstance#service_account}
	ServiceAccount *string `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
}

