package datastackitpostgresflexdatabase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitPostgresflexDatabaseConfig struct {
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
	// Database ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.31.1/docs/data-sources/postgresflex_database#database_id DataStackitPostgresflexDatabase#database_id}
	DatabaseId *string `field:"required" json:"databaseId" yaml:"databaseId"`
	// ID of the Postgres Flex instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.31.1/docs/data-sources/postgresflex_database#instance_id DataStackitPostgresflexDatabase#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.31.1/docs/data-sources/postgresflex_database#project_id DataStackitPostgresflexDatabase#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

