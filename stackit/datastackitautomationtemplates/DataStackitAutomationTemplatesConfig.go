package datastackitautomationtemplates

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitAutomationTemplatesConfig struct {
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
	// STACKIT project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/data-sources/automation_templates#project_id DataStackitAutomationTemplates#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Automation templates data source region. If undefined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/data-sources/automation_templates#region DataStackitAutomationTemplates#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

