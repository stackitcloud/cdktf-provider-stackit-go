package datastackitobjectstoragebucket

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitObjectstorageBucketConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.8.0/docs/data-sources/objectstorage_bucket#name DataStackitObjectstorageBucket#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT Project ID to which the bucket is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.8.0/docs/data-sources/objectstorage_bucket#project_id DataStackitObjectstorageBucket#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

