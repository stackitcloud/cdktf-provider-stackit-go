package dremioinstance


type DremioInstanceAuthenticationOauthJwtClaims struct {
	// Mapped user name claim (e.g. email).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/dremio_instance#user_name DremioInstance#user_name}
	UserName *string `field:"required" json:"userName" yaml:"userName"`
}

