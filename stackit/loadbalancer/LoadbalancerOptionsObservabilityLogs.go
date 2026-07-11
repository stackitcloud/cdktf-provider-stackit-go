package loadbalancer


type LoadbalancerOptionsObservabilityLogs struct {
	// Credentials reference for logs. Not changeable after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/loadbalancer#credentials_ref Loadbalancer#credentials_ref}
	CredentialsRef *string `field:"optional" json:"credentialsRef" yaml:"credentialsRef"`
	// The ARGUS/Loki remote write Push URL to ship the logs to. Not changeable after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/loadbalancer#push_url Loadbalancer#push_url}
	PushUrl *string `field:"optional" json:"pushUrl" yaml:"pushUrl"`
}

