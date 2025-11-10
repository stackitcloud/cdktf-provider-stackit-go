package loadbalancer


type LoadbalancerListeners struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.8.0/docs/resources/loadbalancer#display_name Loadbalancer#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Port number where we listen for traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.8.0/docs/resources/loadbalancer#port Loadbalancer#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Protocol is the highest network protocol we understand to load balance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.8.0/docs/resources/loadbalancer#protocol Loadbalancer#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Reference target pool by target pool name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.8.0/docs/resources/loadbalancer#target_pool Loadbalancer#target_pool}
	TargetPool *string `field:"optional" json:"targetPool" yaml:"targetPool"`
}

