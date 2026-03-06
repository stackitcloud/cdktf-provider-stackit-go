package applicationloadbalancer


type ApplicationLoadBalancerListenersHttpHostsRulesCookiePersistence struct {
	// The name of the cookie to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.85.0/docs/resources/application_load_balancer#name ApplicationLoadBalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// TTL specifies the time-to-live for the cookie.
	//
	// The default value is 0s, and it acts as a session cookie, expiring when the client session ends.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.85.0/docs/resources/application_load_balancer#ttl ApplicationLoadBalancer#ttl}
	Ttl *string `field:"required" json:"ttl" yaml:"ttl"`
}

