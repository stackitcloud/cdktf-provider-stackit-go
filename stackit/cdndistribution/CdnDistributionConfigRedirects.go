package cdndistribution


type CdnDistributionConfigRedirects struct {
	// A list of redirect rules. The order of rules matters for evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.99.0/docs/resources/cdn_distribution#rules CdnDistribution#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

