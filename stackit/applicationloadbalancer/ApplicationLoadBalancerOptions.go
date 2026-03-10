package applicationloadbalancer


type ApplicationLoadBalancerOptions struct {
	// Use this option to limit the IP ranges that can use the Application Load Balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#access_control ApplicationLoadBalancer#access_control}
	AccessControl *ApplicationLoadBalancerOptionsAccessControl `field:"optional" json:"accessControl" yaml:"accessControl"`
	// This option automates the handling of the external IP address for an Application Load Balancer.
	//
	// If set to true a new IP address will be automatically created. It will also be automatically deleted when the Load Balancer is deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#ephemeral_address ApplicationLoadBalancer#ephemeral_address}
	EphemeralAddress interface{} `field:"optional" json:"ephemeralAddress" yaml:"ephemeralAddress"`
	// We offer Load Balancer observability via STACKIT Observability or external solutions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#observability ApplicationLoadBalancer#observability}
	Observability *ApplicationLoadBalancerOptionsObservability `field:"optional" json:"observability" yaml:"observability"`
	// Application Load Balancer is accessible only via a private network ip address. Not changeable after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#private_network_only ApplicationLoadBalancer#private_network_only}
	PrivateNetworkOnly interface{} `field:"optional" json:"privateNetworkOnly" yaml:"privateNetworkOnly"`
}

