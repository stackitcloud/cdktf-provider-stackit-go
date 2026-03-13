package applicationloadbalancer


type ApplicationLoadBalancerListenersHttpHostsRulesPath struct {
	// Exact path match.
	//
	// Only a request path exactly equal to the value will match, e.g. '/foo' matches only '/foo', not '/foo/bar' or '/foobar'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#exact_match ApplicationLoadBalancer#exact_match}
	ExactMatch *string `field:"optional" json:"exactMatch" yaml:"exactMatch"`
	// Prefix path match. Only matches on full segment boundaries, e.g. '/foo' matches '/foo' and '/foo/bar' but NOT '/foobar'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#prefix ApplicationLoadBalancer#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

