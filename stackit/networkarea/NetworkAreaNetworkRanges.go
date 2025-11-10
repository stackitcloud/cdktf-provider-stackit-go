package networkarea


type NetworkAreaNetworkRanges struct {
	// Classless Inter-Domain Routing (CIDR).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.2/docs/resources/network_area#prefix NetworkArea#prefix}
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
}

