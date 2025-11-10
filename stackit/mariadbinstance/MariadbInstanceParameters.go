package mariadbinstance


type MariadbInstanceParameters struct {
	// Enable monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.56.0/docs/resources/mariadb_instance#enable_monitoring MariadbInstance#enable_monitoring}
	EnableMonitoring interface{} `field:"optional" json:"enableMonitoring" yaml:"enableMonitoring"`
	// Graphite server URL (host and port). If set, monitoring with Graphite will be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.56.0/docs/resources/mariadb_instance#graphite MariadbInstance#graphite}
	Graphite *string `field:"optional" json:"graphite" yaml:"graphite"`
	// The maximum disk threshold in MB. If the disk usage exceeds this threshold, the instance will be stopped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.56.0/docs/resources/mariadb_instance#max_disk_threshold MariadbInstance#max_disk_threshold}
	MaxDiskThreshold *float64 `field:"optional" json:"maxDiskThreshold" yaml:"maxDiskThreshold"`
	// The frequency in seconds at which metrics are emitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.56.0/docs/resources/mariadb_instance#metrics_frequency MariadbInstance#metrics_frequency}
	MetricsFrequency *float64 `field:"optional" json:"metricsFrequency" yaml:"metricsFrequency"`
	// The prefix for the metrics.
	//
	// Could be useful when using Graphite monitoring to prefix the metrics with a certain value, like an API key
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.56.0/docs/resources/mariadb_instance#metrics_prefix MariadbInstance#metrics_prefix}
	MetricsPrefix *string `field:"optional" json:"metricsPrefix" yaml:"metricsPrefix"`
	// The ID of the STACKIT monitoring instance. Monitoring instances with the plan "Observability-Monitoring-Starter" are not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.56.0/docs/resources/mariadb_instance#monitoring_instance_id MariadbInstance#monitoring_instance_id}
	MonitoringInstanceId *string `field:"optional" json:"monitoringInstanceId" yaml:"monitoringInstanceId"`
	// Comma separated list of IP networks in CIDR notation which are allowed to access this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.56.0/docs/resources/mariadb_instance#sgw_acl MariadbInstance#sgw_acl}
	SgwAcl *string `field:"optional" json:"sgwAcl" yaml:"sgwAcl"`
	// List of syslog servers to send logs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.56.0/docs/resources/mariadb_instance#syslog MariadbInstance#syslog}
	Syslog *[]*string `field:"optional" json:"syslog" yaml:"syslog"`
}

