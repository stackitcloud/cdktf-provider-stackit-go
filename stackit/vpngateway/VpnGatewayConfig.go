package vpngateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnGatewayConfig struct {
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
	// Availability zones for the two tunnel endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/vpn_gateway#availability_zones VpnGateway#availability_zones}
	AvailabilityZones *VpnGatewayAvailabilityZones `field:"required" json:"availabilityZones" yaml:"availabilityZones"`
	// A user-friendly name for the VPN gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/vpn_gateway#display_name VpnGateway#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The service plan identifier (e.g. `p500`). For guidance on finding available plans, see [List available service plans](https://docs.stackit.cloud/products/network/connectivity-hybrid-multi-cloud/vpn/getting-started/gateway-create/#list-available-service-plans).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/vpn_gateway#plan_id VpnGateway#plan_id}
	PlanId *string `field:"required" json:"planId" yaml:"planId"`
	// STACKIT project ID associated with the VPN gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/vpn_gateway#project_id VpnGateway#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Routing architecture. Possible values are: `POLICY_BASED`, `ROUTE_BASED`, `BGP_ROUTE_BASED`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/vpn_gateway#routing_type VpnGateway#routing_type}
	RoutingType *string `field:"required" json:"routingType" yaml:"routingType"`
	// BGP configuration. Only applicable when routing_type is BGP_ROUTE_BASED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/vpn_gateway#bgp VpnGateway#bgp}
	Bgp *VpnGatewayBgp `field:"optional" json:"bgp" yaml:"bgp"`
	// Map of custom labels (key-value string pairs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/vpn_gateway#labels VpnGateway#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// STACKIT region name the resource is located in. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/vpn_gateway#region VpnGateway#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

