package datastackitsqlserverflexdatabase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitSqlserverflexDatabaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/data-sources/sqlserverflex_database#instance_id DataStackitSqlserverflexDatabase#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Name of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/data-sources/sqlserverflex_database#name DataStackitSqlserverflexDatabase#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/data-sources/sqlserverflex_database#project_id DataStackitSqlserverflexDatabase#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/data-sources/sqlserverflex_database#region DataStackitSqlserverflexDatabase#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/data-sources/sqlserverflex_database#timeouts DataStackitSqlserverflexDatabase#timeouts}.
	Timeouts *DataStackitSqlserverflexDatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

