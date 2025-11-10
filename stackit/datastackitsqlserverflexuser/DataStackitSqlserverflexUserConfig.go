package datastackitsqlserverflexuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitSqlserverflexUserConfig struct {
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
	// ID of the SQLServer Flex instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.27.0/docs/data-sources/sqlserverflex_user#instance_id DataStackitSqlserverflexUser#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.27.0/docs/data-sources/sqlserverflex_user#project_id DataStackitSqlserverflexUser#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// User ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.27.0/docs/data-sources/sqlserverflex_user#user_id DataStackitSqlserverflexUser#user_id}
	UserId *string `field:"required" json:"userId" yaml:"userId"`
}

