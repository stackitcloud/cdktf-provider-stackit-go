package mongodbflexinstance


type MongodbflexInstanceStorage struct {
	// The storage class. You can list available storage classes using the [STACKIT CLI](https://github.com/stackitcloud/stackit-cli): ```bash stackit mongodbflex options --storages --flavor-id FLAVOR_ID ```.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/mongodbflex_instance#class MongodbflexInstance#class}
	Class *string `field:"required" json:"class" yaml:"class"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/resources/mongodbflex_instance#size MongodbflexInstance#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
}

