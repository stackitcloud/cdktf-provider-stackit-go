package vpnconnection


type VpnConnectionTunnel1Peering struct {
	// Local tunnel interface IPv4 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/vpn_connection#local_address VpnConnection#local_address}
	LocalAddress *string `field:"required" json:"localAddress" yaml:"localAddress"`
	// Remote tunnel interface IPv4 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/vpn_connection#remote_address VpnConnection#remote_address}
	RemoteAddress *string `field:"required" json:"remoteAddress" yaml:"remoteAddress"`
}

