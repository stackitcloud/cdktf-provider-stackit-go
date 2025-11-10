package loadbalancer


type LoadbalancerNetworks struct {
	// Openstack network ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.64.0/docs/resources/loadbalancer#network_id Loadbalancer#network_id}
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// The role defines how the load balancer is using the network. Supported values are: `ROLE_UNSPECIFIED`, `ROLE_LISTENERS_AND_TARGETS`, `ROLE_LISTENERS`, `ROLE_TARGETS`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.64.0/docs/resources/loadbalancer#role Loadbalancer#role}
	Role *string `field:"required" json:"role" yaml:"role"`
}

