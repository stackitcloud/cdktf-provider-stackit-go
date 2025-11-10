package opensearchinstance


type OpensearchInstanceParameters struct {
	// Enable monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.64.0/docs/resources/opensearch_instance#enable_monitoring OpensearchInstance#enable_monitoring}
	EnableMonitoring interface{} `field:"optional" json:"enableMonitoring" yaml:"enableMonitoring"`
	// If set, monitoring with Graphite will be enabled.
	//
	// Expects the host and port where the Graphite metrics should be sent to (host:port).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.64.0/docs/resources/opensearch_instance#graphite OpensearchInstance#graphite}
	Graphite *string `field:"optional" json:"graphite" yaml:"graphite"`
	// The garbage collector to use for OpenSearch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.64.0/docs/resources/opensearch_instance#java_garbage_collector OpensearchInstance#java_garbage_collector}
	JavaGarbageCollector *string `field:"optional" json:"javaGarbageCollector" yaml:"javaGarbageCollector"`
	// The amount of memory (in MB) allocated as heap by the JVM for OpenSearch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.64.0/docs/resources/opensearch_instance#java_heapspace OpensearchInstance#java_heapspace}
	JavaHeapspace *float64 `field:"optional" json:"javaHeapspace" yaml:"javaHeapspace"`
	// The amount of memory (in MB) used by the JVM to store metadata for OpenSearch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.64.0/docs/resources/opensearch_instance#java_maxmetaspace OpensearchInstance#java_maxmetaspace}
	JavaMaxmetaspace *float64 `field:"optional" json:"javaMaxmetaspace" yaml:"javaMaxmetaspace"`
	// The maximum disk threshold in MB. If the disk usage exceeds this threshold, the instance will be stopped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.64.0/docs/resources/opensearch_instance#max_disk_threshold OpensearchInstance#max_disk_threshold}
	MaxDiskThreshold *float64 `field:"optional" json:"maxDiskThreshold" yaml:"maxDiskThreshold"`
	// The frequency in seconds at which metrics are emitted (in seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.64.0/docs/resources/opensearch_instance#metrics_frequency OpensearchInstance#metrics_frequency}
	MetricsFrequency *float64 `field:"optional" json:"metricsFrequency" yaml:"metricsFrequency"`
	// The prefix for the metrics.
	//
	// Could be useful when using Graphite monitoring to prefix the metrics with a certain value, like an API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.64.0/docs/resources/opensearch_instance#metrics_prefix OpensearchInstance#metrics_prefix}
	MetricsPrefix *string `field:"optional" json:"metricsPrefix" yaml:"metricsPrefix"`
	// The ID of the STACKIT monitoring instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.64.0/docs/resources/opensearch_instance#monitoring_instance_id OpensearchInstance#monitoring_instance_id}
	MonitoringInstanceId *string `field:"optional" json:"monitoringInstanceId" yaml:"monitoringInstanceId"`
	// List of plugins to install.
	//
	// Must be a supported plugin name. The plugins `repository-s3` and `repository-azure` are enabled by default and cannot be disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.64.0/docs/resources/opensearch_instance#plugins OpensearchInstance#plugins}
	Plugins *[]*string `field:"optional" json:"plugins" yaml:"plugins"`
	// Comma separated list of IP networks in CIDR notation which are allowed to access this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.64.0/docs/resources/opensearch_instance#sgw_acl OpensearchInstance#sgw_acl}
	SgwAcl *string `field:"optional" json:"sgwAcl" yaml:"sgwAcl"`
	// List of syslog servers to send logs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.64.0/docs/resources/opensearch_instance#syslog OpensearchInstance#syslog}
	Syslog *[]*string `field:"optional" json:"syslog" yaml:"syslog"`
	// List of TLS ciphers to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.64.0/docs/resources/opensearch_instance#tls_ciphers OpensearchInstance#tls_ciphers}
	TlsCiphers *[]*string `field:"optional" json:"tlsCiphers" yaml:"tlsCiphers"`
	// The TLS protocol to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.64.0/docs/resources/opensearch_instance#tls_protocols OpensearchInstance#tls_protocols}
	TlsProtocols *[]*string `field:"optional" json:"tlsProtocols" yaml:"tlsProtocols"`
}

