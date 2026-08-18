package albwafcustomrulegroup


type AlbWafCustomRuleGroupRulesConditionsOperator struct {
	// The operational evaluation type definition macro.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.112.0/docs/resources/alb_waf_custom_rule_group#type AlbWafCustomRuleGroup#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The text or rule regex pattern arguments applied inside the operator execution loop.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.112.0/docs/resources/alb_waf_custom_rule_group#value AlbWafCustomRuleGroup#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

