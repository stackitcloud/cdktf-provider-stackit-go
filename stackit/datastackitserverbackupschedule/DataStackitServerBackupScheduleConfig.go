package datastackitserverbackupschedule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitServerBackupScheduleConfig struct {
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
	// Backup schedule ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.22.0/docs/data-sources/server_backup_schedule#backup_schedule_id DataStackitServerBackupSchedule#backup_schedule_id}
	BackupScheduleId *float64 `field:"required" json:"backupScheduleId" yaml:"backupScheduleId"`
	// STACKIT Project ID to which the server is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.22.0/docs/data-sources/server_backup_schedule#project_id DataStackitServerBackupSchedule#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Server ID for the backup schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.22.0/docs/data-sources/server_backup_schedule#server_id DataStackitServerBackupSchedule#server_id}
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
}

