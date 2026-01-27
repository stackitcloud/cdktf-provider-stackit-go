package securitygrouprule


type SecurityGroupRuleProtocol struct {
	// The protocol name which the rule should match.
	//
	// Either `name` or `number` must be provided. Possible values are: `ah`, `dccp`, `egp`, `esp`, `gre`, `icmp`, `igmp`, `ipip`, `ipv6-encap`, `ipv6-frag`, `ipv6-icmp`, `ipv6-nonxt`, `ipv6-opts`, `ipv6-route`, `ospf`, `pgm`, `rsvp`, `sctp`, `tcp`, `udp`, `udplite`, `vrrp`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.1/docs/resources/security_group_rule#name SecurityGroupRule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The protocol number which the rule should match. Either `name` or `number` must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.1/docs/resources/security_group_rule#number SecurityGroupRule#number}
	Number *float64 `field:"optional" json:"number" yaml:"number"`
}

