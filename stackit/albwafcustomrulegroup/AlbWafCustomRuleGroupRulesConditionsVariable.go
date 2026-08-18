package albwafcustomrulegroup


type AlbWafCustomRuleGroupRulesConditionsVariable struct {
	// The targeted validation engine variable macro.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.112.0/docs/resources/alb_waf_custom_rule_group#type AlbWafCustomRuleGroup#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Optional key element context for map variables (e.g., matching a 'Host' header key).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.112.0/docs/resources/alb_waf_custom_rule_group#value AlbWafCustomRuleGroup#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

