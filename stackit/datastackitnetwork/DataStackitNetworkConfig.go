package datastackitnetwork

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitNetworkConfig struct {
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
	// The network ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.0/docs/data-sources/network#network_id DataStackitNetwork#network_id}
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// STACKIT project ID to which the network is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.0/docs/data-sources/network#project_id DataStackitNetwork#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Can only be used when experimental "network" is set.
	//
	// This is likely going to undergo significant changes or be removed in the future.
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.0/docs/data-sources/network#region DataStackitNetwork#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

