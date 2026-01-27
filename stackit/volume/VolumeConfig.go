package volume

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VolumeConfig struct {
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
	// The availability zone of the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.1/docs/resources/volume#availability_zone Volume#availability_zone}
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// STACKIT project ID to which the volume is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.1/docs/resources/volume#project_id Volume#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The description of the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.1/docs/resources/volume#description Volume#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Parameter to connect to a key-encryption-key within the STACKIT-KMS to create encrypted volumes.
	//
	// These parameters never leave the backend again. So these parameters are not present on imports or in the datasource. They live only in your Terraform state after creation of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.1/docs/resources/volume#encryption_parameters Volume#encryption_parameters}
	EncryptionParameters *VolumeEncryptionParameters `field:"optional" json:"encryptionParameters" yaml:"encryptionParameters"`
	// Labels are key-value string pairs which can be attached to a resource container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.1/docs/resources/volume#labels Volume#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.1/docs/resources/volume#name Volume#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The performance class of the volume. Possible values are documented in [Service plans BlockStorage](https://docs.stackit.cloud/products/storage/block-storage/basics/service-plans/#currently-available-service-plans-performance-classes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.1/docs/resources/volume#performance_class Volume#performance_class}
	PerformanceClass *string `field:"optional" json:"performanceClass" yaml:"performanceClass"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.1/docs/resources/volume#region Volume#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The size of the volume in GB.
	//
	// It can only be updated to a larger value than the current size. Either `size` or `source` must be provided
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.1/docs/resources/volume#size Volume#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The source of the volume.
	//
	// It can be either a volume, an image, a snapshot or a backup. Either `size` or `source` must be provided
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.1/docs/resources/volume#source Volume#source}
	Source *VolumeSource `field:"optional" json:"source" yaml:"source"`
}

