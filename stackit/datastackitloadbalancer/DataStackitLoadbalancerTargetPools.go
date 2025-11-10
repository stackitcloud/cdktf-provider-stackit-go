package datastackitloadbalancer


type DataStackitLoadbalancerTargetPools struct {
	// Here you can setup various session persistence options, so far only "`use_source_ip_address`" is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.11.0/docs/data-sources/loadbalancer#session_persistence DataStackitLoadbalancer#session_persistence}
	SessionPersistence *DataStackitLoadbalancerTargetPoolsSessionPersistence `field:"optional" json:"sessionPersistence" yaml:"sessionPersistence"`
}

