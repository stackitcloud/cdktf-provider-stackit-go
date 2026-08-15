package albwafcustomrulegroup


type AlbWafCustomRuleGroupRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/alb_waf_custom_rule_group#behavior AlbWafCustomRuleGroup#behavior}.
	Behavior *AlbWafCustomRuleGroupRulesBehavior `field:"required" json:"behavior" yaml:"behavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/alb_waf_custom_rule_group#conditions AlbWafCustomRuleGroup#conditions}.
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// A clear description explaining the threat vector or criteria addressed by this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/alb_waf_custom_rule_group#description AlbWafCustomRuleGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

