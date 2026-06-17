package cdndistribution


type CdnDistributionConfigRedirectsRulesMatchers struct {
	// A list of glob patterns to match against the request path.
	//
	// At least one value is required. Examples: "/shop/*" or "* /img/*"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.99.0/docs/resources/cdn_distribution#values CdnDistribution#values}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Defines how multiple matchers within this rule are combined (ALL, ANY, NONE). Defaults to ANY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.99.0/docs/resources/cdn_distribution#value_match_condition CdnDistribution#value_match_condition}
	ValueMatchCondition *string `field:"optional" json:"valueMatchCondition" yaml:"valueMatchCondition"`
}

