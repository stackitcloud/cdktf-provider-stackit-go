package dremioinstance


type DremioInstanceAuthenticationOauthParameters struct {
	// Parameter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/dremio_instance#name DremioInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Parameter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/dremio_instance#value DremioInstance#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

