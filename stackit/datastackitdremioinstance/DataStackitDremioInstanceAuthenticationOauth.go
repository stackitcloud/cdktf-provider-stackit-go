package datastackitdremioinstance


type DataStackitDremioInstanceAuthenticationOauth struct {
	// Any additional parameters the Identity Provider requires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/data-sources/dremio_instance#parameters DataStackitDremioInstance#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A list of space-separated scopes. The `openid` scope is always required; other scopes can vary by provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/data-sources/dremio_instance#scope DataStackitDremioInstance#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

