package rabbitmqinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RabbitmqInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.71.0/docs/resources/rabbitmq_instance#name RabbitmqInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The selected plan name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.71.0/docs/resources/rabbitmq_instance#plan_name RabbitmqInstance#plan_name}
	PlanName *string `field:"required" json:"planName" yaml:"planName"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.71.0/docs/resources/rabbitmq_instance#project_id RabbitmqInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The service version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.71.0/docs/resources/rabbitmq_instance#version RabbitmqInstance#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// Configuration parameters.
	//
	// Please note that removing a previously configured field from your Terraform configuration won't replace its value in the API. To update a previously configured field, explicitly set a new value for it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.71.0/docs/resources/rabbitmq_instance#parameters RabbitmqInstance#parameters}
	Parameters *RabbitmqInstanceParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

