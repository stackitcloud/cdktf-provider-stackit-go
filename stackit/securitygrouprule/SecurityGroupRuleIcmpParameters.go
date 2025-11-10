package securitygrouprule


type SecurityGroupRuleIcmpParameters struct {
	// ICMP code. Can be set if the protocol is ICMP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.2/docs/resources/security_group_rule#code SecurityGroupRule#code}
	Code *float64 `field:"required" json:"code" yaml:"code"`
	// ICMP type. Can be set if the protocol is ICMP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.2/docs/resources/security_group_rule#type SecurityGroupRule#type}
	Type *float64 `field:"required" json:"type" yaml:"type"`
}

