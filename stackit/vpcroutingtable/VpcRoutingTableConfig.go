package vpcroutingtable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcRoutingTableConfig struct {
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
	// The name of the regional routing table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpc_routing_table#name VpcRoutingTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the regional routing table is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpc_routing_table#project_id VpcRoutingTable#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The vpc ID to which the regional routing table is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpc_routing_table#vpc_id VpcRoutingTable#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Description of the regional routing table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpc_routing_table#description VpcRoutingTable#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// This controls whether dynamic routes are propagated to this regional routing table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpc_routing_table#dynamic_routes VpcRoutingTable#dynamic_routes}
	DynamicRoutes interface{} `field:"optional" json:"dynamicRoutes" yaml:"dynamicRoutes"`
	// Labels are key-value string pairs which can be attached to a resource container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpc_routing_table#labels VpcRoutingTable#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpc_routing_table#region VpcRoutingTable#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// This allows installation of automatic system routes for connectivity between projects in the same VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpc_routing_table#system_routes VpcRoutingTable#system_routes}
	SystemRoutes interface{} `field:"optional" json:"systemRoutes" yaml:"systemRoutes"`
}

