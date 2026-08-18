package albwafcustomrulegroup


type AlbWafCustomRuleGroupRulesBehavior struct {
	// The protective stance action. ACTION_DENY forces a 403 status response code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.112.0/docs/resources/alb_waf_custom_rule_group#action AlbWafCustomRuleGroup#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Determines whether an entry should be generated in the security ledger upon a rule hit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.112.0/docs/resources/alb_waf_custom_rule_group#log AlbWafCustomRuleGroup#log}
	Log interface{} `field:"optional" json:"log" yaml:"log"`
	// Custom notification message string mapped to underlying logdata contexts. Required if log is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.112.0/docs/resources/alb_waf_custom_rule_group#log_msg AlbWafCustomRuleGroup#log_msg}
	LogMsg *string `field:"optional" json:"logMsg" yaml:"logMsg"`
}

