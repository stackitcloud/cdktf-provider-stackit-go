package cdndistribution


type CdnDistributionConfigBackend struct {
	// The configured backend type for the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.66.0/docs/resources/cdn_distribution#origin_url CdnDistribution#origin_url}
	OriginUrl *string `field:"required" json:"originUrl" yaml:"originUrl"`
	// The configured backend type. Supported values are: `http`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.66.0/docs/resources/cdn_distribution#type CdnDistribution#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The configured origin request headers for the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.66.0/docs/resources/cdn_distribution#origin_request_headers CdnDistribution#origin_request_headers}
	OriginRequestHeaders *map[string]*string `field:"optional" json:"originRequestHeaders" yaml:"originRequestHeaders"`
}

