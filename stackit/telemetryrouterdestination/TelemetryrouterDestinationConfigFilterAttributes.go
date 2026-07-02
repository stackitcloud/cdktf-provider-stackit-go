package telemetryrouterdestination


type TelemetryrouterDestinationConfigFilterAttributes struct {
	// The TelemetryRouter destination's filter attribute key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/telemetryrouter_destination#key TelemetryrouterDestination#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The TelemetryRouter destination's filter attribute level, possible values: Possible values are: `resource`, `scope`, `logRecord`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/telemetryrouter_destination#level TelemetryrouterDestination#level}
	Level *string `field:"required" json:"level" yaml:"level"`
	// The TelemetryRouter destination's filter attribute matcher, possible values: Possible values are: `=`, `!=`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/telemetryrouter_destination#matcher TelemetryrouterDestination#matcher}
	Matcher *string `field:"required" json:"matcher" yaml:"matcher"`
	// The TelemetryRouter destination's filter attribute values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/telemetryrouter_destination#values TelemetryrouterDestination#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

