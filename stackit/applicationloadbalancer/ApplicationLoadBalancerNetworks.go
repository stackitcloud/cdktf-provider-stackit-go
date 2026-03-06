package applicationloadbalancer


type ApplicationLoadBalancerNetworks struct {
	// STACKIT network ID the Application Load Balancer and/or targets are in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.85.0/docs/resources/application_load_balancer#network_id ApplicationLoadBalancer#network_id}
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// The role defines how the Application Load Balancer is using the network. Possible values are: `ROLE_UNSPECIFIED`, `ROLE_LISTENERS_AND_TARGETS`, `ROLE_LISTENERS`, `ROLE_TARGETS`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.85.0/docs/resources/application_load_balancer#role ApplicationLoadBalancer#role}
	Role *string `field:"required" json:"role" yaml:"role"`
}

