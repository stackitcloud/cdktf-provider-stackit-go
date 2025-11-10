package resourcemanagerproject

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ResourcemanagerProjectConfig struct {
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
	// Project name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.36.0/docs/resources/resourcemanager_project#name ResourcemanagerProject#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Email address of the owner of the project.
	//
	// This value is only considered during creation. Changing it afterwards will have no effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.36.0/docs/resources/resourcemanager_project#owner_email ResourcemanagerProject#owner_email}
	OwnerEmail *string `field:"required" json:"ownerEmail" yaml:"ownerEmail"`
	// Parent resource identifier. Both container ID (user-friendly) and UUID are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.36.0/docs/resources/resourcemanager_project#parent_container_id ResourcemanagerProject#parent_container_id}
	ParentContainerId *string `field:"required" json:"parentContainerId" yaml:"parentContainerId"`
	// Labels are key-value string pairs which can be attached to a resource container.
	//
	// A label key must match the regex [A-ZÄÜÖa-zäüöß0-9_-]{1,64}. A label value must match the regex ^$|[A-ZÄÜÖa-zäüöß0-9_-]{1,64}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.36.0/docs/resources/resourcemanager_project#labels ResourcemanagerProject#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The members assigned to the project.
	//
	// At least one subject needs to be a user, and not a client or service account. This value is only considered during creation. Changing it afterwards will have no effect.
	//
	// !> The "members" field has been deprecated in favor of the "owner_email" field. Please use the "owner_email" field to assign the owner role to a user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.36.0/docs/resources/resourcemanager_project#members ResourcemanagerProject#members}
	Members interface{} `field:"optional" json:"members" yaml:"members"`
}

