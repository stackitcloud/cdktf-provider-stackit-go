package sqlserverflexinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SqlserverflexInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.51.0/docs/resources/sqlserverflex_instance#flavor SqlserverflexInstance#flavor}.
	Flavor *SqlserverflexInstanceFlavor `field:"required" json:"flavor" yaml:"flavor"`
	// Instance name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.51.0/docs/resources/sqlserverflex_instance#name SqlserverflexInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.51.0/docs/resources/sqlserverflex_instance#project_id SqlserverflexInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The Access Control List (ACL) for the SQLServer Flex instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.51.0/docs/resources/sqlserverflex_instance#acl SqlserverflexInstance#acl}
	Acl *[]*string `field:"optional" json:"acl" yaml:"acl"`
	// The backup schedule. Should follow the cron scheduling system format (e.g. "0 0 * * *").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.51.0/docs/resources/sqlserverflex_instance#backup_schedule SqlserverflexInstance#backup_schedule}
	BackupSchedule *string `field:"optional" json:"backupSchedule" yaml:"backupSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.51.0/docs/resources/sqlserverflex_instance#options SqlserverflexInstance#options}.
	Options *SqlserverflexInstanceOptions `field:"optional" json:"options" yaml:"options"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.51.0/docs/resources/sqlserverflex_instance#region SqlserverflexInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.51.0/docs/resources/sqlserverflex_instance#storage SqlserverflexInstance#storage}.
	Storage *SqlserverflexInstanceStorage `field:"optional" json:"storage" yaml:"storage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.51.0/docs/resources/sqlserverflex_instance#version SqlserverflexInstance#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

