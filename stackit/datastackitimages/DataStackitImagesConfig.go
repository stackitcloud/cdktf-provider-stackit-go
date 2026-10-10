package datastackitimages

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitImagesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.119.0/docs/data-sources/images#project_id DataStackitImages#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// API-side image filtering options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.119.0/docs/data-sources/images#filter DataStackitImages#filter}
	Filter *DataStackitImagesFilter `field:"optional" json:"filter" yaml:"filter"`
	// Region override. Uses the provider region when omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.119.0/docs/data-sources/images#region DataStackitImages#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.119.0/docs/data-sources/images#timeouts DataStackitImages#timeouts}.
	Timeouts *DataStackitImagesTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

