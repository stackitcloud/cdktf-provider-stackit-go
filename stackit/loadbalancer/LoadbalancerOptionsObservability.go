package loadbalancer


type LoadbalancerOptionsObservability struct {
	// Observability logs configuration. Not changeable after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.56.0/docs/resources/loadbalancer#logs Loadbalancer#logs}
	Logs *LoadbalancerOptionsObservabilityLogs `field:"optional" json:"logs" yaml:"logs"`
	// Observability metrics configuration. Not changeable after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.56.0/docs/resources/loadbalancer#metrics Loadbalancer#metrics}
	Metrics *LoadbalancerOptionsObservabilityMetrics `field:"optional" json:"metrics" yaml:"metrics"`
}

