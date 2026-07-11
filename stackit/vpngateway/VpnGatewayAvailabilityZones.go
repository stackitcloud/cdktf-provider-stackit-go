package vpngateway


type VpnGatewayAvailabilityZones struct {
	// Availability zone for tunnel 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/vpn_gateway#tunnel1 VpnGateway#tunnel1}
	Tunnel1 *string `field:"required" json:"tunnel1" yaml:"tunnel1"`
	// Availability zone for tunnel 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/vpn_gateway#tunnel2 VpnGateway#tunnel2}
	Tunnel2 *string `field:"required" json:"tunnel2" yaml:"tunnel2"`
}

