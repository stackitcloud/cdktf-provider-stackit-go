package cdndistribution


type CdnDistributionConfigA struct {
	// The configured backend for the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.74.0/docs/resources/cdn_distribution#backend CdnDistribution#backend}
	Backend *CdnDistributionConfigBackend `field:"required" json:"backend" yaml:"backend"`
	// The configured regions where content will be hosted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.74.0/docs/resources/cdn_distribution#regions CdnDistribution#regions}
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
	// The configured countries where distribution of content is blocked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.74.0/docs/resources/cdn_distribution#blocked_countries CdnDistribution#blocked_countries}
	BlockedCountries *[]*string `field:"optional" json:"blockedCountries" yaml:"blockedCountries"`
	// Configuration for the Image Optimizer.
	//
	// This is a paid feature that automatically optimizes images to reduce their file size for faster delivery, leading to improved website performance and a better user experience.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.74.0/docs/resources/cdn_distribution#optimizer CdnDistribution#optimizer}
	Optimizer *CdnDistributionConfigOptimizer `field:"optional" json:"optimizer" yaml:"optimizer"`
}

