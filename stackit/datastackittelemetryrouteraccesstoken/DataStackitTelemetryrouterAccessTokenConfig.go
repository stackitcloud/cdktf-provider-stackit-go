package datastackittelemetryrouteraccesstoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitTelemetryrouterAccessTokenConfig struct {
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
	// The access token ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/data-sources/telemetryrouter_access_token#access_token_id DataStackitTelemetryrouterAccessToken#access_token_id}
	AccessTokenId *string `field:"required" json:"accessTokenId" yaml:"accessTokenId"`
	// The TelemetryRouter instance ID associated with the access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/data-sources/telemetryrouter_access_token#instance_id DataStackitTelemetryrouterAccessToken#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// STACKIT project ID associated with the TelemetryRouter access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/data-sources/telemetryrouter_access_token#project_id DataStackitTelemetryrouterAccessToken#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// STACKIT region name the resource is located in. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/data-sources/telemetryrouter_access_token#region DataStackitTelemetryrouterAccessToken#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

