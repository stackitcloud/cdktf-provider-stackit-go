package routingtableroute


type RoutingTableRouteNextHop struct {
	// Type of the next hop. Possible values are: `blackhole`, `internet`, `ipv4`, `ipv6`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.71.0/docs/resources/routing_table_route#type RoutingTableRoute#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Either IPv4 or IPv6 (not set for blackhole and internet). Only IPv4 supported during experimental stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.71.0/docs/resources/routing_table_route#value RoutingTableRoute#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

