package datastackitvolumeautomation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitVolumeAutomationConfig struct {
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
	// ID of the volume automation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.119.0/docs/data-sources/volume_automation#automation_id DataStackitVolumeAutomation#automation_id}
	AutomationId *string `field:"required" json:"automationId" yaml:"automationId"`
	// STACKIT Project ID to which the volume automation is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.119.0/docs/data-sources/volume_automation#project_id DataStackitVolumeAutomation#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.119.0/docs/data-sources/volume_automation#region DataStackitVolumeAutomation#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

