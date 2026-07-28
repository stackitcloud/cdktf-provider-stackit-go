package postgresflexinstance


type PostgresflexInstanceNetwork struct {
	// The network access scope of the instance.
	//
	// This feature is in private preview. Supplying this object is only permitted for enabled accounts. If your account does not have access, the request will be rejected. Possible values are: `PUBLIC`, `SNA`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.105.0/docs/resources/postgresflex_instance#access_scope PostgresflexInstance#access_scope}
	AccessScope *string `field:"optional" json:"accessScope" yaml:"accessScope"`
	// The Access Control List (ACL) for the PostgresFlex instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.105.0/docs/resources/postgresflex_instance#acl PostgresflexInstance#acl}
	Acl *[]*string `field:"optional" json:"acl" yaml:"acl"`
}

