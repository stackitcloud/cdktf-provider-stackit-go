package vpnconnection


type VpnConnectionTunnel1 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpn_connection#phase1 VpnConnection#phase1}.
	Phase1 *VpnConnectionTunnel1Phase1 `field:"required" json:"phase1" yaml:"phase1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpn_connection#phase2 VpnConnection#phase2}.
	Phase2 *VpnConnectionTunnel1Phase2 `field:"required" json:"phase2" yaml:"phase2"`
	// Remote IPv4 address for the tunnel endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpn_connection#remote_address VpnConnection#remote_address}
	RemoteAddress *string `field:"required" json:"remoteAddress" yaml:"remoteAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpn_connection#bgp VpnConnection#bgp}.
	Bgp *VpnConnectionTunnel1Bgp `field:"optional" json:"bgp" yaml:"bgp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpn_connection#peering VpnConnection#peering}.
	Peering *VpnConnectionTunnel1Peering `field:"optional" json:"peering" yaml:"peering"`
	// Pre-shared key for the IPsec tunnel. Minimum 20 characters. Write-only argument `pre_shared_key_wo` should be preferred.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpn_connection#pre_shared_key VpnConnection#pre_shared_key}
	PreSharedKey *string `field:"optional" json:"preSharedKey" yaml:"preSharedKey"`
	// Pre-shared key for the IPsec tunnel.
	//
	// Minimum 20 characters. Write-only - never stored in state and never returned by the API. To rotate the key, update this value AND increment pre_shared_key_wo_version. Changing this field alone will NOT trigger an update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpn_connection#pre_shared_key_wo VpnConnection#pre_shared_key_wo}
	PreSharedKeyWo *string `field:"optional" json:"preSharedKeyWo" yaml:"preSharedKeyWo"`
	// User-managed rotation counter for the pre-shared key.
	//
	// Must be incremented every time pre_shared_key_wo is changed. Terraform diffs this field to detect key rotations - changing pre_shared_key_wo alone will NOT trigger an update because it is write-only and never stored in state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpn_connection#pre_shared_key_wo_version VpnConnection#pre_shared_key_wo_version}
	PreSharedKeyWoVersion *float64 `field:"optional" json:"preSharedKeyWoVersion" yaml:"preSharedKeyWoVersion"`
}

