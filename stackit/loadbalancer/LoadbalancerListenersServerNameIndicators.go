package loadbalancer


type LoadbalancerListenersServerNameIndicators struct {
	// A domain name to match in order to pass TLS traffic to the target pool in the current listener.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.19.0/docs/resources/loadbalancer#name Loadbalancer#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

