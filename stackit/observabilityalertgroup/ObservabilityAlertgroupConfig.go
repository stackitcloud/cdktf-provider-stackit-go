package observabilityalertgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObservabilityAlertgroupConfig struct {
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
	// Observability instance ID to which the alert group is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.65.0/docs/resources/observability_alertgroup#instance_id ObservabilityAlertgroup#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The name of the alert group. Is the identifier and must be unique in the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.65.0/docs/resources/observability_alertgroup#name ObservabilityAlertgroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the alert group is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.65.0/docs/resources/observability_alertgroup#project_id ObservabilityAlertgroup#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Rules for the alert group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.65.0/docs/resources/observability_alertgroup#rules ObservabilityAlertgroup#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// Specifies the frequency at which rules within the group are evaluated.
	//
	// The interval must be at least 60 seconds and defaults to 60 seconds if not set. Supported formats include hours, minutes, and seconds, either singly or in combination. Examples of valid formats are: '5h30m40s', '5h', '5h30m', '60m', and '60s'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.65.0/docs/resources/observability_alertgroup#interval ObservabilityAlertgroup#interval}
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
}

