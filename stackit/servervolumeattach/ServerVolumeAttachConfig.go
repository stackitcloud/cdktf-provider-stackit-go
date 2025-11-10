package servervolumeattach

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServerVolumeAttachConfig struct {
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
	// STACKIT project ID to which the volume attachment is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/server_volume_attach#project_id ServerVolumeAttach#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The server ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/server_volume_attach#server_id ServerVolumeAttach#server_id}
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// The volume ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/server_volume_attach#volume_id ServerVolumeAttach#volume_id}
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
}

