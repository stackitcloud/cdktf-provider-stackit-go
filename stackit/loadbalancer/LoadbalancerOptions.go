package loadbalancer


type LoadbalancerOptions struct {
	// Load Balancer is accessible only from an IP address in this range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.71.0/docs/resources/loadbalancer#acl Loadbalancer#acl}
	Acl *[]*string `field:"optional" json:"acl" yaml:"acl"`
	// We offer Load Balancer metrics observability via ARGUS or external solutions. Not changeable after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.71.0/docs/resources/loadbalancer#observability Loadbalancer#observability}
	Observability *LoadbalancerOptionsObservability `field:"optional" json:"observability" yaml:"observability"`
	// If true, Load Balancer is accessible only via a private network IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.71.0/docs/resources/loadbalancer#private_network_only Loadbalancer#private_network_only}
	PrivateNetworkOnly interface{} `field:"optional" json:"privateNetworkOnly" yaml:"privateNetworkOnly"`
}

