package networkarearoute


type NetworkAreaRouteNextHop struct {
	// Type of the next hop. Possible values are: `blackhole`, `internet`, `ipv4`, `ipv6`. Only `ipv4` supported currently.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network_area_route#type NetworkAreaRoute#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Either IPv4 or IPv6 (not set for blackhole and internet). Only IPv4 supported currently.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network_area_route#value NetworkAreaRoute#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

