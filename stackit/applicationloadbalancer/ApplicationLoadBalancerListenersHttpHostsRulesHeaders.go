package applicationloadbalancer


type ApplicationLoadBalancerListenersHttpHostsRulesHeaders struct {
	// Header name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#name ApplicationLoadBalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Exact match for the header value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#exact_match ApplicationLoadBalancer#exact_match}
	ExactMatch *string `field:"optional" json:"exactMatch" yaml:"exactMatch"`
}

