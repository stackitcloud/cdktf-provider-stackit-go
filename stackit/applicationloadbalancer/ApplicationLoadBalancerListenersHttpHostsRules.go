package applicationloadbalancer


type ApplicationLoadBalancerListenersHttpHostsRules struct {
	// Reference target pool by target pool name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#target_pool ApplicationLoadBalancer#target_pool}
	TargetPool *string `field:"required" json:"targetPool" yaml:"targetPool"`
	// Routing persistence via cookies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#cookie_persistence ApplicationLoadBalancer#cookie_persistence}
	CookiePersistence *ApplicationLoadBalancerListenersHttpHostsRulesCookiePersistence `field:"optional" json:"cookiePersistence" yaml:"cookiePersistence"`
	// Headers for the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#headers ApplicationLoadBalancer#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Routing via path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#path ApplicationLoadBalancer#path}
	Path *ApplicationLoadBalancerListenersHttpHostsRulesPath `field:"optional" json:"path" yaml:"path"`
	// Query parameters for the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#query_parameters ApplicationLoadBalancer#query_parameters}
	QueryParameters interface{} `field:"optional" json:"queryParameters" yaml:"queryParameters"`
	// If enabled, when client sends an HTTP request with and Upgrade header, indicating the desire to establish a Websocket connection, if backend server supports WebSocket, it responds with HTTP 101 status code, switching protocols from HTTP to WebSocket.
	//
	// Hence the client and the server can exchange data in real-time using one long-lived TCP connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#web_socket ApplicationLoadBalancer#web_socket}
	WebSocket interface{} `field:"optional" json:"webSocket" yaml:"webSocket"`
}

