package modelexperimentstoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelexperimentsTokenConfig struct {
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
	// The AI Model Experiments instance ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.113.0/docs/resources/modelexperiments_token#instance_id ModelexperimentsToken#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The display name is a short name chosen by the user to identify the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.113.0/docs/resources/modelexperiments_token#name ModelexperimentsToken#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT Project ID to which the resource is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.113.0/docs/resources/modelexperiments_token#project_id ModelexperimentsToken#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The description is a longer text chosen by the user to provide more context for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.113.0/docs/resources/modelexperiments_token#description ModelexperimentsToken#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A map of arbitrary key/value pairs that can be attached to the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.113.0/docs/resources/modelexperiments_token#labels ModelexperimentsToken#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The STACKIT region name the resource is located in. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.113.0/docs/resources/modelexperiments_token#region ModelexperimentsToken#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// A map of arbitrary key/value pairs that will force recreation of the resource when they change, enabling resource rotation based on external conditions such as a rotating timestamp.
	//
	// Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.113.0/docs/resources/modelexperiments_token#rotate_when_changed ModelexperimentsToken#rotate_when_changed}
	RotateWhenChanged *map[string]*string `field:"optional" json:"rotateWhenChanged" yaml:"rotateWhenChanged"`
	// The TTL duration of the AI Model Experiments instance token. E.g. 5h30m40s,5h,5h30m,30m,30s.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.113.0/docs/resources/modelexperiments_token#ttl_duration ModelexperimentsToken#ttl_duration}
	TtlDuration *string `field:"optional" json:"ttlDuration" yaml:"ttlDuration"`
}

