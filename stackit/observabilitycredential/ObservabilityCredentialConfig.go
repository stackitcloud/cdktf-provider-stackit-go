package observabilitycredential

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObservabilityCredentialConfig struct {
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
	// The Observability Instance ID the credential belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/observability_credential#instance_id ObservabilityCredential#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// STACKIT project ID to which the credential is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/observability_credential#project_id ObservabilityCredential#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// A description of the credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/observability_credential#description ObservabilityCredential#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A map of arbitrary key/value pairs that will force recreation of the resource when they change, enabling resource rotation based on external conditions such as a rotating timestamp.
	//
	// Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/observability_credential#rotate_when_changed ObservabilityCredential#rotate_when_changed}
	RotateWhenChanged *map[string]*string `field:"optional" json:"rotateWhenChanged" yaml:"rotateWhenChanged"`
}

