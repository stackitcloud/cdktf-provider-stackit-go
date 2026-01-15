package network

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkConfig struct {
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
	// The name of the network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network#name Network#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the network is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network#project_id Network#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The IPv4 gateway of a network.
	//
	// If not specified, the first IP of the network will be assigned as the gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network#ipv4_gateway Network#ipv4_gateway}
	Ipv4Gateway *string `field:"optional" json:"ipv4Gateway" yaml:"ipv4Gateway"`
	// The IPv4 nameservers of the network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network#ipv4_nameservers Network#ipv4_nameservers}
	Ipv4Nameservers *[]*string `field:"optional" json:"ipv4Nameservers" yaml:"ipv4Nameservers"`
	// The IPv4 prefix of the network (CIDR).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network#ipv4_prefix Network#ipv4_prefix}
	Ipv4Prefix *string `field:"optional" json:"ipv4Prefix" yaml:"ipv4Prefix"`
	// The IPv4 prefix length of the network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network#ipv4_prefix_length Network#ipv4_prefix_length}
	Ipv4PrefixLength *float64 `field:"optional" json:"ipv4PrefixLength" yaml:"ipv4PrefixLength"`
	// The IPv6 gateway of a network.
	//
	// If not specified, the first IP of the network will be assigned as the gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network#ipv6_gateway Network#ipv6_gateway}
	Ipv6Gateway *string `field:"optional" json:"ipv6Gateway" yaml:"ipv6Gateway"`
	// The IPv6 nameservers of the network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network#ipv6_nameservers Network#ipv6_nameservers}
	Ipv6Nameservers *[]*string `field:"optional" json:"ipv6Nameservers" yaml:"ipv6Nameservers"`
	// The IPv6 prefix of the network (CIDR).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network#ipv6_prefix Network#ipv6_prefix}
	Ipv6Prefix *string `field:"optional" json:"ipv6Prefix" yaml:"ipv6Prefix"`
	// The IPv6 prefix length of the network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network#ipv6_prefix_length Network#ipv6_prefix_length}
	Ipv6PrefixLength *float64 `field:"optional" json:"ipv6PrefixLength" yaml:"ipv6PrefixLength"`
	// Labels are key-value string pairs which can be attached to a resource container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network#labels Network#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The nameservers of the network.
	//
	// This field is deprecated and will be removed in January 2026, use `ipv4_nameservers` to configure the nameservers for IPv4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network#nameservers Network#nameservers}
	Nameservers *[]*string `field:"optional" json:"nameservers" yaml:"nameservers"`
	// If set to `true`, the network doesn't have a gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network#no_ipv4_gateway Network#no_ipv4_gateway}
	NoIpv4Gateway interface{} `field:"optional" json:"noIpv4Gateway" yaml:"noIpv4Gateway"`
	// If set to `true`, the network doesn't have a gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network#no_ipv6_gateway Network#no_ipv6_gateway}
	NoIpv6Gateway interface{} `field:"optional" json:"noIpv6Gateway" yaml:"noIpv6Gateway"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network#region Network#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// If set to `true`, the network is routed and therefore accessible from other networks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network#routed Network#routed}
	Routed interface{} `field:"optional" json:"routed" yaml:"routed"`
	// The ID of the routing table associated with the network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/network#routing_table_id Network#routing_table_id}
	RoutingTableId *string `field:"optional" json:"routingTableId" yaml:"routingTableId"`
}

