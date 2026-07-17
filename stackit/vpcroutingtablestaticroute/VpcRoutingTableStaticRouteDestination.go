package vpcroutingtablestaticroute


type VpcRoutingTableStaticRouteDestination struct {
	// CIDR type. Possible values are: `cidrv4`, `cidrv6`. Currently cidrv6 is unsupported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_routing_table_static_route#type VpcRoutingTableStaticRoute#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// CIDR value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_routing_table_static_route#value VpcRoutingTableStaticRoute#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

