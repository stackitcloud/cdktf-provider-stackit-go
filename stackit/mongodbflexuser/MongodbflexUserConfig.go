package mongodbflexuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MongodbflexUserConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.2/docs/resources/mongodbflex_user#database MongodbflexUser#database}.
	Database *string `field:"required" json:"database" yaml:"database"`
	// ID of the MongoDB Flex instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.2/docs/resources/mongodbflex_user#instance_id MongodbflexUser#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.2/docs/resources/mongodbflex_user#project_id MongodbflexUser#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Database access levels for the user. Some of the possible values are: [`read`, `readWrite`, `readWriteAnyDatabase`].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.2/docs/resources/mongodbflex_user#roles MongodbflexUser#roles}
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.2/docs/resources/mongodbflex_user#region MongodbflexUser#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.2/docs/resources/mongodbflex_user#username MongodbflexUser#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

