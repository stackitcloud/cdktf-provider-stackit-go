package server

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServerConfig struct {
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
	// Name of the type of the machine for the server. Possible values are documented in [Virtual machine flavors](https://docs.stackit.cloud/stackit/en/virtual-machine-flavors-75137231.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.59.0/docs/resources/server#machine_type Server#machine_type}
	MachineType *string `field:"required" json:"machineType" yaml:"machineType"`
	// The name of the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.59.0/docs/resources/server#name Server#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the server is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.59.0/docs/resources/server#project_id Server#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The affinity group the server is assigned to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.59.0/docs/resources/server#affinity_group Server#affinity_group}
	AffinityGroup *string `field:"optional" json:"affinityGroup" yaml:"affinityGroup"`
	// The availability zone of the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.59.0/docs/resources/server#availability_zone Server#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The boot volume for the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.59.0/docs/resources/server#boot_volume Server#boot_volume}
	BootVolume *ServerBootVolume `field:"optional" json:"bootVolume" yaml:"bootVolume"`
	// The desired status of the server resource. Supported values are: `active`, `inactive`, `deallocated`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.59.0/docs/resources/server#desired_status Server#desired_status}
	DesiredStatus *string `field:"optional" json:"desiredStatus" yaml:"desiredStatus"`
	// The image ID to be used for an ephemeral disk on the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.59.0/docs/resources/server#image_id Server#image_id}
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// The name of the keypair used during server creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.59.0/docs/resources/server#keypair_name Server#keypair_name}
	KeypairName *string `field:"optional" json:"keypairName" yaml:"keypairName"`
	// Labels are key-value string pairs which can be attached to a resource container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.59.0/docs/resources/server#labels Server#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The IDs of network interfaces which should be attached to the server. Updating it will recreate the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.59.0/docs/resources/server#network_interfaces Server#network_interfaces}
	NetworkInterfaces *[]*string `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// User data that is passed via cloud-init to the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.59.0/docs/resources/server#user_data Server#user_data}
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

