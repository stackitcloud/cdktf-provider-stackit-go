package servernetworkinterfaceattach

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServerNetworkInterfaceAttachConfig struct {
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
	// The network interface ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/server_network_interface_attach#network_interface_id ServerNetworkInterfaceAttach#network_interface_id}
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// STACKIT project ID to which the network interface attachment is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/server_network_interface_attach#project_id ServerNetworkInterfaceAttach#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The server ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/server_network_interface_attach#server_id ServerNetworkInterfaceAttach#server_id}
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
}

