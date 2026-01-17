package networkarearoute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkAreaRouteConfig struct {
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
	// Destination of the route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/resources/network_area_route#destination NetworkAreaRoute#destination}
	Destination *NetworkAreaRouteDestination `field:"required" json:"destination" yaml:"destination"`
	// The network area ID to which the network area route is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/resources/network_area_route#network_area_id NetworkAreaRoute#network_area_id}
	NetworkAreaId *string `field:"required" json:"networkAreaId" yaml:"networkAreaId"`
	// Next hop destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/resources/network_area_route#next_hop NetworkAreaRoute#next_hop}
	NextHop *NetworkAreaRouteNextHop `field:"required" json:"nextHop" yaml:"nextHop"`
	// STACKIT organization ID to which the network area is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/resources/network_area_route#organization_id NetworkAreaRoute#organization_id}
	OrganizationId *string `field:"required" json:"organizationId" yaml:"organizationId"`
	// Labels are key-value string pairs which can be attached to a resource container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/resources/network_area_route#labels NetworkAreaRoute#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/resources/network_area_route#region NetworkAreaRoute#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

