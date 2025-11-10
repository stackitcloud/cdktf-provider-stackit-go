package objectstoragecredentialsgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObjectstorageCredentialsGroupConfig struct {
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
	// The credentials group's display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.42.0/docs/resources/objectstorage_credentials_group#name ObjectstorageCredentialsGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Project ID to which the credentials group is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.42.0/docs/resources/objectstorage_credentials_group#project_id ObjectstorageCredentialsGroup#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.42.0/docs/resources/objectstorage_credentials_group#region ObjectstorageCredentialsGroup#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

