package datastackittelemetryrouterdestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitTelemetryrouterDestinationConfig struct {
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
	// The TelemetryRouter destination ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/data-sources/telemetryrouter_destination#destination_id DataStackitTelemetryrouterDestination#destination_id}
	DestinationId *string `field:"required" json:"destinationId" yaml:"destinationId"`
	// The TelemetryRouter instance ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/data-sources/telemetryrouter_destination#instance_id DataStackitTelemetryrouterDestination#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// STACKIT project ID associated with the TelemetryRouter instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/data-sources/telemetryrouter_destination#project_id DataStackitTelemetryrouterDestination#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// STACKIT region name the resource is located in. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/data-sources/telemetryrouter_destination#region DataStackitTelemetryrouterDestination#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

