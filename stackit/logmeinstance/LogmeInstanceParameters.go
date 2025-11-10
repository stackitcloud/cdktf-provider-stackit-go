package logmeinstance


type LogmeInstanceParameters struct {
	// Enable monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#enable_monitoring LogmeInstance#enable_monitoring}
	EnableMonitoring interface{} `field:"optional" json:"enableMonitoring" yaml:"enableMonitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#fluentd_tcp LogmeInstance#fluentd_tcp}.
	FluentdTcp *float64 `field:"optional" json:"fluentdTcp" yaml:"fluentdTcp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#fluentd_tls LogmeInstance#fluentd_tls}.
	FluentdTls *float64 `field:"optional" json:"fluentdTls" yaml:"fluentdTls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#fluentd_tls_ciphers LogmeInstance#fluentd_tls_ciphers}.
	FluentdTlsCiphers *string `field:"optional" json:"fluentdTlsCiphers" yaml:"fluentdTlsCiphers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#fluentd_tls_max_version LogmeInstance#fluentd_tls_max_version}.
	FluentdTlsMaxVersion *string `field:"optional" json:"fluentdTlsMaxVersion" yaml:"fluentdTlsMaxVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#fluentd_tls_min_version LogmeInstance#fluentd_tls_min_version}.
	FluentdTlsMinVersion *string `field:"optional" json:"fluentdTlsMinVersion" yaml:"fluentdTlsMinVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#fluentd_tls_version LogmeInstance#fluentd_tls_version}.
	FluentdTlsVersion *string `field:"optional" json:"fluentdTlsVersion" yaml:"fluentdTlsVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#fluentd_udp LogmeInstance#fluentd_udp}.
	FluentdUdp *float64 `field:"optional" json:"fluentdUdp" yaml:"fluentdUdp"`
	// If set, monitoring with Graphite will be enabled.
	//
	// Expects the host and port where the Graphite metrics should be sent to (host:port).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#graphite LogmeInstance#graphite}
	Graphite *string `field:"optional" json:"graphite" yaml:"graphite"`
	// Combination of an integer and a timerange when an index will be considered "old" and can be deleted.
	//
	// Possible values for the timerange are `s`, `m`, `h` and `d`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#ism_deletion_after LogmeInstance#ism_deletion_after}
	IsmDeletionAfter *string `field:"optional" json:"ismDeletionAfter" yaml:"ismDeletionAfter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#ism_jitter LogmeInstance#ism_jitter}.
	IsmJitter *float64 `field:"optional" json:"ismJitter" yaml:"ismJitter"`
	// Jitter of the execution time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#ism_job_interval LogmeInstance#ism_job_interval}
	IsmJobInterval *float64 `field:"optional" json:"ismJobInterval" yaml:"ismJobInterval"`
	// The amount of memory (in MB) allocated as heap by the JVM for OpenSearch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#java_heapspace LogmeInstance#java_heapspace}
	JavaHeapspace *float64 `field:"optional" json:"javaHeapspace" yaml:"javaHeapspace"`
	// The amount of memory (in MB) used by the JVM to store metadata for OpenSearch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#java_maxmetaspace LogmeInstance#java_maxmetaspace}
	JavaMaxmetaspace *float64 `field:"optional" json:"javaMaxmetaspace" yaml:"javaMaxmetaspace"`
	// The maximum disk threshold in MB. If the disk usage exceeds this threshold, the instance will be stopped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#max_disk_threshold LogmeInstance#max_disk_threshold}
	MaxDiskThreshold *float64 `field:"optional" json:"maxDiskThreshold" yaml:"maxDiskThreshold"`
	// The frequency in seconds at which metrics are emitted (in seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#metrics_frequency LogmeInstance#metrics_frequency}
	MetricsFrequency *float64 `field:"optional" json:"metricsFrequency" yaml:"metricsFrequency"`
	// The prefix for the metrics.
	//
	// Could be useful when using Graphite monitoring to prefix the metrics with a certain value, like an API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#metrics_prefix LogmeInstance#metrics_prefix}
	MetricsPrefix *string `field:"optional" json:"metricsPrefix" yaml:"metricsPrefix"`
	// The ID of the STACKIT monitoring instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#monitoring_instance_id LogmeInstance#monitoring_instance_id}
	MonitoringInstanceId *string `field:"optional" json:"monitoringInstanceId" yaml:"monitoringInstanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#opensearch_tls_ciphers LogmeInstance#opensearch_tls_ciphers}.
	OpensearchTlsCiphers *[]*string `field:"optional" json:"opensearchTlsCiphers" yaml:"opensearchTlsCiphers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#opensearch_tls_protocols LogmeInstance#opensearch_tls_protocols}.
	OpensearchTlsProtocols *[]*string `field:"optional" json:"opensearchTlsProtocols" yaml:"opensearchTlsProtocols"`
	// Comma separated list of IP networks in CIDR notation which are allowed to access this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#sgw_acl LogmeInstance#sgw_acl}
	SgwAcl *string `field:"optional" json:"sgwAcl" yaml:"sgwAcl"`
	// List of syslog servers to send logs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/logme_instance#syslog LogmeInstance#syslog}
	Syslog *[]*string `field:"optional" json:"syslog" yaml:"syslog"`
}

