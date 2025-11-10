package rabbitmqinstance


type RabbitmqInstanceParameters struct {
	// The timeout in milliseconds for the consumer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/rabbitmq_instance#consumer_timeout RabbitmqInstance#consumer_timeout}
	ConsumerTimeout *float64 `field:"optional" json:"consumerTimeout" yaml:"consumerTimeout"`
	// Enable monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/rabbitmq_instance#enable_monitoring RabbitmqInstance#enable_monitoring}
	EnableMonitoring interface{} `field:"optional" json:"enableMonitoring" yaml:"enableMonitoring"`
	// Graphite server URL (host and port). If set, monitoring with Graphite will be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/rabbitmq_instance#graphite RabbitmqInstance#graphite}
	Graphite *string `field:"optional" json:"graphite" yaml:"graphite"`
	// The maximum disk threshold in MB. If the disk usage exceeds this threshold, the instance will be stopped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/rabbitmq_instance#max_disk_threshold RabbitmqInstance#max_disk_threshold}
	MaxDiskThreshold *float64 `field:"optional" json:"maxDiskThreshold" yaml:"maxDiskThreshold"`
	// The frequency in seconds at which metrics are emitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/rabbitmq_instance#metrics_frequency RabbitmqInstance#metrics_frequency}
	MetricsFrequency *float64 `field:"optional" json:"metricsFrequency" yaml:"metricsFrequency"`
	// The prefix for the metrics.
	//
	// Could be useful when using Graphite monitoring to prefix the metrics with a certain value, like an API key
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/rabbitmq_instance#metrics_prefix RabbitmqInstance#metrics_prefix}
	MetricsPrefix *string `field:"optional" json:"metricsPrefix" yaml:"metricsPrefix"`
	// The ID of the STACKIT monitoring instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/rabbitmq_instance#monitoring_instance_id RabbitmqInstance#monitoring_instance_id}
	MonitoringInstanceId *string `field:"optional" json:"monitoringInstanceId" yaml:"monitoringInstanceId"`
	// List of plugins to install. Must be a supported plugin name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/rabbitmq_instance#plugins RabbitmqInstance#plugins}
	Plugins *[]*string `field:"optional" json:"plugins" yaml:"plugins"`
	// List of roles to assign to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/rabbitmq_instance#roles RabbitmqInstance#roles}
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
	// Comma separated list of IP networks in CIDR notation which are allowed to access this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/rabbitmq_instance#sgw_acl RabbitmqInstance#sgw_acl}
	SgwAcl *string `field:"optional" json:"sgwAcl" yaml:"sgwAcl"`
	// List of syslog servers to send logs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/rabbitmq_instance#syslog RabbitmqInstance#syslog}
	Syslog *[]*string `field:"optional" json:"syslog" yaml:"syslog"`
	// List of TLS ciphers to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/rabbitmq_instance#tls_ciphers RabbitmqInstance#tls_ciphers}
	TlsCiphers *[]*string `field:"optional" json:"tlsCiphers" yaml:"tlsCiphers"`
	// TLS protocol to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/rabbitmq_instance#tls_protocols RabbitmqInstance#tls_protocols}
	TlsProtocols *string `field:"optional" json:"tlsProtocols" yaml:"tlsProtocols"`
}

