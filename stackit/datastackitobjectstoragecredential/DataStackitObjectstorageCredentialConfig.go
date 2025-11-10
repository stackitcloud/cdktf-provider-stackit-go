package datastackitobjectstoragecredential

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitObjectstorageCredentialConfig struct {
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
	// The credential ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.48.0/docs/data-sources/objectstorage_credential#credential_id DataStackitObjectstorageCredential#credential_id}
	CredentialId *string `field:"required" json:"credentialId" yaml:"credentialId"`
	// The credential group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.48.0/docs/data-sources/objectstorage_credential#credentials_group_id DataStackitObjectstorageCredential#credentials_group_id}
	CredentialsGroupId *string `field:"required" json:"credentialsGroupId" yaml:"credentialsGroupId"`
	// STACKIT Project ID to which the credential group is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.48.0/docs/data-sources/objectstorage_credential#project_id DataStackitObjectstorageCredential#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.48.0/docs/data-sources/objectstorage_credential#region DataStackitObjectstorageCredential#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

