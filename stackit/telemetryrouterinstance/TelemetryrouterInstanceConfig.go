package telemetryrouterinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TelemetryrouterInstanceConfig struct {
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
	// The display name of the TelemetryRouter instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/telemetryrouter_instance#display_name TelemetryrouterInstance#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// STACKIT project ID associated with the TelemetryRouter instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/telemetryrouter_instance#project_id TelemetryrouterInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The description of the TelemetryRouter instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/telemetryrouter_instance#description TelemetryrouterInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The TelemetryRouter global filter settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/telemetryrouter_instance#filter TelemetryrouterInstance#filter}
	Filter *TelemetryrouterInstanceFilter `field:"optional" json:"filter" yaml:"filter"`
	// STACKIT region name the resource is located in. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/telemetryrouter_instance#region TelemetryrouterInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

