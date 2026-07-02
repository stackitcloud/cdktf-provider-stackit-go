package datastackitintakerunner

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitIntakeRunnerConfig struct {
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
	// STACKIT Project ID to which the runner is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/data-sources/intake_runner#project_id DataStackitIntakeRunner#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The runner ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/data-sources/intake_runner#runner_id DataStackitIntakeRunner#runner_id}
	RunnerId *string `field:"required" json:"runnerId" yaml:"runnerId"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/data-sources/intake_runner#region DataStackitIntakeRunner#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

