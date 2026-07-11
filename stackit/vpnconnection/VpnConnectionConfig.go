package vpnconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnConnectionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A user-friendly name for the connection.
	//
	// Must start and end with an alphanumeric character, may contain hyphens, and be 1-63 characters long.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/vpn_connection#display_name VpnConnection#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The UUID of the parent VPN gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/vpn_connection#gateway_id VpnConnection#gateway_id}
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// STACKIT project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/vpn_connection#project_id VpnConnection#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Configuration for the IPsec tunnel1.
	//
	// ~> Write-Only argument `pre_shared_key_wo` is available to use in place of `pre_shared_key`. Write-Only arguments are supported in HashiCorp Terraform 1.11.0 and later. [Learn more](https://developer.hashicorp.com/terraform/language/resources/ephemeral#write-only-arguments).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/vpn_connection#tunnel1 VpnConnection#tunnel1}
	Tunnel1 *VpnConnectionTunnel1 `field:"required" json:"tunnel1" yaml:"tunnel1"`
	// Configuration for the IPsec tunnel2.
	//
	// ~> Write-Only argument `pre_shared_key_wo` is available to use in place of `pre_shared_key`. Write-Only arguments are supported in HashiCorp Terraform 1.11.0 and later. [Learn more](https://developer.hashicorp.com/terraform/language/resources/ephemeral#write-only-arguments).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/vpn_connection#tunnel2 VpnConnection#tunnel2}
	Tunnel2 *VpnConnectionTunnel2 `field:"required" json:"tunnel2" yaml:"tunnel2"`
	// Whether this connection is enabled. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/vpn_connection#enabled VpnConnection#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Map of custom labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/vpn_connection#labels VpnConnection#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// List of local IPv4 CIDRs to route through this connection.
	//
	// Optional for route-based and BGP configurations (defaults to 0.0.0.0/0). Mandatory for policy-based.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/vpn_connection#local_subnets VpnConnection#local_subnets}
	LocalSubnets *[]*string `field:"optional" json:"localSubnets" yaml:"localSubnets"`
	// STACKIT region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/vpn_connection#region VpnConnection#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// List of remote IPv4 CIDRs accessible via this connection.
	//
	// Optional for route-based and BGP configurations (defaults to 0.0.0.0/0). Mandatory for policy-based.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/vpn_connection#remote_subnets VpnConnection#remote_subnets}
	RemoteSubnets *[]*string `field:"optional" json:"remoteSubnets" yaml:"remoteSubnets"`
	// List of static routes (IPv4 CIDRs) for route-based VPN. Mandatory for ROUTE_BASED gateways.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/vpn_connection#static_routes VpnConnection#static_routes}
	StaticRoutes *[]*string `field:"optional" json:"staticRoutes" yaml:"staticRoutes"`
}

