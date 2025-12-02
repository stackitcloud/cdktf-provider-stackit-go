package datastackitcdndistribution


type DataStackitCdnDistributionConfigA struct {
	// The configured countries where distribution of content is blocked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.72.0/docs/data-sources/cdn_distribution#blocked_countries DataStackitCdnDistribution#blocked_countries}
	BlockedCountries *[]*string `field:"optional" json:"blockedCountries" yaml:"blockedCountries"`
}

