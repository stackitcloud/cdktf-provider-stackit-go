package datastackitsqlserverflexinstance


type DataStackitSqlserverflexInstanceNetwork struct {
	// The network access scope of the instance.
	//
	// This feature is in private preview. Supplying this object is only permitted for enabled accounts. If your account does not have access, the request will be rejected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/data-sources/sqlserverflex_instance#access_scope DataStackitSqlserverflexInstance#access_scope}
	AccessScope *string `field:"optional" json:"accessScope" yaml:"accessScope"`
}

