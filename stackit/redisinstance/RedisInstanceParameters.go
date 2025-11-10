package redisinstance


type RedisInstanceParameters struct {
	// The number of milliseconds after which the instance is considered down.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#down_after_milliseconds RedisInstance#down_after_milliseconds}
	DownAfterMilliseconds *float64 `field:"optional" json:"downAfterMilliseconds" yaml:"downAfterMilliseconds"`
	// Enable monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#enable_monitoring RedisInstance#enable_monitoring}
	EnableMonitoring interface{} `field:"optional" json:"enableMonitoring" yaml:"enableMonitoring"`
	// The failover timeout in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#failover_timeout RedisInstance#failover_timeout}
	FailoverTimeout *float64 `field:"optional" json:"failoverTimeout" yaml:"failoverTimeout"`
	// Graphite server URL (host and port). If set, monitoring with Graphite will be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#graphite RedisInstance#graphite}
	Graphite *string `field:"optional" json:"graphite" yaml:"graphite"`
	// The lazy eviction enablement (yes or no).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#lazyfree_lazy_eviction RedisInstance#lazyfree_lazy_eviction}
	LazyfreeLazyEviction *string `field:"optional" json:"lazyfreeLazyEviction" yaml:"lazyfreeLazyEviction"`
	// The lazy expire enablement (yes or no).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#lazyfree_lazy_expire RedisInstance#lazyfree_lazy_expire}
	LazyfreeLazyExpire *string `field:"optional" json:"lazyfreeLazyExpire" yaml:"lazyfreeLazyExpire"`
	// The Lua time limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#lua_time_limit RedisInstance#lua_time_limit}
	LuaTimeLimit *float64 `field:"optional" json:"luaTimeLimit" yaml:"luaTimeLimit"`
	// The maximum number of clients.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#maxclients RedisInstance#maxclients}
	Maxclients *float64 `field:"optional" json:"maxclients" yaml:"maxclients"`
	// The maximum disk threshold in MB. If the disk usage exceeds this threshold, the instance will be stopped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#max_disk_threshold RedisInstance#max_disk_threshold}
	MaxDiskThreshold *float64 `field:"optional" json:"maxDiskThreshold" yaml:"maxDiskThreshold"`
	// The policy to handle the maximum memory (volatile-lru, noeviction, etc).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#maxmemory_policy RedisInstance#maxmemory_policy}
	MaxmemoryPolicy *string `field:"optional" json:"maxmemoryPolicy" yaml:"maxmemoryPolicy"`
	// The maximum memory samples.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#maxmemory_samples RedisInstance#maxmemory_samples}
	MaxmemorySamples *float64 `field:"optional" json:"maxmemorySamples" yaml:"maxmemorySamples"`
	// The frequency in seconds at which metrics are emitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#metrics_frequency RedisInstance#metrics_frequency}
	MetricsFrequency *float64 `field:"optional" json:"metricsFrequency" yaml:"metricsFrequency"`
	// The prefix for the metrics.
	//
	// Could be useful when using Graphite monitoring to prefix the metrics with a certain value, like an API key
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#metrics_prefix RedisInstance#metrics_prefix}
	MetricsPrefix *string `field:"optional" json:"metricsPrefix" yaml:"metricsPrefix"`
	// The minimum replicas maximum lag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#min_replicas_max_lag RedisInstance#min_replicas_max_lag}
	MinReplicasMaxLag *float64 `field:"optional" json:"minReplicasMaxLag" yaml:"minReplicasMaxLag"`
	// The ID of the STACKIT monitoring instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#monitoring_instance_id RedisInstance#monitoring_instance_id}
	MonitoringInstanceId *string `field:"optional" json:"monitoringInstanceId" yaml:"monitoringInstanceId"`
	// The notify keyspace events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#notify_keyspace_events RedisInstance#notify_keyspace_events}
	NotifyKeyspaceEvents *string `field:"optional" json:"notifyKeyspaceEvents" yaml:"notifyKeyspaceEvents"`
	// Comma separated list of IP networks in CIDR notation which are allowed to access this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#sgw_acl RedisInstance#sgw_acl}
	SgwAcl *string `field:"optional" json:"sgwAcl" yaml:"sgwAcl"`
	// The snapshot configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#snapshot RedisInstance#snapshot}
	Snapshot *string `field:"optional" json:"snapshot" yaml:"snapshot"`
	// List of syslog servers to send logs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#syslog RedisInstance#syslog}
	Syslog *[]*string `field:"optional" json:"syslog" yaml:"syslog"`
	// List of TLS ciphers to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#tls_ciphers RedisInstance#tls_ciphers}
	TlsCiphers *[]*string `field:"optional" json:"tlsCiphers" yaml:"tlsCiphers"`
	// TLS cipher suites to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#tls_ciphersuites RedisInstance#tls_ciphersuites}
	TlsCiphersuites *string `field:"optional" json:"tlsCiphersuites" yaml:"tlsCiphersuites"`
	// TLS protocol to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/redis_instance#tls_protocols RedisInstance#tls_protocols}
	TlsProtocols *string `field:"optional" json:"tlsProtocols" yaml:"tlsProtocols"`
}

