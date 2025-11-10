package datastackitresourcemanagerproject

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitResourcemanagerProjectConfig struct {
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
	// Project container ID. Globally unique, user-friendly identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.27.0/docs/data-sources/resourcemanager_project#container_id DataStackitResourcemanagerProject#container_id}
	ContainerId *string `field:"optional" json:"containerId" yaml:"containerId"`
	// Email address of the owner of the project.
	//
	// This value is only considered during creation. Changing it afterwards will have no effect.
	//
	// !> The "owner_email" field has been deprecated in favor of the "members" field. Please use the "members" field to assign the owner role to a user, by setting the "role" field to `owner`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.27.0/docs/data-sources/resourcemanager_project#owner_email DataStackitResourcemanagerProject#owner_email}
	OwnerEmail *string `field:"optional" json:"ownerEmail" yaml:"ownerEmail"`
	// Project UUID identifier.
	//
	// This is the ID that can be used in most of the other resources to identify the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.27.0/docs/data-sources/resourcemanager_project#project_id DataStackitResourcemanagerProject#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

