package loadbalancer


type LoadbalancerListenersTcp struct {
	// Time after which an idle connection is closed.
	//
	// The default value is set to 300 seconds, and the maximum value is 3600 seconds. The format is a duration and the unit must be seconds. Example: 30s
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.77.0/docs/resources/loadbalancer#idle_timeout Loadbalancer#idle_timeout}
	IdleTimeout *string `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
}

