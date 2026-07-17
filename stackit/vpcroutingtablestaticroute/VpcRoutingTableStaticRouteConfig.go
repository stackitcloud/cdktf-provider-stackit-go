package vpcroutingtablestaticroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcRoutingTableStaticRouteConfig struct {
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
	// The destination of the static route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_routing_table_static_route#destination VpcRoutingTableStaticRoute#destination}
	Destination *VpcRoutingTableStaticRouteDestination `field:"required" json:"destination" yaml:"destination"`
	// The nexthop of the static route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_routing_table_static_route#nexthop VpcRoutingTableStaticRoute#nexthop}
	Nexthop *VpcRoutingTableStaticRouteNexthop `field:"required" json:"nexthop" yaml:"nexthop"`
	// STACKIT Project ID to which the static route is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_routing_table_static_route#project_id VpcRoutingTableStaticRoute#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The routing table ID to which the static route is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_routing_table_static_route#routing_table_id VpcRoutingTableStaticRoute#routing_table_id}
	RoutingTableId *string `field:"required" json:"routingTableId" yaml:"routingTableId"`
	// The VPC ID to which the static route is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_routing_table_static_route#vpc_id VpcRoutingTableStaticRoute#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Labels are key-value string pairs which can be attached to a resource container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_routing_table_static_route#labels VpcRoutingTableStaticRoute#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The region of the static route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_routing_table_static_route#region VpcRoutingTableStaticRoute#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_routing_table_static_route#timeouts VpcRoutingTableStaticRoute#timeouts}.
	Timeouts *VpcRoutingTableStaticRouteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

