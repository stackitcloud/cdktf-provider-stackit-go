package datastackitobservabilityalertgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitObservabilityAlertgroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.53.0/docs/data-sources/observability_alertgroup#instance_id DataStackitObservabilityAlertgroup#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The name of the alert group. Is the identifier and must be unique in the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.53.0/docs/data-sources/observability_alertgroup#name DataStackitObservabilityAlertgroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the alert group is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.53.0/docs/data-sources/observability_alertgroup#project_id DataStackitObservabilityAlertgroup#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

