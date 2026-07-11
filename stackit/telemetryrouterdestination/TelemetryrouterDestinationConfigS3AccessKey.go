package telemetryrouterdestination


type TelemetryrouterDestinationConfigS3AccessKey struct {
	// S3 access key ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/telemetryrouter_destination#id TelemetryrouterDestination#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// S3 access key secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/telemetryrouter_destination#secret TelemetryrouterDestination#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
}

