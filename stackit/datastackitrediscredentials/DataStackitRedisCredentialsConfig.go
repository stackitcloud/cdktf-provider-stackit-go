package datastackitrediscredentials

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitRedisCredentialsConfig struct {
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
	// The credentials ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs/data-sources/redis_credentials#credentials_id DataStackitRedisCredentials#credentials_id}
	CredentialsId *string `field:"required" json:"credentialsId" yaml:"credentialsId"`
	// ID of the Redis instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs/data-sources/redis_credentials#instance_id DataStackitRedisCredentials#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs/data-sources/redis_credentials#project_id DataStackitRedisCredentials#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

