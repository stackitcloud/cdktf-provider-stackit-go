package telemetryrouterdestination


type TelemetryrouterDestinationConfigOpentelemetry struct {
	// OpenTelemetry destination URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/telemetryrouter_destination#uri TelemetryrouterDestination#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// OpenTelemetry basic auth configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/telemetryrouter_destination#basic_auth TelemetryrouterDestination#basic_auth}
	BasicAuth *TelemetryrouterDestinationConfigOpentelemetryBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	// OpenTelemetry bearer token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/telemetryrouter_destination#bearer_token TelemetryrouterDestination#bearer_token}
	BearerToken *string `field:"optional" json:"bearerToken" yaml:"bearerToken"`
}

