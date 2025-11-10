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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.0/docs/resources/volume#availability_zone Volume#availability_zone}
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// STACKIT project ID to which the volume is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.0/docs/resources/volume#project_id Volume#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The description of the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.0/docs/resources/volume#description Volume#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Labels are key-value string pairs which can be attached to a resource container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.0/docs/resources/volume#labels Volume#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.0/docs/resources/volume#name Volume#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The performance class of the volume. Possible values are documented in [Service plans BlockStorage](https://docs.stackit.cloud/stackit/en/service-plans-blockstorage-75137974.html#ServiceplansBlockStorage-CurrentlyavailableServicePlans%28performanceclasses%29).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.0/docs/resources/volume#performance_class Volume#performance_class}
	PerformanceClass *string `field:"optional" json:"performanceClass" yaml:"performanceClass"`
	// The server ID of the server to which the volume is attached to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.0/docs/resources/volume#server_id Volume#server_id}
	ServerId *string `field:"optional" json:"serverId" yaml:"serverId"`
	// The size of the volume in GB.
	//
	// It can only be updated to a larger value than the current size. Either `size` or `source` must be provided
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.0/docs/resources/volume#size Volume#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The source of the volume.
	//
	// It can be either a volume, an image, a snapshot or a backup. Either `size` or `source` must be provided
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.0/docs/resources/volume#source Volume#source}
	Source *VolumeSource `field:"optional" json:"source" yaml:"source"`
}

