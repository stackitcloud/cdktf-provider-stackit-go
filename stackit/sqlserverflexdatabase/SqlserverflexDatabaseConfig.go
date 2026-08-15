package sqlserverflexdatabase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SqlserverflexDatabaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/sqlserverflex_database#instance_id SqlserverflexDatabase#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Name of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/sqlserverflex_database#name SqlserverflexDatabase#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The owner of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/sqlserverflex_database#owner SqlserverflexDatabase#owner}
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// STACKIT project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/sqlserverflex_database#project_id SqlserverflexDatabase#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The collation of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/sqlserverflex_database#collation SqlserverflexDatabase#collation}
	Collation *string `field:"optional" json:"collation" yaml:"collation"`
	// Compatibility level of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/sqlserverflex_database#compatibility SqlserverflexDatabase#compatibility}
	Compatibility *float64 `field:"optional" json:"compatibility" yaml:"compatibility"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/sqlserverflex_database#region SqlserverflexDatabase#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/sqlserverflex_database#timeouts SqlserverflexDatabase#timeouts}.
	Timeouts *SqlserverflexDatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

