package argusinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ArgusInstanceConfig struct {
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
	// The name of the Argus instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.7.1/docs/resources/argus_instance#name ArgusInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the Argus plan. E.g. `Monitoring-Medium-EU01`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.7.1/docs/resources/argus_instance#plan_name ArgusInstance#plan_name}
	PlanName *string `field:"required" json:"planName" yaml:"planName"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.7.1/docs/resources/argus_instance#project_id ArgusInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Additional parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.7.1/docs/resources/argus_instance#parameters ArgusInstance#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

