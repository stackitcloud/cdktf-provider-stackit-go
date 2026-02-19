package logsaccesstoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsAccessTokenConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/logs_access_token#display_name LogsAccessToken#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The Logs instance ID associated with the access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/logs_access_token#instance_id LogsAccessToken#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The access permissions granted to the access token. Possible values: `read`, `write`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/logs_access_token#permissions LogsAccessToken#permissions}
	Permissions *[]*string `field:"required" json:"permissions" yaml:"permissions"`
	// STACKIT project ID associated with the Logs access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/logs_access_token#project_id LogsAccessToken#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The description of the access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/logs_access_token#description LogsAccessToken#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A lifetime period for an access token in days. If unset the token will not expire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/logs_access_token#lifetime LogsAccessToken#lifetime}
	Lifetime *float64 `field:"optional" json:"lifetime" yaml:"lifetime"`
	// STACKIT region name the resource is located in. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/logs_access_token#region LogsAccessToken#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

