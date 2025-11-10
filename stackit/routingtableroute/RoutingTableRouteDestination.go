package routingtableroute


type RoutingTableRouteDestination struct {
	// CIDRV type. Possible values are: `cidrv4`, `cidrv6`. Only `cidrv4` is supported during experimental stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.69.0/docs/resources/routing_table_route#type RoutingTableRoute#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// An CIDR string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.69.0/docs/resources/routing_table_route#value RoutingTableRoute#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

