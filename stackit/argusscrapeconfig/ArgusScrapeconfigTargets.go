package argusscrapeconfig


type ArgusScrapeconfigTargets struct {
	// Specifies target URLs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.5.0/docs/resources/argus_scrapeconfig#urls ArgusScrapeconfig#urls}
	Urls *[]*string `field:"required" json:"urls" yaml:"urls"`
	// Specifies labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.5.0/docs/resources/argus_scrapeconfig#labels ArgusScrapeconfig#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

