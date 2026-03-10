package applicationloadbalancer


type ApplicationLoadBalancerTargetPoolsActiveHealthCheck struct {
	// Healthy threshold of the health checking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#healthy_threshold ApplicationLoadBalancer#healthy_threshold}
	HealthyThreshold *float64 `field:"required" json:"healthyThreshold" yaml:"healthyThreshold"`
	// Interval duration of health checking in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#interval ApplicationLoadBalancer#interval}
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// Interval duration threshold of the health checking in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#interval_jitter ApplicationLoadBalancer#interval_jitter}
	IntervalJitter *string `field:"required" json:"intervalJitter" yaml:"intervalJitter"`
	// Active health checking timeout duration in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#timeout ApplicationLoadBalancer#timeout}
	Timeout *string `field:"required" json:"timeout" yaml:"timeout"`
	// Unhealthy threshold of the health checking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#unhealthy_threshold ApplicationLoadBalancer#unhealthy_threshold}
	UnhealthyThreshold *float64 `field:"required" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
	// Options for the HTTP health checking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#http_health_checks ApplicationLoadBalancer#http_health_checks}
	HttpHealthChecks *ApplicationLoadBalancerTargetPoolsActiveHealthCheckHttpHealthChecks `field:"optional" json:"httpHealthChecks" yaml:"httpHealthChecks"`
}

