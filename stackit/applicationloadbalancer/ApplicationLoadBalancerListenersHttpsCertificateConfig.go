package applicationloadbalancer


type ApplicationLoadBalancerListenersHttpsCertificateConfig struct {
	// Certificate IDs for TLS termination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.86.0/docs/resources/application_load_balancer#certificate_ids ApplicationLoadBalancer#certificate_ids}
	CertificateIds *[]*string `field:"required" json:"certificateIds" yaml:"certificateIds"`
}

