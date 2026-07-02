package telemetryrouterdestination


type TelemetryrouterDestinationConfigS3 struct {
	// S3 bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/telemetryrouter_destination#bucket TelemetryrouterDestination#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// S3 endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/telemetryrouter_destination#endpoint TelemetryrouterDestination#endpoint}
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// S3 access key configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/telemetryrouter_destination#access_key TelemetryrouterDestination#access_key}
	AccessKey *TelemetryrouterDestinationConfigS3AccessKey `field:"optional" json:"accessKey" yaml:"accessKey"`
}

