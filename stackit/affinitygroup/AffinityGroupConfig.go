package affinitygroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AffinityGroupConfig struct {
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
	// The name of the affinity group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.63.0/docs/resources/affinity_group#name AffinityGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The policy of the affinity group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.63.0/docs/resources/affinity_group#policy AffinityGroup#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// STACKIT Project ID to which the affinity group is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.63.0/docs/resources/affinity_group#project_id AffinityGroup#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

