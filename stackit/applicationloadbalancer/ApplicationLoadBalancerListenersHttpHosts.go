package applicationloadbalancer


type ApplicationLoadBalancerListenersHttpHosts struct {
	// Hostname to match. Supports wildcards (e.g. *.example.com).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.85.0/docs/resources/application_load_balancer#host ApplicationLoadBalancer#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// Routing rules under the specified host, matched by path prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.85.0/docs/resources/application_load_balancer#rules ApplicationLoadBalancer#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

