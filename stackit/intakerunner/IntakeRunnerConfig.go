package intakerunner

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntakeRunnerConfig struct {
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
	// The maximum message size in KiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/intake_runner#max_message_size_kib IntakeRunner#max_message_size_kib}
	MaxMessageSizeKib *float64 `field:"required" json:"maxMessageSizeKib" yaml:"maxMessageSizeKib"`
	// The maximum number of messages per hour.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/intake_runner#max_messages_per_hour IntakeRunner#max_messages_per_hour}
	MaxMessagesPerHour *float64 `field:"required" json:"maxMessagesPerHour" yaml:"maxMessagesPerHour"`
	// The name of the runner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/intake_runner#name IntakeRunner#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT Project ID to which the runner is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/intake_runner#project_id IntakeRunner#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The description of the runner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/intake_runner#description IntakeRunner#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User-defined labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/intake_runner#labels IntakeRunner#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/intake_runner#region IntakeRunner#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

