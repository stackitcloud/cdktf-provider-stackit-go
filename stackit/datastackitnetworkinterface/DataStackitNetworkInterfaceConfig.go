package datastackitnetworkinterface

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitNetworkInterfaceConfig struct {
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
	// The network ID to which the network interface is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.38.1/docs/data-sources/network_interface#network_id DataStackitNetworkInterface#network_id}
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// The network interface ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.38.1/docs/data-sources/network_interface#network_interface_id DataStackitNetworkInterface#network_interface_id}
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// STACKIT project ID to which the network interface is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.38.1/docs/data-sources/network_interface#project_id DataStackitNetworkInterface#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

