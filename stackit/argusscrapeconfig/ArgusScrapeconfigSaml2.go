package argusscrapeconfig


type ArgusScrapeconfigSaml2 struct {
	// Specifies if URL parameters are enabled. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.41.0/docs/resources/argus_scrapeconfig#enable_url_parameters ArgusScrapeconfig#enable_url_parameters}
	EnableUrlParameters interface{} `field:"optional" json:"enableUrlParameters" yaml:"enableUrlParameters"`
}

