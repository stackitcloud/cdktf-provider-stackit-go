package telemetryrouterinstance


type TelemetryrouterInstanceFilterAttributes struct {
	// The TelemetryRouter global filter attribute key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/telemetryrouter_instance#key TelemetryrouterInstance#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The TelemetryRouter global filter attribute level, possible values: Possible values are: `resource`, `scope`, `logRecord`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/telemetryrouter_instance#level TelemetryrouterInstance#level}
	Level *string `field:"required" json:"level" yaml:"level"`
	// The TelemetryRouter global filter attribute matcher, possible values: Possible values are: `=`, `!=`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/telemetryrouter_instance#matcher TelemetryrouterInstance#matcher}
	Matcher *string `field:"required" json:"matcher" yaml:"matcher"`
	// The TelemetryRouter global filter attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/telemetryrouter_instance#values TelemetryrouterInstance#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

