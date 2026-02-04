package networkarearoute


type NetworkAreaRouteDestination struct {
	// CIDRV type. Possible values are: `cidrv4`, `cidrv6`. Only `cidrv4` is supported currently.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.80.0/docs/resources/network_area_route#type NetworkAreaRoute#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// An CIDR string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.80.0/docs/resources/network_area_route#value NetworkAreaRoute#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

