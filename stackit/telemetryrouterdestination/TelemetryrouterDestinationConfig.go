package telemetryrouterdestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TelemetryrouterDestinationConfig struct {
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
	// The configuration of the TelemetryRouter destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/telemetryrouter_destination#config TelemetryrouterDestination#config}
	Config *TelemetryrouterDestinationConfigA `field:"required" json:"config" yaml:"config"`
	// The displayed name of the TelemetryRouter destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/telemetryrouter_destination#display_name TelemetryrouterDestination#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The TelemetryRouter instance ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/telemetryrouter_destination#instance_id TelemetryrouterDestination#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// STACKIT project ID associated with the TelemetryRouter instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/telemetryrouter_destination#project_id TelemetryrouterDestination#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The description of the TelemetryRouter destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/telemetryrouter_destination#description TelemetryrouterDestination#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// STACKIT region name the resource is located in. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/telemetryrouter_destination#region TelemetryrouterDestination#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

