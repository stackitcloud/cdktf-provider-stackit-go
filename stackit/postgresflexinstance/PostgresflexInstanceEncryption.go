package postgresflexinstance


type PostgresflexInstanceEncryption struct {
	// The ID of the Key within the STACKIT-KMS to use for the encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.109.0/docs/resources/postgresflex_instance#kek_key_id PostgresflexInstance#kek_key_id}
	KekKeyId *string `field:"required" json:"kekKeyId" yaml:"kekKeyId"`
	// The ID of the keyring where the key is located within the STACKTI-KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.109.0/docs/resources/postgresflex_instance#kek_keyring_id PostgresflexInstance#kek_keyring_id}
	KekKeyringId *string `field:"required" json:"kekKeyringId" yaml:"kekKeyringId"`
	// Version of the key within the STACKIT-KMS to use for the encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.109.0/docs/resources/postgresflex_instance#kek_key_version PostgresflexInstance#kek_key_version}
	KekKeyVersion *string `field:"required" json:"kekKeyVersion" yaml:"kekKeyVersion"`
	// Service-Account linked to the Key within the STACKIT-KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.109.0/docs/resources/postgresflex_instance#service_account PostgresflexInstance#service_account}
	ServiceAccount *string `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
}

