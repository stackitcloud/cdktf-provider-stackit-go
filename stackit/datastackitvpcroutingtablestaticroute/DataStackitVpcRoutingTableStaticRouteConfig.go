package datastackitvpcroutingtablestaticroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitVpcRoutingTableStaticRouteConfig struct {
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
	// STACKIT Project ID to which the static route is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/data-sources/vpc_routing_table_static_route#project_id DataStackitVpcRoutingTableStaticRoute#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The static route ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/data-sources/vpc_routing_table_static_route#route_id DataStackitVpcRoutingTableStaticRoute#route_id}
	RouteId *string `field:"required" json:"routeId" yaml:"routeId"`
	// The routing table ID to which the static route is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/data-sources/vpc_routing_table_static_route#routing_table_id DataStackitVpcRoutingTableStaticRoute#routing_table_id}
	RoutingTableId *string `field:"required" json:"routingTableId" yaml:"routingTableId"`
	// The VPC ID to which the static route is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/data-sources/vpc_routing_table_static_route#vpc_id DataStackitVpcRoutingTableStaticRoute#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The region of the static route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/data-sources/vpc_routing_table_static_route#region DataStackitVpcRoutingTableStaticRoute#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/data-sources/vpc_routing_table_static_route#timeouts DataStackitVpcRoutingTableStaticRoute#timeouts}.
	Timeouts *DataStackitVpcRoutingTableStaticRouteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

