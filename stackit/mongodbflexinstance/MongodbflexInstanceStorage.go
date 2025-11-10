package mongodbflexinstance


type MongodbflexInstanceStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.6.2/docs/resources/mongodbflex_instance#class MongodbflexInstance#class}.
	Class *string `field:"required" json:"class" yaml:"class"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.6.2/docs/resources/mongodbflex_instance#size MongodbflexInstance#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
}

