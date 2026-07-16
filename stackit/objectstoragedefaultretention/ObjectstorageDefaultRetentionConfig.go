package objectstoragedefaultretention

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObjectstorageDefaultRetentionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/objectstorage_default_retention#bucket_name ObjectstorageDefaultRetention#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The number retention period in days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/objectstorage_default_retention#days ObjectstorageDefaultRetention#days}
	Days *float64 `field:"required" json:"days" yaml:"days"`
	// The retention mode for default retention on a bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/objectstorage_default_retention#mode ObjectstorageDefaultRetention#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// STACKIT Project ID to which the default-retention is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/objectstorage_default_retention#project_id ObjectstorageDefaultRetention#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/objectstorage_default_retention#region ObjectstorageDefaultRetention#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

