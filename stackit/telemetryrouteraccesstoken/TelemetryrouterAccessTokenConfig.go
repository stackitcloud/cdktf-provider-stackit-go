package telemetryrouteraccesstoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TelemetryrouterAccessTokenConfig struct {
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
	// The displayed name of the access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/telemetryrouter_access_token#display_name TelemetryrouterAccessToken#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The TelemetryRouter instance ID associated with the access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/telemetryrouter_access_token#instance_id TelemetryrouterAccessToken#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// STACKIT project ID associated with the TelemetryRouter access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/telemetryrouter_access_token#project_id TelemetryrouterAccessToken#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The description of the access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/telemetryrouter_access_token#description TelemetryrouterAccessToken#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// STACKIT region name the resource is located in. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/telemetryrouter_access_token#region TelemetryrouterAccessToken#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The time-to-live (TTL) in days for the access token. If not set, token will not expire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/telemetryrouter_access_token#ttl TelemetryrouterAccessToken#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

