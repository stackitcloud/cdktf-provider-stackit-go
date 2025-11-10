package observabilityinstance


type ObservabilityInstanceAlertConfigReceiversOpsgenieConfigs struct {
	// The API key for OpsGenie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/observability_instance#api_key ObservabilityInstance#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// The host to send OpsGenie API requests to. Must be a valid URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/observability_instance#api_url ObservabilityInstance#api_url}
	ApiUrl *string `field:"optional" json:"apiUrl" yaml:"apiUrl"`
	// Comma separated list of tags attached to the notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/observability_instance#tags ObservabilityInstance#tags}
	Tags *string `field:"optional" json:"tags" yaml:"tags"`
}

