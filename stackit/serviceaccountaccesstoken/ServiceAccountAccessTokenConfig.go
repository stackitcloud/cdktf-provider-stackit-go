package serviceaccountaccesstoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceAccountAccessTokenConfig struct {
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
	// STACKIT project ID associated with the service account token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.54.0/docs/resources/service_account_access_token#project_id ServiceAccountAccessToken#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Email address linked to the service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.54.0/docs/resources/service_account_access_token#service_account_email ServiceAccountAccessToken#service_account_email}
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// A map of arbitrary key/value pairs that will force recreation of the token when they change, enabling token rotation based on external conditions such as a rotating timestamp.
	//
	// Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.54.0/docs/resources/service_account_access_token#rotate_when_changed ServiceAccountAccessToken#rotate_when_changed}
	RotateWhenChanged *map[string]*string `field:"optional" json:"rotateWhenChanged" yaml:"rotateWhenChanged"`
	// Specifies the token's validity duration in days. If unspecified, defaults to 90 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.54.0/docs/resources/service_account_access_token#ttl_days ServiceAccountAccessToken#ttl_days}
	TtlDays *float64 `field:"optional" json:"ttlDays" yaml:"ttlDays"`
}

