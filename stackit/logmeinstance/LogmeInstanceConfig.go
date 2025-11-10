package logmeinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogmeInstanceConfig struct {
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
	// Instance name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.16.1/docs/resources/logme_instance#name LogmeInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The selected plan name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.16.1/docs/resources/logme_instance#plan_name LogmeInstance#plan_name}
	PlanName *string `field:"required" json:"planName" yaml:"planName"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.16.1/docs/resources/logme_instance#project_id LogmeInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The service version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.16.1/docs/resources/logme_instance#version LogmeInstance#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.16.1/docs/resources/logme_instance#parameters LogmeInstance#parameters}.
	Parameters *LogmeInstanceParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

