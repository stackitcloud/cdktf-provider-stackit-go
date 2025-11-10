package modelservingtoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelservingTokenConfig struct {
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
	// Name of the AI model serving auth token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.69.0/docs/resources/modelserving_token#name ModelservingToken#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the AI model serving auth token is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.69.0/docs/resources/modelserving_token#project_id ModelservingToken#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The description of the AI model serving auth token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.69.0/docs/resources/modelserving_token#description ModelservingToken#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Region to which the AI model serving auth token is associated. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.69.0/docs/resources/modelserving_token#region ModelservingToken#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// A map of arbitrary key/value pairs that will force recreation of the token when they change, enabling token rotation based on external conditions such as a rotating timestamp.
	//
	// Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.69.0/docs/resources/modelserving_token#rotate_when_changed ModelservingToken#rotate_when_changed}
	RotateWhenChanged *map[string]*string `field:"optional" json:"rotateWhenChanged" yaml:"rotateWhenChanged"`
	// The TTL duration of the AI model serving auth token. E.g. 5h30m40s,5h,5h30m,30m,30s.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.69.0/docs/resources/modelserving_token#ttl_duration ModelservingToken#ttl_duration}
	TtlDuration *string `field:"optional" json:"ttlDuration" yaml:"ttlDuration"`
}

