package datastackitloadbalancer


type DataStackitLoadbalancerListeners struct {
	// A list of domain names to match in order to pass TLS traffic to the target pool in the current listener.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.0/docs/data-sources/loadbalancer#server_name_indicators DataStackitLoadbalancer#server_name_indicators}
	ServerNameIndicators interface{} `field:"optional" json:"serverNameIndicators" yaml:"serverNameIndicators"`
}

