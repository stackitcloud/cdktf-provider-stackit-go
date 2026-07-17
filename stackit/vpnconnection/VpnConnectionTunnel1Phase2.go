package vpnconnection


type VpnConnectionTunnel1Phase2 struct {
	// Encryption algorithms for Phase 2. Possible values are: `aes256`, `aes128gcm16`, `aes256gcm16`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpn_connection#encryption_algorithms VpnConnection#encryption_algorithms}
	EncryptionAlgorithms *[]*string `field:"required" json:"encryptionAlgorithms" yaml:"encryptionAlgorithms"`
	// Integrity algorithms for Phase 2. Possible values are: `sha1`, `sha2_256`, `sha2_384`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpn_connection#integrity_algorithms VpnConnection#integrity_algorithms}
	IntegrityAlgorithms *[]*string `field:"required" json:"integrityAlgorithms" yaml:"integrityAlgorithms"`
	// Diffie-Hellman groups for Phase 2. Possible values are: `modp1024`, `modp2048`, `ecp256`, `ecp384`, `modp2048s256`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpn_connection#dh_groups VpnConnection#dh_groups}
	DhGroups *[]*string `field:"optional" json:"dhGroups" yaml:"dhGroups"`
	// Action to perform on DPD timeout. Default: 'restart'. Possible values are: `clear`, `restart`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpn_connection#dpd_action VpnConnection#dpd_action}
	DpdAction *string `field:"optional" json:"dpdAction" yaml:"dpdAction"`
	// Time to schedule a Child SA re-keying in seconds. Range: 900-3600. Default: 3600.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpn_connection#rekey_time VpnConnection#rekey_time}
	RekeyTime *float64 `field:"optional" json:"rekeyTime" yaml:"rekeyTime"`
	// Action to perform after loading the connection configuration. Default: 'start'. Possible values are: `none`, `start`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpn_connection#start_action VpnConnection#start_action}
	StartAction *string `field:"optional" json:"startAction" yaml:"startAction"`
}

