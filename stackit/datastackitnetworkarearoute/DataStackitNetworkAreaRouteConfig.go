package datastackitnetworkarearoute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitNetworkAreaRouteConfig struct {
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
	// The network area ID to which the network area route is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.68.0/docs/data-sources/network_area_route#network_area_id DataStackitNetworkAreaRoute#network_area_id}
	NetworkAreaId *string `field:"required" json:"networkAreaId" yaml:"networkAreaId"`
	// The network area route ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.68.0/docs/data-sources/network_area_route#network_area_route_id DataStackitNetworkAreaRoute#network_area_route_id}
	NetworkAreaRouteId *string `field:"required" json:"networkAreaRouteId" yaml:"networkAreaRouteId"`
	// STACKIT organization ID to which the network area is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.68.0/docs/data-sources/network_area_route#organization_id DataStackitNetworkAreaRoute#organization_id}
	OrganizationId *string `field:"required" json:"organizationId" yaml:"organizationId"`
}

