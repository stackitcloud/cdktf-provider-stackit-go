package datastackitskekubernetesversions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitSkeKubernetesVersionsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/data-sources/ske_kubernetes_versions#region DataStackitSkeKubernetesVersions#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// If specified, only returns Kubernetes versions with this version state. Possible values are: `UNSPECIFIED`, `SUPPORTED`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/data-sources/ske_kubernetes_versions#version_state DataStackitSkeKubernetesVersions#version_state}
	VersionState *string `field:"optional" json:"versionState" yaml:"versionState"`
}

