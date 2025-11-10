package publicip

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PublicIpConfig struct {
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
	// STACKIT project ID to which the public IP is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.65.0/docs/resources/public_ip#project_id PublicIp#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Labels are key-value string pairs which can be attached to a resource container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.65.0/docs/resources/public_ip#labels PublicIp#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Associates the public IP with a network interface or a virtual IP (ID).
	//
	// If you are using this resource with a Kubernetes Load Balancer or any other resource which associates a network interface implicitly, use the lifecycle `ignore_changes` property in this field to prevent unintentional removal of the network interface due to drift in the Terraform state
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.65.0/docs/resources/public_ip#network_interface_id PublicIp#network_interface_id}
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
}

