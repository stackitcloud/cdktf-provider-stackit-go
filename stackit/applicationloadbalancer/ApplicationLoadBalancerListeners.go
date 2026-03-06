package applicationloadbalancer


type ApplicationLoadBalancerListeners struct {
	// Configuration for HTTP traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.85.0/docs/resources/application_load_balancer#http ApplicationLoadBalancer#http}
	Http *ApplicationLoadBalancerListenersHttp `field:"required" json:"http" yaml:"http"`
	// Unique name for the listener.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.85.0/docs/resources/application_load_balancer#name ApplicationLoadBalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Port number on which the listener receives incoming traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.85.0/docs/resources/application_load_balancer#port ApplicationLoadBalancer#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Protocol is the highest network protocol we understand to load balance. Possible values are: `PROTOCOL_UNSPECIFIED`, `PROTOCOL_HTTP`, `PROTOCOL_HTTPS`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.85.0/docs/resources/application_load_balancer#protocol ApplicationLoadBalancer#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Configuration for handling HTTPS traffic on this listener.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.85.0/docs/resources/application_load_balancer#https ApplicationLoadBalancer#https}
	Https *ApplicationLoadBalancerListenersHttps `field:"optional" json:"https" yaml:"https"`
	// Enable Web Application Firewall (WAF), referenced by name.
	//
	// See "Application Load Balancer - Web Application Firewall API" for more information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.85.0/docs/resources/application_load_balancer#waf_config_name ApplicationLoadBalancer#waf_config_name}
	WafConfigName *string `field:"optional" json:"wafConfigName" yaml:"wafConfigName"`
}

