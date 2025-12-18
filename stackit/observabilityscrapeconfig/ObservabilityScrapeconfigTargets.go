package observabilityscrapeconfig


type ObservabilityScrapeconfigTargets struct {
	// Specifies target URLs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.75.0/docs/resources/observability_scrapeconfig#urls ObservabilityScrapeconfig#urls}
	Urls *[]*string `field:"required" json:"urls" yaml:"urls"`
	// Specifies labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.75.0/docs/resources/observability_scrapeconfig#labels ObservabilityScrapeconfig#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

