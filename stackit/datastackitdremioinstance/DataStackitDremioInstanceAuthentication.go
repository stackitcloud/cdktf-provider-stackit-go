package datastackitdremioinstance


type DataStackitDremioInstanceAuthentication struct {
	// Azure Active Directory authentication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/data-sources/dremio_instance#azuread DataStackitDremioInstance#azuread}
	Azuread *DataStackitDremioInstanceAuthenticationAzuread `field:"optional" json:"azuread" yaml:"azuread"`
	// OIDC authentication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/data-sources/dremio_instance#oauth DataStackitDremioInstance#oauth}
	Oauth *DataStackitDremioInstanceAuthenticationOauth `field:"optional" json:"oauth" yaml:"oauth"`
}

