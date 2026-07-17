package datastackittelemetrylink

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitTelemetrylinkConfig struct {
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
	// STACKIT project ID, folder ID, or organization ID associated with the Telemetry Link resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/data-sources/telemetrylink#resource_id DataStackitTelemetrylink#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The resource type of the TelemetryLink resource, possible values: Possible values are: `organization`, `folder`, `project`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/data-sources/telemetrylink#resource_type DataStackitTelemetrylink#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// STACKIT region name the resource is located in. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/data-sources/telemetrylink#region DataStackitTelemetrylink#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

