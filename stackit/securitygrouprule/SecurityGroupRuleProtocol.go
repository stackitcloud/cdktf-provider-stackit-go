package securitygrouprule


type SecurityGroupRuleProtocol struct {
	// The protocol name which the rule should match. Either `name` or `number` must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.3/docs/resources/security_group_rule#name SecurityGroupRule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The protocol number which the rule should match. Either `name` or `number` must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.3/docs/resources/security_group_rule#number SecurityGroupRule#number}
	Number *float64 `field:"optional" json:"number" yaml:"number"`
}

