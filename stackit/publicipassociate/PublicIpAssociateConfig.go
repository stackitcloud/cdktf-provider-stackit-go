package publicipassociate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PublicIpAssociateConfig struct {
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
	// The ID of the network interface (or virtual IP) to which the public IP should be attached to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs/resources/public_ip_associate#network_interface_id PublicIpAssociate#network_interface_id}
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// STACKIT project ID to which the public IP is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs/resources/public_ip_associate#project_id PublicIpAssociate#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The public IP ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs/resources/public_ip_associate#public_ip_id PublicIpAssociate#public_ip_id}
	PublicIpId *string `field:"required" json:"publicIpId" yaml:"publicIpId"`
}

