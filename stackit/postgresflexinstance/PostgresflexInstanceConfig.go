package postgresflexinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PostgresflexInstanceConfig struct {
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
	// The schedule for on what time and how often the database backup will be created.
	//
	// Must be a valid cron expression using numeric minute and hour values, e.g: '0 2 * * *'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.107.1/docs/resources/postgresflex_instance#backup_schedule PostgresflexInstance#backup_schedule}
	BackupSchedule *string `field:"required" json:"backupSchedule" yaml:"backupSchedule"`
	// Instance name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.107.1/docs/resources/postgresflex_instance#name PostgresflexInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.107.1/docs/resources/postgresflex_instance#project_id PostgresflexInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.107.1/docs/resources/postgresflex_instance#storage PostgresflexInstance#storage}.
	Storage *PostgresflexInstanceStorage `field:"required" json:"storage" yaml:"storage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.107.1/docs/resources/postgresflex_instance#version PostgresflexInstance#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// The Access Control List (ACL) for the PostgresFlex instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.107.1/docs/resources/postgresflex_instance#acl PostgresflexInstance#acl}
	Acl *[]*string `field:"optional" json:"acl" yaml:"acl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.107.1/docs/resources/postgresflex_instance#encryption PostgresflexInstance#encryption}.
	Encryption *PostgresflexInstanceEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.107.1/docs/resources/postgresflex_instance#flavor PostgresflexInstance#flavor}.
	Flavor *PostgresflexInstanceFlavor `field:"optional" json:"flavor" yaml:"flavor"`
	// The flavor ID of the PostgreSQL Flex instance.
	//
	// Can only be set when `flavor` and `replicas` are not set. You can list available storage classes using the [STACKIT CLI](https://github.com/stackitcloud/stackit-cli):
	// ```bash
	// stackit curl https://postgres-flex-service.api.stackit.cloud/v3/projects/{project_id}/regions/{region}/flavors\?size=100
	// ```
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.107.1/docs/resources/postgresflex_instance#flavor_id PostgresflexInstance#flavor_id}
	FlavorId *string `field:"optional" json:"flavorId" yaml:"flavorId"`
	// The network configuration of the instance. Will be required after February 2027. Set a value to prevent breaking changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.107.1/docs/resources/postgresflex_instance#network PostgresflexInstance#network}
	Network *PostgresflexInstanceNetwork `field:"optional" json:"network" yaml:"network"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.107.1/docs/resources/postgresflex_instance#region PostgresflexInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// How many replicas the instance should have.
	//
	// Valid values are 1 for single mode or 3 for replication. Can only be set together with `flavor`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.107.1/docs/resources/postgresflex_instance#replicas PostgresflexInstance#replicas}
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
	// How long backups are retained.
	//
	// The value can only be between 32 and 90 days. Will be required after February 2027. Set a value to prevent breaking changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.107.1/docs/resources/postgresflex_instance#retention_days PostgresflexInstance#retention_days}
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
}

