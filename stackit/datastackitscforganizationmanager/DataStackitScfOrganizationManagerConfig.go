package datastackitscforganizationmanager

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitScfOrganizationManagerConfig struct {
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
	// The ID of the Cloud Foundry Organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.71.0/docs/data-sources/scf_organization_manager#org_id DataStackitScfOrganizationManager#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// The ID of the project associated with the organization of the organization manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.71.0/docs/data-sources/scf_organization_manager#project_id DataStackitScfOrganizationManager#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The region where the organization of the organization manager is located. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.71.0/docs/data-sources/scf_organization_manager#region DataStackitScfOrganizationManager#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

