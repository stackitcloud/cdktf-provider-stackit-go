package mongodbflexinstance


type MongodbflexInstanceFlavor struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs/resources/mongodbflex_instance#cpu MongodbflexInstance#cpu}.
	Cpu *float64 `field:"required" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs/resources/mongodbflex_instance#ram MongodbflexInstance#ram}.
	Ram *float64 `field:"required" json:"ram" yaml:"ram"`
}

