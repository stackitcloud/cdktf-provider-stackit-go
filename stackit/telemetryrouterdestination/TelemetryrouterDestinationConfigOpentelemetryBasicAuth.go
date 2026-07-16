package telemetryrouterdestination


type TelemetryrouterDestinationConfigOpentelemetryBasicAuth struct {
	// OpenTelemetry basic auth username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/telemetryrouter_destination#password TelemetryrouterDestination#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// OpenTelemetry basic auth username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/resources/telemetryrouter_destination#username TelemetryrouterDestination#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

