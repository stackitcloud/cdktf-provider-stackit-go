package loadbalancer


type LoadbalancerTargetPoolsActiveHealthCheck struct {
	// Healthy threshold of the health checking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.20.2/docs/resources/loadbalancer#healthy_threshold Loadbalancer#healthy_threshold}
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// Interval duration of health checking in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.20.2/docs/resources/loadbalancer#interval Loadbalancer#interval}
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
	// Interval duration threshold of the health checking in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.20.2/docs/resources/loadbalancer#interval_jitter Loadbalancer#interval_jitter}
	IntervalJitter *string `field:"optional" json:"intervalJitter" yaml:"intervalJitter"`
	// Active health checking timeout duration in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.20.2/docs/resources/loadbalancer#timeout Loadbalancer#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	// Unhealthy threshold of the health checking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.20.2/docs/resources/loadbalancer#unhealthy_threshold Loadbalancer#unhealthy_threshold}
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

