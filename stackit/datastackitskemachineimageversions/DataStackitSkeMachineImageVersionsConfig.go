package datastackitskemachineimageversions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitSkeMachineImageVersionsConfig struct {
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
	// Region override. If omitted, the provider’s region will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.80.0/docs/data-sources/ske_machine_image_versions#region DataStackitSkeMachineImageVersions#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Filter returned machine image versions by their state. Possible values are: `UNSPECIFIED`, `SUPPORTED`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.80.0/docs/data-sources/ske_machine_image_versions#version_state DataStackitSkeMachineImageVersions#version_state}
	VersionState *string `field:"optional" json:"versionState" yaml:"versionState"`
}

