package applicationloadbalancer


type ApplicationLoadBalancerTargetPoolsTargets struct {
	// Private target IP, which must by unique within a target pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#ip ApplicationLoadBalancer#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// Target display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#display_name ApplicationLoadBalancer#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

