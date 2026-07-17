package datastackitobjectstoragedefaultretention

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitObjectstorageDefaultRetentionConfig struct {
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
	// The associated bucket's name. It must be DNS conform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/data-sources/objectstorage_default_retention#bucket_name DataStackitObjectstorageDefaultRetention#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// STACKIT Project ID to which the default-retention is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/data-sources/objectstorage_default_retention#project_id DataStackitObjectstorageDefaultRetention#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/data-sources/objectstorage_default_retention#region DataStackitObjectstorageDefaultRetention#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

