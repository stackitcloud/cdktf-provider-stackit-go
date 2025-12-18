package datastackitimagev2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitImageV2Config struct {
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
	// STACKIT project ID to which the image is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.74.0/docs/data-sources/image_v2#project_id DataStackitImageV2#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Additional filtering options based on image properties. Can be used independently or in conjunction with `name` or `name_regex`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.74.0/docs/data-sources/image_v2#filter DataStackitImageV2#filter}
	Filter *DataStackitImageV2Filter `field:"optional" json:"filter" yaml:"filter"`
	// Image ID to fetch directly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.74.0/docs/data-sources/image_v2#image_id DataStackitImageV2#image_id}
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// Exact image name to match.
	//
	// Optionally applies a `filter` block to further refine results in case multiple images share the same name. The first match is returned, optionally sorted by name in ascending order. Cannot be used together with `name_regex`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.74.0/docs/data-sources/image_v2#name DataStackitImageV2#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Regular expression to match against image names.
	//
	// Optionally applies a `filter` block to narrow down results when multiple image names match the regex. The first match is returned, optionally sorted by name in ascending order. Cannot be used together with `name`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.74.0/docs/data-sources/image_v2#name_regex DataStackitImageV2#name_regex}
	NameRegex *string `field:"optional" json:"nameRegex" yaml:"nameRegex"`
	// If set to `true`, images are sorted in ascending lexicographical order by image name (such as `Ubuntu 18.04`, `Ubuntu 20.04`, `Ubuntu 22.04`) before selecting the first match. Defaults to `false` (descending such as `Ubuntu 22.04`, `Ubuntu 20.04`, `Ubuntu 18.04`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.74.0/docs/data-sources/image_v2#sort_ascending DataStackitImageV2#sort_ascending}
	SortAscending interface{} `field:"optional" json:"sortAscending" yaml:"sortAscending"`
}

