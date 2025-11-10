package server


type ServerBootVolume struct {
	// The ID of the source, either image ID or volume ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.3/docs/resources/server#source_id Server#source_id}
	SourceId *string `field:"required" json:"sourceId" yaml:"sourceId"`
	// The type of the source. Supported values are: `volume`, `image`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.3/docs/resources/server#source_type Server#source_type}
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
	// Delete the volume during the termination of the server. Only allowed when `source_type` is `image`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.3/docs/resources/server#delete_on_termination Server#delete_on_termination}
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// The performance class of the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.3/docs/resources/server#performance_class Server#performance_class}
	PerformanceClass *string `field:"optional" json:"performanceClass" yaml:"performanceClass"`
	// The size of the boot volume in GB. Must be provided when `source_type` is `image`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.3/docs/resources/server#size Server#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

