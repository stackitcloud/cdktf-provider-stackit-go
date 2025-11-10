package cdndistribution


type CdnDistributionConfigA struct {
	// The configured backend for the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.55.0/docs/resources/cdn_distribution#backend CdnDistribution#backend}
	Backend *CdnDistributionConfigBackend `field:"required" json:"backend" yaml:"backend"`
	// The configured regions where content will be hosted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.55.0/docs/resources/cdn_distribution#regions CdnDistribution#regions}
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
}

