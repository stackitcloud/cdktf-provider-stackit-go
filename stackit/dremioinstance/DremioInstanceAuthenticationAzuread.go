package dremioinstance


type DremioInstanceAuthenticationAzuread struct {
	// The Azure AD authority URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/dremio_instance#authority_url DremioInstance#authority_url}
	AuthorityUrl *string `field:"required" json:"authorityUrl" yaml:"authorityUrl"`
	// The Azure AD client ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/dremio_instance#client_id DremioInstance#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The Azure AD client secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/dremio_instance#client_secret DremioInstance#client_secret}
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
}

