package observabilityinstance


type ObservabilityInstanceAlertConfigReceiversEmailConfigs struct {
	// SMTP authentication information. Must be a valid email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.69.0/docs/resources/observability_instance#auth_identity ObservabilityInstance#auth_identity}
	AuthIdentity *string `field:"optional" json:"authIdentity" yaml:"authIdentity"`
	// SMTP authentication password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.69.0/docs/resources/observability_instance#auth_password ObservabilityInstance#auth_password}
	AuthPassword *string `field:"optional" json:"authPassword" yaml:"authPassword"`
	// SMTP authentication username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.69.0/docs/resources/observability_instance#auth_username ObservabilityInstance#auth_username}
	AuthUsername *string `field:"optional" json:"authUsername" yaml:"authUsername"`
	// The sender email address. Must be a valid email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.69.0/docs/resources/observability_instance#from ObservabilityInstance#from}
	From *string `field:"optional" json:"from" yaml:"from"`
	// Whether to notify about resolved alerts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.69.0/docs/resources/observability_instance#send_resolved ObservabilityInstance#send_resolved}
	SendResolved interface{} `field:"optional" json:"sendResolved" yaml:"sendResolved"`
	// The SMTP host through which emails are sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.69.0/docs/resources/observability_instance#smart_host ObservabilityInstance#smart_host}
	SmartHost *string `field:"optional" json:"smartHost" yaml:"smartHost"`
	// The email address to send notifications to. Must be a valid email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.69.0/docs/resources/observability_instance#to ObservabilityInstance#to}
	To *string `field:"optional" json:"to" yaml:"to"`
}

