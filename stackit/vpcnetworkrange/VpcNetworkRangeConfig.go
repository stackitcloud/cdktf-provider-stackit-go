package vpcnetworkrange

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcNetworkRangeConfig struct {
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
	// Description of the VPC network range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_network_range#description VpcNetworkRange#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// IP version of the VPC network range. Possible values are: `ipv4`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_network_range#ip_version VpcNetworkRange#ip_version}
	IpVersion *string `field:"required" json:"ipVersion" yaml:"ipVersion"`
	// The IP prefix for the VPC network range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_network_range#prefix VpcNetworkRange#prefix}
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// STACKIT project ID to which the VPC network range is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_network_range#project_id VpcNetworkRange#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The VPC ID to which the VPC network range is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_network_range#vpc_id VpcNetworkRange#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Default prefix length for the VPC network range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_network_range#default_prefix_length VpcNetworkRange#default_prefix_length}
	DefaultPrefixLength *float64 `field:"optional" json:"defaultPrefixLength" yaml:"defaultPrefixLength"`
	// Labels are key-value string pairs which can be attached to a resource container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_network_range#labels VpcNetworkRange#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Maximum prefix length for the VPC network range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_network_range#max_prefix_length VpcNetworkRange#max_prefix_length}
	MaxPrefixLength *float64 `field:"optional" json:"maxPrefixLength" yaml:"maxPrefixLength"`
	// Minimum prefix length for the VPC network range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_network_range#min_prefix_length VpcNetworkRange#min_prefix_length}
	MinPrefixLength *float64 `field:"optional" json:"minPrefixLength" yaml:"minPrefixLength"`
	// List of nameservers for the VPC network range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_network_range#nameservers VpcNetworkRange#nameservers}
	Nameservers *[]*string `field:"optional" json:"nameservers" yaml:"nameservers"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_network_range#region VpcNetworkRange#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/vpc_network_range#timeouts VpcNetworkRange#timeouts}.
	Timeouts *VpcNetworkRangeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

