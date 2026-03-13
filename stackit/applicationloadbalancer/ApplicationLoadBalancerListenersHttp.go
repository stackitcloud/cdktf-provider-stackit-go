package applicationloadbalancer


type ApplicationLoadBalancerListenersHttp struct {
	// Defines routing rules grouped by hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#hosts ApplicationLoadBalancer#hosts}
	Hosts interface{} `field:"required" json:"hosts" yaml:"hosts"`
}

