package datastackitvpngatewaystatus

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitVpnGatewayStatusConfig struct {
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
	// The server-generated UUID of the VPN gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/data-sources/vpn_gateway_status#gateway_id DataStackitVpnGatewayStatus#gateway_id}
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// STACKIT project ID associated with the VPN gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/data-sources/vpn_gateway_status#project_id DataStackitVpnGatewayStatus#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

