package dremioinstance


type DremioInstanceAuthentication struct {
	// Type of authentication (local-only, azuread, oauth).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/dremio_instance#type DremioInstance#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Azure Active Directory authentication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/dremio_instance#azuread DremioInstance#azuread}
	Azuread *DremioInstanceAuthenticationAzuread `field:"optional" json:"azuread" yaml:"azuread"`
	// OIDC authentication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/dremio_instance#oauth DremioInstance#oauth}
	Oauth *DremioInstanceAuthenticationOauth `field:"optional" json:"oauth" yaml:"oauth"`
}

