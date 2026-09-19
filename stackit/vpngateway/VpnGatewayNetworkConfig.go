package vpngateway


type VpnGatewayNetworkConfig struct {
	// The IPv4 network prefix (CIDR notation) allocated for the VPN gateway.
	//
	// Must have a prefix length of /28 or larger. Cannot be changed after the gateway is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.116.0/docs/resources/vpn_gateway#predefined_network_prefix VpnGateway#predefined_network_prefix}
	PredefinedNetworkPrefix *string `field:"optional" json:"predefinedNetworkPrefix" yaml:"predefinedNetworkPrefix"`
	// Custom routing table ID for the VPN gateway. If omitted, a default routing table is assigned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.116.0/docs/resources/vpn_gateway#routing_table_id VpnGateway#routing_table_id}
	RoutingTableId *string `field:"optional" json:"routingTableId" yaml:"routingTableId"`
}

