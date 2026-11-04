package cdndistribution


type CdnDistributionConfigLogSink struct {
	// Authentication credentials the CDN uses when pushing logs.
	//
	// Loki requires `username` and `password`. OTLP requires `type` to be set to `basic` (with `username` and `password`) or `bearer` (with `token`).
	//
	// ~> Write-Only arguments `password_wo` and `token_wo` are available to use in place of `password` and `token`. Write-Only arguments are supported in HashiCorp Terraform 1.11.0 and later. [Learn more](https://developer.hashicorp.com/terraform/language/resources/ephemeral#write-only-arguments).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.122.0/docs/resources/cdn_distribution#credentials CdnDistribution#credentials}
	Credentials *CdnDistributionConfigLogSinkCredentials `field:"required" json:"credentials" yaml:"credentials"`
	// The fully qualified URL where the CDN should push access logs (for example, `https://loki.example.com/loki/api/v1/push` for Loki or `https://otlp.example.com/otlp/v1/logs` for OTLP).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.122.0/docs/resources/cdn_distribution#push_url CdnDistribution#push_url}
	PushUrl *string `field:"required" json:"pushUrl" yaml:"pushUrl"`
	// The log sink protocolPossible values are: `loki`, `otlp`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.122.0/docs/resources/cdn_distribution#type CdnDistribution#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

