package volumeautomation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VolumeAutomationConfig struct {
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
	// STACKIT Project ID to which the volume automation is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.118.0/docs/resources/volume_automation#project_id VolumeAutomation#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// ID of the automation template this volume automation is based on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.118.0/docs/resources/volume_automation#template_id VolumeAutomation#template_id}
	TemplateId *string `field:"required" json:"templateId" yaml:"templateId"`
	// The volume automation description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.118.0/docs/resources/volume_automation#description VolumeAutomation#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configuration input for the volume automation. See [API Docs](https://docs.api.stackit.cloud/documentation/automation-service/version/v1#tag/Volume-Automations/operation/CreateVolumeAutomation) for possible configuration options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.118.0/docs/resources/volume_automation#input VolumeAutomation#input}
	Input *string `field:"optional" json:"input" yaml:"input"`
	// The volume automation name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.118.0/docs/resources/volume_automation#name VolumeAutomation#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.118.0/docs/resources/volume_automation#region VolumeAutomation#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Triggers that determine when the automation runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.118.0/docs/resources/volume_automation#triggers VolumeAutomation#triggers}
	Triggers *VolumeAutomationTriggers `field:"optional" json:"triggers" yaml:"triggers"`
}

