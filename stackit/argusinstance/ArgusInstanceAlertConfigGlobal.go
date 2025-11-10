package argusinstance


type ArgusInstanceAlertConfigGlobal struct {
	// The API key for OpsGenie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.1/docs/resources/argus_instance#opsgenie_api_key ArgusInstance#opsgenie_api_key}
	OpsgenieApiKey *string `field:"optional" json:"opsgenieApiKey" yaml:"opsgenieApiKey"`
	// The host to send OpsGenie API requests to. Must be a valid URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.1/docs/resources/argus_instance#opsgenie_api_url ArgusInstance#opsgenie_api_url}
	OpsgenieApiUrl *string `field:"optional" json:"opsgenieApiUrl" yaml:"opsgenieApiUrl"`
	// The default value used by alertmanager if the alert does not include EndsAt.
	//
	// After this time passes, it can declare the alert as resolved if it has not been updated. This has no impact on alerts from Prometheus, as they always include EndsAt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.1/docs/resources/argus_instance#resolve_timeout ArgusInstance#resolve_timeout}
	ResolveTimeout *string `field:"optional" json:"resolveTimeout" yaml:"resolveTimeout"`
	// SMTP authentication information. Must be a valid email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.1/docs/resources/argus_instance#smtp_auth_identity ArgusInstance#smtp_auth_identity}
	SmtpAuthIdentity *string `field:"optional" json:"smtpAuthIdentity" yaml:"smtpAuthIdentity"`
	// SMTP Auth using LOGIN and PLAIN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.1/docs/resources/argus_instance#smtp_auth_password ArgusInstance#smtp_auth_password}
	SmtpAuthPassword *string `field:"optional" json:"smtpAuthPassword" yaml:"smtpAuthPassword"`
	// SMTP Auth using CRAM-MD5, LOGIN and PLAIN. If empty, Alertmanager doesn't authenticate to the SMTP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.1/docs/resources/argus_instance#smtp_auth_username ArgusInstance#smtp_auth_username}
	SmtpAuthUsername *string `field:"optional" json:"smtpAuthUsername" yaml:"smtpAuthUsername"`
	// The default SMTP From header field. Must be a valid email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.1/docs/resources/argus_instance#smtp_from ArgusInstance#smtp_from}
	SmtpFrom *string `field:"optional" json:"smtpFrom" yaml:"smtpFrom"`
	// The default SMTP smarthost used for sending emails, including port number in format `host:port` (eg.
	//
	// `smtp.example.com:587`). Port number usually is 25, or 587 for SMTP over TLS (sometimes referred to as STARTTLS).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.1/docs/resources/argus_instance#smtp_smart_host ArgusInstance#smtp_smart_host}
	SmtpSmartHost *string `field:"optional" json:"smtpSmartHost" yaml:"smtpSmartHost"`
}

