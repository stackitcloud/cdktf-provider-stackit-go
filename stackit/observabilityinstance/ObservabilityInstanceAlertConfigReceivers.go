package observabilityinstance


type ObservabilityInstanceAlertConfigReceivers struct {
	// Name of the receiver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.57.0/docs/resources/observability_instance#name ObservabilityInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of email configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.57.0/docs/resources/observability_instance#email_configs ObservabilityInstance#email_configs}
	EmailConfigs interface{} `field:"optional" json:"emailConfigs" yaml:"emailConfigs"`
	// List of OpsGenie configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.57.0/docs/resources/observability_instance#opsgenie_configs ObservabilityInstance#opsgenie_configs}
	OpsgenieConfigs interface{} `field:"optional" json:"opsgenieConfigs" yaml:"opsgenieConfigs"`
	// List of Webhooks configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.57.0/docs/resources/observability_instance#webhooks_configs ObservabilityInstance#webhooks_configs}
	WebhooksConfigs interface{} `field:"optional" json:"webhooksConfigs" yaml:"webhooksConfigs"`
}

