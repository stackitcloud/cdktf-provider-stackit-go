package observabilityinstance


type ObservabilityInstanceAlertConfig struct {
	// List of alert receivers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/observability_instance#receivers ObservabilityInstance#receivers}
	Receivers interface{} `field:"required" json:"receivers" yaml:"receivers"`
	// Route configuration for the alerts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/observability_instance#route ObservabilityInstance#route}
	Route *ObservabilityInstanceAlertConfigRoute `field:"required" json:"route" yaml:"route"`
	// Global configuration for the alerts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/observability_instance#global ObservabilityInstance#global}
	Global *ObservabilityInstanceAlertConfigGlobal `field:"optional" json:"global" yaml:"global"`
}

