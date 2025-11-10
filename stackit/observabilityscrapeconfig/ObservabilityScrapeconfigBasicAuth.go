package observabilityscrapeconfig


type ObservabilityScrapeconfigBasicAuth struct {
	// Specifies basic auth password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.41.0/docs/resources/observability_scrapeconfig#password ObservabilityScrapeconfig#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Specifies basic auth username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.41.0/docs/resources/observability_scrapeconfig#username ObservabilityScrapeconfig#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

