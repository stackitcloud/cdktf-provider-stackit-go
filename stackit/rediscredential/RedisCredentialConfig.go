package rediscredential

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedisCredentialConfig struct {
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
	// ID of the Redis instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/redis_credential#instance_id RedisCredential#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// STACKIT Project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/redis_credential#project_id RedisCredential#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// A map of arbitrary key/value pairs that will force recreation of the resource when they change, enabling resource rotation based on external conditions such as a rotating timestamp.
	//
	// Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/redis_credential#rotate_when_changed RedisCredential#rotate_when_changed}
	RotateWhenChanged *map[string]*string `field:"optional" json:"rotateWhenChanged" yaml:"rotateWhenChanged"`
}

