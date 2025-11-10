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
	// The Access Control List (ACL) for the PostgresFlex instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs/resources/postgresflex_instance#acl PostgresflexInstance#acl}
	Acl *[]*string `field:"required" json:"acl" yaml:"acl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs/resources/postgresflex_instance#backup_schedule PostgresflexInstance#backup_schedule}.
	BackupSchedule *string `field:"required" json:"backupSchedule" yaml:"backupSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs/resources/postgresflex_instance#flavor PostgresflexInstance#flavor}.
	Flavor *PostgresflexInstanceFlavor `field:"required" json:"flavor" yaml:"flavor"`
	// Instance name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs/resources/postgresflex_instance#name PostgresflexInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs/resources/postgresflex_instance#project_id PostgresflexInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs/resources/postgresflex_instance#replicas PostgresflexInstance#replicas}.
	Replicas *float64 `field:"required" json:"replicas" yaml:"replicas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs/resources/postgresflex_instance#storage PostgresflexInstance#storage}.
	Storage *PostgresflexInstanceStorage `field:"required" json:"storage" yaml:"storage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs/resources/postgresflex_instance#version PostgresflexInstance#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs/resources/postgresflex_instance#region PostgresflexInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

