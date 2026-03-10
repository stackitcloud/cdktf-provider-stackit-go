package applicationloadbalancer


type ApplicationLoadBalancerOptionsObservabilityLogs struct {
	// Credentials reference for logging.
	//
	// This reference is created via the observability create endpoint and the credential needs to contain the basic auth username and password for the logging solution the push URL points to. Then this enables monitoring via remote write for the Application Load Balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#credentials_ref ApplicationLoadBalancer#credentials_ref}
	CredentialsRef *string `field:"required" json:"credentialsRef" yaml:"credentialsRef"`
	// Credentials reference for logging.
	//
	// This reference is created via the observability create endpoint and the credential needs to contain the basic auth username and password for the logging solution the push URL points to. Then this enables monitoring via remote write for the Application Load Balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#push_url ApplicationLoadBalancer#push_url}
	PushUrl *string `field:"required" json:"pushUrl" yaml:"pushUrl"`
}

