package cdndistribution


type CdnDistributionConfigLogSinkCredentials struct {
	// The password corresponding to `username`.
	//
	// Required for Loki and for OTLP when `credentials.type` is `basic`. Write-only argument `password_wo` should be preferred.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.122.0/docs/resources/cdn_distribution#password CdnDistribution#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Write-only password for log sink authentication.
	//
	// Never stored in state and never returned by the API. To rotate the password, update this value AND increment `password_wo_version`. Changing this field alone will NOT trigger an update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.122.0/docs/resources/cdn_distribution#password_wo CdnDistribution#password_wo}
	PasswordWo *string `field:"optional" json:"passwordWo" yaml:"passwordWo"`
	// User-managed rotation counter for `password_wo`.
	//
	// Must be incremented every time `password_wo` is changed. Note: because the API treats credentials atomically (username + password together for Loki/OTLP basic), the password is also automatically re-sent when `username`, `log_sink.type`, or `credentials.type` changes, even without a version bump.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.122.0/docs/resources/cdn_distribution#password_wo_version CdnDistribution#password_wo_version}
	PasswordWoVersion *float64 `field:"optional" json:"passwordWoVersion" yaml:"passwordWoVersion"`
	// The bearer token used to authenticate. Required for OTLP when `credentials.type` is `bearer`. Write-only argument `token_wo` should be preferred.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.122.0/docs/resources/cdn_distribution#token CdnDistribution#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Write-only bearer token for OTLP log sink authentication.
	//
	// Never stored in state and never returned by the API. To rotate the token, update this value AND increment `token_wo_version`. Changing this field alone will NOT trigger an update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.122.0/docs/resources/cdn_distribution#token_wo CdnDistribution#token_wo}
	TokenWo *string `field:"optional" json:"tokenWo" yaml:"tokenWo"`
	// User-managed rotation counter for `token_wo`.
	//
	// Must be incremented every time `token_wo` is changed. Note: the token is also automatically re-sent when `log_sink.type` or `credentials.type` changes, even without a version bump.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.122.0/docs/resources/cdn_distribution#token_wo_version CdnDistribution#token_wo_version}
	TokenWoVersion *float64 `field:"optional" json:"tokenWoVersion" yaml:"tokenWoVersion"`
	// The authentication type when using an OTLP log sink. Leave unset for Loki.Possible values are: `basic`, `bearer`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.122.0/docs/resources/cdn_distribution#type CdnDistribution#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The username used to authenticate. Required for Loki and for OTLP when `credentials.type` is `basic`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.122.0/docs/resources/cdn_distribution#username CdnDistribution#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

