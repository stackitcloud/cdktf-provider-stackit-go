package mongodbflexinstance


type MongodbflexInstanceOptions struct {
	// Type of the MongoDB Flex instance. Supported values are: `Replica`, `Sharded`, `Single`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.34.0/docs/resources/mongodbflex_instance#type MongodbflexInstance#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The number of days that daily backups will be retained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.34.0/docs/resources/mongodbflex_instance#daily_snapshot_retention_days MongodbflexInstance#daily_snapshot_retention_days}
	DailySnapshotRetentionDays *float64 `field:"optional" json:"dailySnapshotRetentionDays" yaml:"dailySnapshotRetentionDays"`
	// The number of months that monthly backups will be retained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.34.0/docs/resources/mongodbflex_instance#monthly_snapshot_retention_months MongodbflexInstance#monthly_snapshot_retention_months}
	MonthlySnapshotRetentionMonths *float64 `field:"optional" json:"monthlySnapshotRetentionMonths" yaml:"monthlySnapshotRetentionMonths"`
	// The number of hours back in time the point-in-time recovery feature will be able to recover.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.34.0/docs/resources/mongodbflex_instance#point_in_time_window_hours MongodbflexInstance#point_in_time_window_hours}
	PointInTimeWindowHours *float64 `field:"optional" json:"pointInTimeWindowHours" yaml:"pointInTimeWindowHours"`
	// The number of days that continuous backups (controlled via the `backup_schedule`) will be retained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.34.0/docs/resources/mongodbflex_instance#snapshot_retention_days MongodbflexInstance#snapshot_retention_days}
	SnapshotRetentionDays *float64 `field:"optional" json:"snapshotRetentionDays" yaml:"snapshotRetentionDays"`
	// The number of weeks that weekly backups will be retained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.34.0/docs/resources/mongodbflex_instance#weekly_snapshot_retention_weeks MongodbflexInstance#weekly_snapshot_retention_weeks}
	WeeklySnapshotRetentionWeeks *float64 `field:"optional" json:"weeklySnapshotRetentionWeeks" yaml:"weeklySnapshotRetentionWeeks"`
}

