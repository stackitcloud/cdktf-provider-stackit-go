package scforganization

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ScfOrganizationConfig struct {
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
	// The name of the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/scf_organization#name ScfOrganization#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the project associated with the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/scf_organization#project_id ScfOrganization#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The ID of the platform associated with the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/scf_organization#platform_id ScfOrganization#platform_id}
	PlatformId *string `field:"optional" json:"platformId" yaml:"platformId"`
	// The ID of the quota associated with the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/scf_organization#quota_id ScfOrganization#quota_id}
	QuotaId *string `field:"optional" json:"quotaId" yaml:"quotaId"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/scf_organization#region ScfOrganization#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// A boolean indicating whether the organization is suspended.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/scf_organization#suspended ScfOrganization#suspended}
	Suspended interface{} `field:"optional" json:"suspended" yaml:"suspended"`
}

