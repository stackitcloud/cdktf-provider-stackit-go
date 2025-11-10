package serverbackupschedule


type ServerBackupScheduleBackupProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.26.2/docs/resources/server_backup_schedule#name ServerBackupSchedule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.26.2/docs/resources/server_backup_schedule#retention_period ServerBackupSchedule#retention_period}.
	RetentionPeriod *float64 `field:"required" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.26.2/docs/resources/server_backup_schedule#volume_ids ServerBackupSchedule#volume_ids}.
	VolumeIds *[]*string `field:"optional" json:"volumeIds" yaml:"volumeIds"`
}

