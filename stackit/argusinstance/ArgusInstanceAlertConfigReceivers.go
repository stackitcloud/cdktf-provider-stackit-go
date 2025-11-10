package argusinstance


type ArgusInstanceAlertConfigReceivers struct {
	// Name of the receiver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.30.1/docs/resources/argus_instance#name ArgusInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of email configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.30.1/docs/resources/argus_instance#email_configs ArgusInstance#email_configs}
	EmailConfigs interface{} `field:"optional" json:"emailConfigs" yaml:"emailConfigs"`
	// List of OpsGenie configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.30.1/docs/resources/argus_instance#opsgenie_configs ArgusInstance#opsgenie_configs}
	OpsgenieConfigs interface{} `field:"optional" json:"opsgenieConfigs" yaml:"opsgenieConfigs"`
	// List of Webhooks configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.30.1/docs/resources/argus_instance#webhooks_configs ArgusInstance#webhooks_configs}
	WebhooksConfigs interface{} `field:"optional" json:"webhooksConfigs" yaml:"webhooksConfigs"`
}

