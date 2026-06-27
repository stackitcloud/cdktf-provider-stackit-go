package telemetrylink

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TelemetrylinkConfig struct {
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
	// The displayed name of the Telemetry Link resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/telemetrylink#display_name Telemetrylink#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// STACKIT project ID, folder ID, or organization ID associated with the Telemetry Link resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/telemetrylink#resource_id Telemetrylink#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The resource type of the TelemetryLink resource, possible values: Possible values are: `organization`, `folder`, `project`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/telemetrylink#resource_type Telemetrylink#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// The Telemetry Router ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/telemetrylink#telemetry_router_id Telemetrylink#telemetry_router_id}
	TelemetryRouterId *string `field:"required" json:"telemetryRouterId" yaml:"telemetryRouterId"`
	// The access token of the Telemetry Router instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/telemetrylink#access_token Telemetrylink#access_token}
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// The description of the Telemetry Link resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/telemetrylink#description Telemetrylink#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// STACKIT region name the resource is located in. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/telemetrylink#region Telemetrylink#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

