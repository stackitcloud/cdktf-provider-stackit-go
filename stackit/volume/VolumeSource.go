package volume


type VolumeSource struct {
	// The ID of the source, e.g. image ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/resources/volume#id Volume#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The type of the source. Possible values are: `volume`, `image`, `snapshot`, `backup`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/resources/volume#type Volume#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

