package datastackitobjectstoragecredentialsgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitObjectstorageCredentialsGroupConfig struct {
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
	// Object Storage Project ID to which the credentials group is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.1/docs/data-sources/objectstorage_credentials_group#project_id DataStackitObjectstorageCredentialsGroup#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The credentials group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.1/docs/data-sources/objectstorage_credentials_group#credentials_group_id DataStackitObjectstorageCredentialsGroup#credentials_group_id}
	CredentialsGroupId *string `field:"optional" json:"credentialsGroupId" yaml:"credentialsGroupId"`
	// The credentials group's display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.1/docs/data-sources/objectstorage_credentials_group#name DataStackitObjectstorageCredentialsGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

