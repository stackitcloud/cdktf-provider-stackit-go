package redisinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedisInstanceConfig struct {
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
	// Instance name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/redis_instance#name RedisInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The selected plan name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/redis_instance#plan_name RedisInstance#plan_name}
	PlanName *string `field:"required" json:"planName" yaml:"planName"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/redis_instance#project_id RedisInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The service version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/redis_instance#version RedisInstance#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// Configuration parameters.
	//
	// Please note that removing a previously configured field from your Terraform configuration won't replace its value in the API. To update a previously configured field, explicitly set a new value for it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/redis_instance#parameters RedisInstance#parameters}
	Parameters *RedisInstanceParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

