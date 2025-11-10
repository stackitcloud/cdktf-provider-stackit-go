package postgresqlinstance


type PostgresqlInstanceParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.14.1/docs/resources/postgresql_instance#enable_monitoring PostgresqlInstance#enable_monitoring}.
	EnableMonitoring interface{} `field:"optional" json:"enableMonitoring" yaml:"enableMonitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.14.1/docs/resources/postgresql_instance#metrics_frequency PostgresqlInstance#metrics_frequency}.
	MetricsFrequency *float64 `field:"optional" json:"metricsFrequency" yaml:"metricsFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.14.1/docs/resources/postgresql_instance#metrics_prefix PostgresqlInstance#metrics_prefix}.
	MetricsPrefix *string `field:"optional" json:"metricsPrefix" yaml:"metricsPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.14.1/docs/resources/postgresql_instance#monitoring_instance_id PostgresqlInstance#monitoring_instance_id}.
	MonitoringInstanceId *string `field:"optional" json:"monitoringInstanceId" yaml:"monitoringInstanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.14.1/docs/resources/postgresql_instance#plugins PostgresqlInstance#plugins}.
	Plugins *[]*string `field:"optional" json:"plugins" yaml:"plugins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.14.1/docs/resources/postgresql_instance#sgw_acl PostgresqlInstance#sgw_acl}.
	SgwAcl *string `field:"optional" json:"sgwAcl" yaml:"sgwAcl"`
}

