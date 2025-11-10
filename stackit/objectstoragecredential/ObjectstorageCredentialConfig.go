package objectstoragecredential

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObjectstorageCredentialConfig struct {
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
	// The credential group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.27.0/docs/resources/objectstorage_credential#credentials_group_id ObjectstorageCredential#credentials_group_id}
	CredentialsGroupId *string `field:"required" json:"credentialsGroupId" yaml:"credentialsGroupId"`
	// STACKIT Project ID to which the credential group is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.27.0/docs/resources/objectstorage_credential#project_id ObjectstorageCredential#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Expiration timestamp, in RFC339 format without fractional seconds. Example: "2025-01-01T00:00:00Z". If not set, the credential never expires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.27.0/docs/resources/objectstorage_credential#expiration_timestamp ObjectstorageCredential#expiration_timestamp}
	ExpirationTimestamp *string `field:"optional" json:"expirationTimestamp" yaml:"expirationTimestamp"`
}

