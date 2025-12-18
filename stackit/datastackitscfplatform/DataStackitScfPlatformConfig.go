package datastackitscfplatform

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitScfPlatformConfig struct {
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
	// The unique id of the platform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.74.0/docs/data-sources/scf_platform#platform_id DataStackitScfPlatform#platform_id}
	PlatformId *string `field:"required" json:"platformId" yaml:"platformId"`
	// The ID of the project associated with the platform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.74.0/docs/data-sources/scf_platform#project_id DataStackitScfPlatform#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The region where the platform is located. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.74.0/docs/data-sources/scf_platform#region DataStackitScfPlatform#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

