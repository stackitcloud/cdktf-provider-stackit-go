package mariadbinstance


type MariadbInstanceParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.15.2/docs/resources/mariadb_instance#sgw_acl MariadbInstance#sgw_acl}.
	SgwAcl *string `field:"optional" json:"sgwAcl" yaml:"sgwAcl"`
}

