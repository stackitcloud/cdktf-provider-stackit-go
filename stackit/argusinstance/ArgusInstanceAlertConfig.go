package argusinstance


type ArgusInstanceAlertConfig struct {
	// List of alert receivers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.40.0/docs/resources/argus_instance#receivers ArgusInstance#receivers}
	Receivers interface{} `field:"required" json:"receivers" yaml:"receivers"`
	// Route configuration for the alerts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.40.0/docs/resources/argus_instance#route ArgusInstance#route}
	Route *ArgusInstanceAlertConfigRoute `field:"required" json:"route" yaml:"route"`
	// Global configuration for the alerts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.40.0/docs/resources/argus_instance#global ArgusInstance#global}
	Global *ArgusInstanceAlertConfigGlobal `field:"optional" json:"global" yaml:"global"`
}

