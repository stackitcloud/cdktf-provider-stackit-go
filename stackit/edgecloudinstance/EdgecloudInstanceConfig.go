package edgecloudinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EdgecloudInstanceConfig struct {
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
	// Display name shown for the Edge Cloud instance.
	//
	// Has to be a valid hostname, with a length between 4 and 8 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/edgecloud_instance#display_name EdgecloudInstance#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// STACKIT Edge Plan ID for the Edge Cloud instance, has to be the UUID of an existing plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/edgecloud_instance#plan_id EdgecloudInstance#plan_id}
	PlanId *string `field:"required" json:"planId" yaml:"planId"`
	// STACKIT project ID to which the Edge Cloud instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/edgecloud_instance#project_id EdgecloudInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Description for your STACKIT Edge Cloud instance. Max length is 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/edgecloud_instance#description EdgecloudInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// STACKIT region to use for the instance, providers default_region will be used if unset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/edgecloud_instance#region EdgecloudInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

