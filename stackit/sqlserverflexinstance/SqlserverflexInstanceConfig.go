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
	// Instance name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/sqlserverflex_instance#name SqlserverflexInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/sqlserverflex_instance#project_id SqlserverflexInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The Access Control List (ACL) for the SQLServer Flex instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/sqlserverflex_instance#acl SqlserverflexInstance#acl}
	Acl *[]*string `field:"optional" json:"acl" yaml:"acl"`
	// The backup schedule.
	//
	// Should follow the cron scheduling system format (e.g. "0 0 * * *") Will be required in the future. Set a value to prevent breaking changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/sqlserverflex_instance#backup_schedule SqlserverflexInstance#backup_schedule}
	BackupSchedule *string `field:"optional" json:"backupSchedule" yaml:"backupSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/sqlserverflex_instance#flavor SqlserverflexInstance#flavor}.
	Flavor *SqlserverflexInstanceFlavor `field:"optional" json:"flavor" yaml:"flavor"`
	// The flavor ID of the SQLServer Flex instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/sqlserverflex_instance#flavor_id SqlserverflexInstance#flavor_id}
	FlavorId *string `field:"optional" json:"flavorId" yaml:"flavorId"`
	// The network configuration of the instance. Will be required in the future. Set a value to prevent breaking changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/sqlserverflex_instance#network SqlserverflexInstance#network}
	Network *SqlserverflexInstanceNetwork `field:"optional" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/sqlserverflex_instance#options SqlserverflexInstance#options}.
	Options *SqlserverflexInstanceOptions `field:"optional" json:"options" yaml:"options"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/sqlserverflex_instance#region SqlserverflexInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The days (30 to 90) for how long the backup files should be stored before cleaned up.
	//
	// Will be required in the future. Set a value to prevent breaking changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/sqlserverflex_instance#retention_days SqlserverflexInstance#retention_days}
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
	// The object containing information about the storage size and class.
	//
	// Will be required in the future. Set a value to prevent breaking changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/sqlserverflex_instance#storage SqlserverflexInstance#storage}
	Storage *SqlserverflexInstanceStorage `field:"optional" json:"storage" yaml:"storage"`
	// The sqlserver version used for the instance.
	//
	// Possible values are: `2022`. Will be required in the future. Set a value to prevent breaking changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/sqlserverflex_instance#version SqlserverflexInstance#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

