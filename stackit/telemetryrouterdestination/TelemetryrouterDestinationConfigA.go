package telemetryrouterdestination


type TelemetryrouterDestinationConfigA struct {
	// The TelemetryRouter destinations's configuration type, possible values: Possible values are: `OpenTelemetry`, `S3`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/telemetryrouter_destination#config_type TelemetryrouterDestination#config_type}
	ConfigType *string `field:"required" json:"configType" yaml:"configType"`
	// The TelemetryRouter destination's filter settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/telemetryrouter_destination#filter TelemetryrouterDestination#filter}
	Filter *TelemetryrouterDestinationConfigFilter `field:"optional" json:"filter" yaml:"filter"`
	// OpenTelemetry configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/telemetryrouter_destination#opentelemetry TelemetryrouterDestination#opentelemetry}
	Opentelemetry *TelemetryrouterDestinationConfigOpentelemetry `field:"optional" json:"opentelemetry" yaml:"opentelemetry"`
	// S3 configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/telemetryrouter_destination#s3 TelemetryrouterDestination#s3}
	S3 *TelemetryrouterDestinationConfigS3 `field:"optional" json:"s3" yaml:"s3"`
}

