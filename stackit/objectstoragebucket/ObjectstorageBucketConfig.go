package objectstoragebucket

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObjectstorageBucketConfig struct {
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
	// The bucket name. It must be DNS conform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.16.0/docs/resources/objectstorage_bucket#name ObjectstorageBucket#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT Project ID to which the bucket is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.16.0/docs/resources/objectstorage_bucket#project_id ObjectstorageBucket#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

