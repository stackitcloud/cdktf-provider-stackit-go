package datastackitaffinitygroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitAffinityGroupConfig struct {
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
	// The affinity group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.68.0/docs/data-sources/affinity_group#affinity_group_id DataStackitAffinityGroup#affinity_group_id}
	AffinityGroupId *string `field:"required" json:"affinityGroupId" yaml:"affinityGroupId"`
	// STACKIT Project ID to which the affinity group is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.68.0/docs/data-sources/affinity_group#project_id DataStackitAffinityGroup#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

