package resourcemanagerfolder

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ResourcemanagerFolderConfig struct {
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
	// The name of the folder.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/resourcemanager_folder#name ResourcemanagerFolder#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Email address of the owner of the folder.
	//
	// This value is only considered during creation. Changing it afterwards will have no effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/resourcemanager_folder#owner_email ResourcemanagerFolder#owner_email}
	OwnerEmail *string `field:"required" json:"ownerEmail" yaml:"ownerEmail"`
	// Parent resource identifier. Both container ID (user-friendly) and UUID are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/resourcemanager_folder#parent_container_id ResourcemanagerFolder#parent_container_id}
	ParentContainerId *string `field:"required" json:"parentContainerId" yaml:"parentContainerId"`
	// Labels are key-value string pairs which can be attached to a resource container.
	//
	// A label key must match the regex [A-ZÄÜÖa-zäüöß0-9_-]{1,64}. A label value must match the regex ^$|[A-ZÄÜÖa-zäüöß0-9_-]{1,64}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/resources/resourcemanager_folder#labels ResourcemanagerFolder#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

