package securitygrouprule


type SecurityGroupRulePortRange struct {
	// The maximum port number. Should be greater or equal to the minimum.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.37.0/docs/resources/security_group_rule#max SecurityGroupRule#max}
	Max *float64 `field:"required" json:"max" yaml:"max"`
	// The minimum port number. Should be less or equal to the maximum.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.37.0/docs/resources/security_group_rule#min SecurityGroupRule#min}
	Min *float64 `field:"required" json:"min" yaml:"min"`
}

