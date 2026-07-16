package vpcroutingtablestaticroute


type VpcRoutingTableStaticRouteNexthop struct {
	// Type of the nexthop. Possible values are: `blackhole`, `internet`, `ipv4`, `ipv6`. Currently ipv6 is unsupported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpc_routing_table_static_route#type VpcRoutingTableStaticRoute#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Value of the nexthop.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpc_routing_table_static_route#value VpcRoutingTableStaticRoute#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

