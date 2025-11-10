package serverbackupschedule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServerBackupScheduleConfig struct {
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
	// Backup schedule details for the backups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.53.0/docs/resources/server_backup_schedule#backup_properties ServerBackupSchedule#backup_properties}
	BackupProperties *ServerBackupScheduleBackupProperties `field:"required" json:"backupProperties" yaml:"backupProperties"`
	// Is the backup schedule enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.53.0/docs/resources/server_backup_schedule#enabled ServerBackupSchedule#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The schedule name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.53.0/docs/resources/server_backup_schedule#name ServerBackupSchedule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT Project ID to which the server is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.53.0/docs/resources/server_backup_schedule#project_id ServerBackupSchedule#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Backup schedule described in `rrule` (recurrence rule) format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.53.0/docs/resources/server_backup_schedule#rrule ServerBackupSchedule#rrule}
	Rrule *string `field:"required" json:"rrule" yaml:"rrule"`
	// Server ID for the backup schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.53.0/docs/resources/server_backup_schedule#server_id ServerBackupSchedule#server_id}
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.53.0/docs/resources/server_backup_schedule#region ServerBackupSchedule#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

