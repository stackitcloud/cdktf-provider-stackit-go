package dremioinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DremioInstanceConfig struct {
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
	// Dremio instance authentication settings. A change here triggers a Dremio restart and will incur downtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/dremio_instance#authentication DremioInstance#authentication}
	Authentication *DremioInstanceAuthentication `field:"required" json:"authentication" yaml:"authentication"`
	// The display name is a short name chosen by the user to identify the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/dremio_instance#display_name DremioInstance#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// STACKIT Project ID to which the resource is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/dremio_instance#project_id DremioInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The description is a longer text chosen by the user to provide more context for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/dremio_instance#description DremioInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The STACKIT region name the resource is located in. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/dremio_instance#region DremioInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/dremio_instance#timeouts DremioInstance#timeouts}.
	Timeouts *DremioInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

