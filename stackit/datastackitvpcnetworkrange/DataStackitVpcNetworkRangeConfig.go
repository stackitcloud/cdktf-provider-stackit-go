package datastackitvpcnetworkrange

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitVpcNetworkRangeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/data-sources/vpc_network_range#network_range_id DataStackitVpcNetworkRange#network_range_id}.
	NetworkRangeId *string `field:"required" json:"networkRangeId" yaml:"networkRangeId"`
	// STACKIT project ID to which the VPC network range is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/data-sources/vpc_network_range#project_id DataStackitVpcNetworkRange#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The VPC ID to which the VPC network range is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/data-sources/vpc_network_range#vpc_id DataStackitVpcNetworkRange#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/data-sources/vpc_network_range#region DataStackitVpcNetworkRange#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/data-sources/vpc_network_range#timeouts DataStackitVpcNetworkRange#timeouts}.
	Timeouts *DataStackitVpcNetworkRangeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

