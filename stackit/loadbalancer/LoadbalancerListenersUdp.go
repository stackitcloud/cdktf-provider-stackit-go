package loadbalancer


type LoadbalancerListenersUdp struct {
	// Time after which an idle session is closed.
	//
	// The default value is set to 1 minute, and the maximum value is 2 minutes. The format is a duration and the unit must be seconds. Example: 30s
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.71.0/docs/resources/loadbalancer#idle_timeout Loadbalancer#idle_timeout}
	IdleTimeout *string `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
}

