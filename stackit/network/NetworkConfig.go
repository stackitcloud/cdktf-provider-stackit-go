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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.18.1/docs/resources/network#name Network#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The nameservers of the network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.18.1/docs/resources/network#nameservers Network#nameservers}
	Nameservers *[]*string `field:"required" json:"nameservers" yaml:"nameservers"`
	// STACKIT project ID to which the network is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.18.1/docs/resources/network#project_id Network#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The IPv4 prefix length of the network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.18.1/docs/resources/network#ipv4_prefix_length Network#ipv4_prefix_length}
	Ipv4PrefixLength *float64 `field:"optional" json:"ipv4PrefixLength" yaml:"ipv4PrefixLength"`
}

