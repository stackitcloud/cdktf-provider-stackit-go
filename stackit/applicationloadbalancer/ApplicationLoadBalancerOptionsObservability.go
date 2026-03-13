package applicationloadbalancer


type ApplicationLoadBalancerOptionsObservability struct {
	// Observability logs configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#logs ApplicationLoadBalancer#logs}
	Logs *ApplicationLoadBalancerOptionsObservabilityLogs `field:"optional" json:"logs" yaml:"logs"`
	// Observability metrics configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#metrics ApplicationLoadBalancer#metrics}
	Metrics *ApplicationLoadBalancerOptionsObservabilityMetrics `field:"optional" json:"metrics" yaml:"metrics"`
}

