package datastackitloadbalancer


type DataStackitLoadbalancerTargetPoolsSessionPersistence struct {
	// If true then all connections from one source IP address are redirected to the same target.
	//
	// This setting changes the load balancing algorithm to Maglev.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.36.0/docs/data-sources/loadbalancer#use_source_ip_address DataStackitLoadbalancer#use_source_ip_address}
	UseSourceIpAddress interface{} `field:"optional" json:"useSourceIpAddress" yaml:"useSourceIpAddress"`
}

