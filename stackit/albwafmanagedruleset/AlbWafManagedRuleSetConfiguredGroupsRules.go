package albwafmanagedruleset


type AlbWafManagedRuleSetConfiguredGroupsRules struct {
	// The mode of the rule. Possible values are: `MODE_ENABLED`, `MODE_DISABLED`, `MODE_LOG_ONLY`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.122.0/docs/resources/alb_waf_managed_rule_set#mode AlbWafManagedRuleSet#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

