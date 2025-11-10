package postgresflexuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PostgresflexUserConfig struct {
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
	// ID of the PostgresFlex instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs/resources/postgresflex_user#instance_id PostgresflexUser#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs/resources/postgresflex_user#project_id PostgresflexUser#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Database access levels for the user. Supported values are: `login`, `createdb`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs/resources/postgresflex_user#roles PostgresflexUser#roles}
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs/resources/postgresflex_user#username PostgresflexUser#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs/resources/postgresflex_user#region PostgresflexUser#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

