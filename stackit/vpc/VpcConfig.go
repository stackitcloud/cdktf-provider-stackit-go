package vpc

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcConfig struct {
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
	// The description of the VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpc#description Vpc#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpc#name Vpc#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the VPC is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpc#project_id Vpc#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Labels are key-value string pairs which can be attached to a resource container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpc#labels Vpc#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/vpc#timeouts Vpc#timeouts}.
	Timeouts *VpcTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

