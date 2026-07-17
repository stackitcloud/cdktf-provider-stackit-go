package vpnconnection


type VpnConnectionTunnel2Bgp struct {
	// Remote ASN for BGP peering (private ASN range, 64512-4294967294).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpn_connection#remote_asn VpnConnection#remote_asn}
	RemoteAsn *float64 `field:"required" json:"remoteAsn" yaml:"remoteAsn"`
}

