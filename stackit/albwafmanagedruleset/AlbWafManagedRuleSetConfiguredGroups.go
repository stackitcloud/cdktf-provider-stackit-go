package albwafmanagedruleset


type AlbWafManagedRuleSetConfiguredGroups struct {
	// Rules of the rule group, keyed by rule ID (e.g. `911100`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.122.0/docs/resources/alb_waf_managed_rule_set#rules AlbWafManagedRuleSet#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

