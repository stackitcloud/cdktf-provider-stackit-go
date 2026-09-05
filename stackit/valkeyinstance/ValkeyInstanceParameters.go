package valkeyinstance


type ValkeyInstanceParameters struct {
	// The number of milliseconds after which the instance is considered down.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#down_after_milliseconds ValkeyInstance#down_after_milliseconds}
	DownAfterMilliseconds *float64 `field:"optional" json:"downAfterMilliseconds" yaml:"downAfterMilliseconds"`
	// Enable monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#enable_monitoring ValkeyInstance#enable_monitoring}
	EnableMonitoring interface{} `field:"optional" json:"enableMonitoring" yaml:"enableMonitoring"`
	// The failover timeout in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#failover_timeout ValkeyInstance#failover_timeout}
	FailoverTimeout *float64 `field:"optional" json:"failoverTimeout" yaml:"failoverTimeout"`
	// Graphite server URL (host and port). If set, monitoring with Graphite will be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#graphite ValkeyInstance#graphite}
	Graphite *string `field:"optional" json:"graphite" yaml:"graphite"`
	// The lazy eviction enablement (yes or no).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#lazyfree_lazy_eviction ValkeyInstance#lazyfree_lazy_eviction}
	LazyfreeLazyEviction *string `field:"optional" json:"lazyfreeLazyEviction" yaml:"lazyfreeLazyEviction"`
	// The lazy expire enablement (yes or no).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#lazyfree_lazy_expire ValkeyInstance#lazyfree_lazy_expire}
	LazyfreeLazyExpire *string `field:"optional" json:"lazyfreeLazyExpire" yaml:"lazyfreeLazyExpire"`
	// The Lua time limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#lua_time_limit ValkeyInstance#lua_time_limit}
	LuaTimeLimit *float64 `field:"optional" json:"luaTimeLimit" yaml:"luaTimeLimit"`
	// The maximum number of clients.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#maxclients ValkeyInstance#maxclients}
	Maxclients *float64 `field:"optional" json:"maxclients" yaml:"maxclients"`
	// The maximum disk threshold in MB. If the disk usage exceeds this threshold, the instance will be stopped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#max_disk_threshold ValkeyInstance#max_disk_threshold}
	MaxDiskThreshold *float64 `field:"optional" json:"maxDiskThreshold" yaml:"maxDiskThreshold"`
	// The policy to handle the maximum memory (volatile-lru, noeviction, etc).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#maxmemory_policy ValkeyInstance#maxmemory_policy}
	MaxmemoryPolicy *string `field:"optional" json:"maxmemoryPolicy" yaml:"maxmemoryPolicy"`
	// The maximum memory samples.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#maxmemory_samples ValkeyInstance#maxmemory_samples}
	MaxmemorySamples *float64 `field:"optional" json:"maxmemorySamples" yaml:"maxmemorySamples"`
	// The frequency in seconds at which metrics are emitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#metrics_frequency ValkeyInstance#metrics_frequency}
	MetricsFrequency *float64 `field:"optional" json:"metricsFrequency" yaml:"metricsFrequency"`
	// The prefix for the metrics.
	//
	// Could be useful when using Graphite monitoring to prefix the metrics with a certain value, like an API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#metrics_prefix ValkeyInstance#metrics_prefix}
	MetricsPrefix *string `field:"optional" json:"metricsPrefix" yaml:"metricsPrefix"`
	// The minimum replicas maximum lag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#min_replicas_max_lag ValkeyInstance#min_replicas_max_lag}
	MinReplicasMaxLag *float64 `field:"optional" json:"minReplicasMaxLag" yaml:"minReplicasMaxLag"`
	// The amount of connected replicas that are required for the primary to accept write operations.
	//
	// It can be set to 0 to disable it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#min_replicas_to_write ValkeyInstance#min_replicas_to_write}
	MinReplicasToWrite *float64 `field:"optional" json:"minReplicasToWrite" yaml:"minReplicasToWrite"`
	// The ID of the STACKIT monitoring instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#monitoring_instance_id ValkeyInstance#monitoring_instance_id}
	MonitoringInstanceId *string `field:"optional" json:"monitoringInstanceId" yaml:"monitoringInstanceId"`
	// The notify keyspace events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#notify_keyspace_events ValkeyInstance#notify_keyspace_events}
	NotifyKeyspaceEvents *string `field:"optional" json:"notifyKeyspaceEvents" yaml:"notifyKeyspaceEvents"`
	// The replication backlog size for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#repl_backlog_size ValkeyInstance#repl_backlog_size}
	ReplBacklogSize *string `field:"optional" json:"replBacklogSize" yaml:"replBacklogSize"`
	// Comma separated list of IP networks in CIDR notation which are allowed to access this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#sgw_acl ValkeyInstance#sgw_acl}
	SgwAcl *string `field:"optional" json:"sgwAcl" yaml:"sgwAcl"`
	// The snapshot configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#snapshot ValkeyInstance#snapshot}
	Snapshot *string `field:"optional" json:"snapshot" yaml:"snapshot"`
	// List of syslog servers to send logs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.114.0/docs/resources/valkey_instance#syslog ValkeyInstance#syslog}
	Syslog *[]*string `field:"optional" json:"syslog" yaml:"syslog"`
}

