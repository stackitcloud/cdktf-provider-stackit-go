package telemetryrouterdestination


type TelemetryrouterDestinationConfigFilter struct {
	// The TelemetryRouter destination's filter attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/telemetryrouter_destination#attributes TelemetryrouterDestination#attributes}
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
}

