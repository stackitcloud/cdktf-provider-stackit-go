package networkinterface

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkInterfaceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.67.0/docs/resources/network_interface#network_id NetworkInterface#network_id}
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// STACKIT project ID to which the network is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.67.0/docs/resources/network_interface#project_id NetworkInterface#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The list of CIDR (Classless Inter-Domain Routing) notations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.67.0/docs/resources/network_interface#allowed_addresses NetworkInterface#allowed_addresses}
	AllowedAddresses *[]*string `field:"optional" json:"allowedAddresses" yaml:"allowedAddresses"`
	// The IPv4 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.67.0/docs/resources/network_interface#ipv4 NetworkInterface#ipv4}
	Ipv4 *string `field:"optional" json:"ipv4" yaml:"ipv4"`
	// Labels are key-value string pairs which can be attached to a network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.67.0/docs/resources/network_interface#labels NetworkInterface#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.67.0/docs/resources/network_interface#name NetworkInterface#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Network Interface Security. If set to false, then no security groups will apply to this network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.67.0/docs/resources/network_interface#security NetworkInterface#security}
	Security interface{} `field:"optional" json:"security" yaml:"security"`
	// The list of security group UUIDs.
	//
	// If security is set to false, setting this field will lead to an error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.67.0/docs/resources/network_interface#security_group_ids NetworkInterface#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

