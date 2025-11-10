package image

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImageConfig struct {
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
	// The disk format of the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.50.0/docs/resources/image#disk_format Image#disk_format}
	DiskFormat *string `field:"required" json:"diskFormat" yaml:"diskFormat"`
	// The filepath of the raw image file to be uploaded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.50.0/docs/resources/image#local_file_path Image#local_file_path}
	LocalFilePath *string `field:"required" json:"localFilePath" yaml:"localFilePath"`
	// The name of the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.50.0/docs/resources/image#name Image#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the image is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.50.0/docs/resources/image#project_id Image#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Properties to set hardware and scheduling settings for an image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.50.0/docs/resources/image#config Image#config}
	Config *ImageConfigA `field:"optional" json:"config" yaml:"config"`
	// Labels are key-value string pairs which can be attached to a resource container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.50.0/docs/resources/image#labels Image#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The minimum disk size of the image in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.50.0/docs/resources/image#min_disk_size Image#min_disk_size}
	MinDiskSize *float64 `field:"optional" json:"minDiskSize" yaml:"minDiskSize"`
	// The minimum RAM of the image in MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.50.0/docs/resources/image#min_ram Image#min_ram}
	MinRam *float64 `field:"optional" json:"minRam" yaml:"minRam"`
}

