package loadbalancer


type LoadbalancerOptionsObservabilityMetrics struct {
	// Credentials reference for metrics. Not changeable after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.95.0/docs/resources/loadbalancer#credentials_ref Loadbalancer#credentials_ref}
	CredentialsRef *string `field:"optional" json:"credentialsRef" yaml:"credentialsRef"`
	// The ARGUS/Prometheus remote write Push URL to ship the metrics to. Not changeable after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.95.0/docs/resources/loadbalancer#push_url Loadbalancer#push_url}
	PushUrl *string `field:"optional" json:"pushUrl" yaml:"pushUrl"`
}

