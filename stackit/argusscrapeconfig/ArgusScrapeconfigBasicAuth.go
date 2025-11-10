package argusscrapeconfig


type ArgusScrapeconfigBasicAuth struct {
	// Specifies basic auth password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.1/docs/resources/argus_scrapeconfig#password ArgusScrapeconfig#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Specifies basic auth username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.1/docs/resources/argus_scrapeconfig#username ArgusScrapeconfig#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

