package datastackittelemetryrouterdestination


type DataStackitTelemetryrouterDestinationConfigA struct {
	// The TelemetryRouter destination's filter settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/data-sources/telemetryrouter_destination#filter DataStackitTelemetryrouterDestination#filter}
	Filter *DataStackitTelemetryrouterDestinationConfigFilter `field:"optional" json:"filter" yaml:"filter"`
	// OpenTelemetry configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/data-sources/telemetryrouter_destination#opentelemetry DataStackitTelemetryrouterDestination#opentelemetry}
	Opentelemetry *DataStackitTelemetryrouterDestinationConfigOpentelemetry `field:"optional" json:"opentelemetry" yaml:"opentelemetry"`
	// S3 configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/data-sources/telemetryrouter_destination#s3 DataStackitTelemetryrouterDestination#s3}
	S3 *DataStackitTelemetryrouterDestinationConfigS3 `field:"optional" json:"s3" yaml:"s3"`
}

