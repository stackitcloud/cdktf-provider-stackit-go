package loadbalancer


type LoadbalancerListeners struct {
	// Port number where we listen for traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.59.0/docs/resources/loadbalancer#port Loadbalancer#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Protocol is the highest network protocol we understand to load balance. Supported values are: `PROTOCOL_UNSPECIFIED`, `PROTOCOL_TCP`, `PROTOCOL_UDP`, `PROTOCOL_TCP_PROXY`, `PROTOCOL_TLS_PASSTHROUGH`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.59.0/docs/resources/loadbalancer#protocol Loadbalancer#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Reference target pool by target pool name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.59.0/docs/resources/loadbalancer#target_pool Loadbalancer#target_pool}
	TargetPool *string `field:"required" json:"targetPool" yaml:"targetPool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.59.0/docs/resources/loadbalancer#display_name Loadbalancer#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// A list of domain names to match in order to pass TLS traffic to the target pool in the current listener.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.59.0/docs/resources/loadbalancer#server_name_indicators Loadbalancer#server_name_indicators}
	ServerNameIndicators interface{} `field:"optional" json:"serverNameIndicators" yaml:"serverNameIndicators"`
}

