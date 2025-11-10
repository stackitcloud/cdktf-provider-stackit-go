package postgresflexdatabase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PostgresflexDatabaseConfig struct {
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
	// ID of the Postgres Flex instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.26.3/docs/resources/postgresflex_database#instance_id PostgresflexDatabase#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.26.3/docs/resources/postgresflex_database#name PostgresflexDatabase#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Username of the database owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.26.3/docs/resources/postgresflex_database#owner PostgresflexDatabase#owner}
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.26.3/docs/resources/postgresflex_database#project_id PostgresflexDatabase#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

