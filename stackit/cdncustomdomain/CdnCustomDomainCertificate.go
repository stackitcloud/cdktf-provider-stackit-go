package cdncustomdomain


type CdnCustomDomainCertificate struct {
	// The PEM-encoded TLS certificate. Required for custom certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/resources/cdn_custom_domain#certificate CdnCustomDomain#certificate}
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// The PEM-encoded private key for the certificate.
	//
	// Required for custom certificates. The certificate will be updated if this field is changed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/resources/cdn_custom_domain#private_key CdnCustomDomain#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
}

