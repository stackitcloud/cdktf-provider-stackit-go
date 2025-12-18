package networkarearegion


type NetworkAreaRegionIpv4NetworkRanges struct {
	// Classless Inter-Domain Routing (CIDR).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.75.0/docs/resources/network_area_region#prefix NetworkAreaRegion#prefix}
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
}

