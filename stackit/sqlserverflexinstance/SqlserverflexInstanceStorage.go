package sqlserverflexinstance


type SqlserverflexInstanceStorage struct {
	// The storage class. You can list available storage classes using the [STACKIT CLI](https://github.com/stackitcloud/stackit-cli): ```bash stackit beta sqlserverflex options --storages --flavor-id FLAVOR_ID ```.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.99.0/docs/resources/sqlserverflex_instance#class SqlserverflexInstance#class}
	Class *string `field:"optional" json:"class" yaml:"class"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.99.0/docs/resources/sqlserverflex_instance#size SqlserverflexInstance#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

