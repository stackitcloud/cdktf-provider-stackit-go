package datastackitroutingtableroutes

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitRoutingTableRoutesConfig struct {
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
	// The network area ID to which the routing table is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/data-sources/routing_table_routes#network_area_id DataStackitRoutingTableRoutes#network_area_id}
	NetworkAreaId *string `field:"required" json:"networkAreaId" yaml:"networkAreaId"`
	// STACKIT organization ID to which the routing table is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/data-sources/routing_table_routes#organization_id DataStackitRoutingTableRoutes#organization_id}
	OrganizationId *string `field:"required" json:"organizationId" yaml:"organizationId"`
	// The routing tables ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/data-sources/routing_table_routes#routing_table_id DataStackitRoutingTableRoutes#routing_table_id}
	RoutingTableId *string `field:"required" json:"routingTableId" yaml:"routingTableId"`
	// The datasource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/data-sources/routing_table_routes#region DataStackitRoutingTableRoutes#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

