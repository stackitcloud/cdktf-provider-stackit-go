package cdncustomdomain


type CdnCustomDomainCertificate struct {
	// The PEM-encoded TLS certificate. Required for custom certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.116.0/docs/resources/cdn_custom_domain#certificate CdnCustomDomain#certificate}
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// The PEM-encoded private key for the certificate.
	//
	// Required for custom certificates. The certificate will be updated if this field is changed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.116.0/docs/resources/cdn_custom_domain#private_key CdnCustomDomain#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// When true, skips the verification check that the custom domain points to the distribution domain via CNAME.
	//
	// Useful for zero-downtime migrations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.116.0/docs/resources/cdn_custom_domain#skip_dns_check CdnCustomDomain#skip_dns_check}
	SkipDnsCheck interface{} `field:"optional" json:"skipDnsCheck" yaml:"skipDnsCheck"`
}

