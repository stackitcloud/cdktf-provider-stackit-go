package argusinstance


type ArgusInstanceAlertConfigReceiversEmailConfigs struct {
	// SMTP authentication information. Must be a valid email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.32.0/docs/resources/argus_instance#auth_identity ArgusInstance#auth_identity}
	AuthIdentity *string `field:"optional" json:"authIdentity" yaml:"authIdentity"`
	// SMTP authentication password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.32.0/docs/resources/argus_instance#auth_password ArgusInstance#auth_password}
	AuthPassword *string `field:"optional" json:"authPassword" yaml:"authPassword"`
	// SMTP authentication username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.32.0/docs/resources/argus_instance#auth_username ArgusInstance#auth_username}
	AuthUsername *string `field:"optional" json:"authUsername" yaml:"authUsername"`
	// The sender email address. Must be a valid email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.32.0/docs/resources/argus_instance#from ArgusInstance#from}
	From *string `field:"optional" json:"from" yaml:"from"`
	// The SMTP host through which emails are sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.32.0/docs/resources/argus_instance#smart_host ArgusInstance#smart_host}
	SmartHost *string `field:"optional" json:"smartHost" yaml:"smartHost"`
	// The email address to send notifications to. Must be a valid email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.32.0/docs/resources/argus_instance#to ArgusInstance#to}
	To *string `field:"optional" json:"to" yaml:"to"`
}

