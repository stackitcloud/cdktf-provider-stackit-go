package cdndistribution


type CdnDistributionConfigTls struct {
	// If set to true, the distribution will accept connections using TLS 1.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/cdn_distribution#enable_tls_10 CdnDistribution#enable_tls_10}
	EnableTls10 interface{} `field:"optional" json:"enableTls10" yaml:"enableTls10"`
	// If set to true, the distribution will accept connections using TLS 1.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/cdn_distribution#enable_tls_11 CdnDistribution#enable_tls_11}
	EnableTls11 interface{} `field:"optional" json:"enableTls11" yaml:"enableTls11"`
}

