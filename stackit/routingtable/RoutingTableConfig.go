package routingtable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RoutingTableConfig struct {
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
	// The name of the routing table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/routing_table#name RoutingTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The network area ID to which the routing table is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/routing_table#network_area_id RoutingTable#network_area_id}
	NetworkAreaId *string `field:"required" json:"networkAreaId" yaml:"networkAreaId"`
	// STACKIT organization ID to which the routing table is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/routing_table#organization_id RoutingTable#organization_id}
	OrganizationId *string `field:"required" json:"organizationId" yaml:"organizationId"`
	// Description of the routing table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/routing_table#description RoutingTable#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// This controls whether dynamic routes are propagated to this routing table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/routing_table#dynamic_routes RoutingTable#dynamic_routes}
	DynamicRoutes interface{} `field:"optional" json:"dynamicRoutes" yaml:"dynamicRoutes"`
	// Labels are key-value string pairs which can be attached to a resource container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/routing_table#labels RoutingTable#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/routing_table#region RoutingTable#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// This controls whether the routes for project-to-project communication are created automatically or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/routing_table#system_routes RoutingTable#system_routes}
	SystemRoutes interface{} `field:"optional" json:"systemRoutes" yaml:"systemRoutes"`
}

