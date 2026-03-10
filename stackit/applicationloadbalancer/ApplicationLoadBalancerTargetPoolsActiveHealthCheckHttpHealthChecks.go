package applicationloadbalancer


type ApplicationLoadBalancerTargetPoolsActiveHealthCheckHttpHealthChecks struct {
	// List of HTTP status codes that indicate a healthy response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#ok_status ApplicationLoadBalancer#ok_status}
	OkStatus *[]*string `field:"required" json:"okStatus" yaml:"okStatus"`
	// Path to send the health check request to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#path ApplicationLoadBalancer#path}
	Path *string `field:"required" json:"path" yaml:"path"`
}

