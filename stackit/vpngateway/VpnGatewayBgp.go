package vpngateway


type VpnGatewayBgp struct {
	// Local ASN for BGP (private ASN range, 64512-4294967294).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpn_gateway#local_asn VpnGateway#local_asn}
	LocalAsn *float64 `field:"required" json:"localAsn" yaml:"localAsn"`
	// List of IPv4 CIDRs to advertise via BGP. If omitted, SNA network ranges are advertised.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpn_gateway#override_advertised_routes VpnGateway#override_advertised_routes}
	OverrideAdvertisedRoutes *[]*string `field:"optional" json:"overrideAdvertisedRoutes" yaml:"overrideAdvertisedRoutes"`
}

