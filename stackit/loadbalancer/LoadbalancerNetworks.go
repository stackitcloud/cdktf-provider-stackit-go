package loadbalancer


type LoadbalancerNetworks struct {
	// Openstack network ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.16.1/docs/resources/loadbalancer#network_id Loadbalancer#network_id}
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// The role defines how the load balancer is using the network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.16.1/docs/resources/loadbalancer#role Loadbalancer#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

