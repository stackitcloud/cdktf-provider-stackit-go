package albwafcustomrulegroup


type AlbWafCustomRuleGroupRulesConditions struct {
	// The comparison logic executed against the transformed variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/alb_waf_custom_rule_group#operator AlbWafCustomRuleGroup#operator}
	Operator *AlbWafCustomRuleGroupRulesConditionsOperator `field:"required" json:"operator" yaml:"operator"`
	// The part of the HTTP transaction to inspect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/alb_waf_custom_rule_group#variable AlbWafCustomRuleGroup#variable}
	Variable *AlbWafCustomRuleGroupRulesConditionsVariable `field:"required" json:"variable" yaml:"variable"`
	// Ordered normalization steps applied before the operator runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/alb_waf_custom_rule_group#transformations AlbWafCustomRuleGroup#transformations}
	Transformations *[]*string `field:"optional" json:"transformations" yaml:"transformations"`
}

