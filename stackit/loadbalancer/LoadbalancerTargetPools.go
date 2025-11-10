package loadbalancer


type LoadbalancerTargetPools struct {
	// Target pool name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.6.3/docs/resources/loadbalancer#name Loadbalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Identical port number where each target listens for traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.6.3/docs/resources/loadbalancer#target_port Loadbalancer#target_port}
	TargetPort *float64 `field:"required" json:"targetPort" yaml:"targetPort"`
	// List of all targets which will be used in the pool. Limited to 250.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.6.3/docs/resources/loadbalancer#targets Loadbalancer#targets}
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.6.3/docs/resources/loadbalancer#active_health_check Loadbalancer#active_health_check}.
	ActiveHealthCheck *LoadbalancerTargetPoolsActiveHealthCheck `field:"optional" json:"activeHealthCheck" yaml:"activeHealthCheck"`
}

