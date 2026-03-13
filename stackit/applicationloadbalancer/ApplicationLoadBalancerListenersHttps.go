package applicationloadbalancer


type ApplicationLoadBalancerListenersHttps struct {
	// TLS termination certificate configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#certificate_config ApplicationLoadBalancer#certificate_config}
	CertificateConfig *ApplicationLoadBalancerListenersHttpsCertificateConfig `field:"required" json:"certificateConfig" yaml:"certificateConfig"`
}

