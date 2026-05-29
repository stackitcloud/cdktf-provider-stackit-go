package cdndistribution


type CdnDistributionConfigRedirectsRules struct {
	// A list of matchers that define when this rule should apply. At least one matcher is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#matchers CdnDistribution#matchers}
	Matchers interface{} `field:"required" json:"matchers" yaml:"matchers"`
	// The HTTP status code for the redirect. Must be one of 301, 302, 303, 307, or 308.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#status_code CdnDistribution#status_code}
	StatusCode *float64 `field:"required" json:"statusCode" yaml:"statusCode"`
	// The target URL to redirect to. Must be a valid URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#target_url CdnDistribution#target_url}
	TargetUrl *string `field:"required" json:"targetUrl" yaml:"targetUrl"`
	// An optional description for the redirect rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#description CdnDistribution#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A toggle to enable or disable the redirect rule. Default to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#enabled CdnDistribution#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Defines how multiple matchers within this rule are combined (ALL, ANY, NONE). Defaults to ANY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#rule_match_condition CdnDistribution#rule_match_condition}
	RuleMatchCondition *string `field:"optional" json:"ruleMatchCondition" yaml:"ruleMatchCondition"`
}

