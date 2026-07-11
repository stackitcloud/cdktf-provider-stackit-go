package sqlserverflexinstance


type SqlserverflexInstanceNetwork struct {
	// The network access scope of the instance.
	//
	// This feature is in private preview. Supplying this object is only permitted for enabled accounts. If your account does not have access, the request will be rejected. Possible values are: `PUBLIC`, `SNA`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/sqlserverflex_instance#access_scope SqlserverflexInstance#access_scope}
	AccessScope *string `field:"optional" json:"accessScope" yaml:"accessScope"`
	// List of IPV4 cidr.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/sqlserverflex_instance#acl SqlserverflexInstance#acl}
	Acl *[]*string `field:"optional" json:"acl" yaml:"acl"`
}

