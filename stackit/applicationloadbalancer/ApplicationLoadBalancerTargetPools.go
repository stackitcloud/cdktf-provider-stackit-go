package applicationloadbalancer


type ApplicationLoadBalancerTargetPools struct {
	// Target pool name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#name ApplicationLoadBalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The number identifying the port where each target listens for traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#target_port ApplicationLoadBalancer#target_port}
	TargetPort *float64 `field:"required" json:"targetPort" yaml:"targetPort"`
	// List of all targets which will be used in the pool. Limited to 250.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#targets ApplicationLoadBalancer#targets}
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#active_health_check ApplicationLoadBalancer#active_health_check}.
	ActiveHealthCheck *ApplicationLoadBalancerTargetPoolsActiveHealthCheck `field:"optional" json:"activeHealthCheck" yaml:"activeHealthCheck"`
	// Configuration for TLS bridging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#tls_config ApplicationLoadBalancer#tls_config}
	TlsConfig *ApplicationLoadBalancerTargetPoolsTlsConfig `field:"optional" json:"tlsConfig" yaml:"tlsConfig"`
}

