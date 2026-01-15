package networkarearegion


type NetworkAreaRegionIpv4 struct {
	// List of Network ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network_area_region#network_ranges NetworkAreaRegion#network_ranges}
	NetworkRanges interface{} `field:"required" json:"networkRanges" yaml:"networkRanges"`
	// IPv4 Classless Inter-Domain Routing (CIDR).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network_area_region#transfer_network NetworkAreaRegion#transfer_network}
	TransferNetwork *string `field:"required" json:"transferNetwork" yaml:"transferNetwork"`
	// List of DNS Servers/Nameservers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network_area_region#default_nameservers NetworkAreaRegion#default_nameservers}
	DefaultNameservers *[]*string `field:"optional" json:"defaultNameservers" yaml:"defaultNameservers"`
	// The default prefix length for networks in the network area.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network_area_region#default_prefix_length NetworkAreaRegion#default_prefix_length}
	DefaultPrefixLength *float64 `field:"optional" json:"defaultPrefixLength" yaml:"defaultPrefixLength"`
	// The maximal prefix length for networks in the network area.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network_area_region#max_prefix_length NetworkAreaRegion#max_prefix_length}
	MaxPrefixLength *float64 `field:"optional" json:"maxPrefixLength" yaml:"maxPrefixLength"`
	// The minimal prefix length for networks in the network area.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network_area_region#min_prefix_length NetworkAreaRegion#min_prefix_length}
	MinPrefixLength *float64 `field:"optional" json:"minPrefixLength" yaml:"minPrefixLength"`
}

