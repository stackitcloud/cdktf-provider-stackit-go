package applicationloadbalancer


type ApplicationLoadBalancerTargetPoolsTlsConfig struct {
	// Specifies a custom Certificate Authority (CA).
	//
	// When provided, the target pool will trust certificates signed by this CA, in addition to any system-trusted CAs. This is useful for scenarios where the target pool needs to communicate with servers using self-signed or internally-issued certificates. Enabled needs to be set to true and skip validation to false for this option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.85.0/docs/resources/application_load_balancer#custom_ca ApplicationLoadBalancer#custom_ca}
	CustomCa *string `field:"optional" json:"customCa" yaml:"customCa"`
	// Enable TLS (Transport Layer Security) bridging for the connection between Application Load Balancer and targets in this pool.
	//
	// When enabled, public CAs are trusted. Can be used in tandem with the options either custom CA or skip validation or alone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.85.0/docs/resources/application_load_balancer#enabled ApplicationLoadBalancer#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Bypass certificate validation for TLS bridging in this target pool.
	//
	// This option is insecure and can only be used with public CAs by setting enabled true. Meant to be used for testing purposes only!
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.85.0/docs/resources/application_load_balancer#skip_certificate_validation ApplicationLoadBalancer#skip_certificate_validation}
	SkipCertificateValidation interface{} `field:"optional" json:"skipCertificateValidation" yaml:"skipCertificateValidation"`
}

