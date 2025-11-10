package postgresflexinstance


type PostgresflexInstanceFlavor struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/postgresflex_instance#cpu PostgresflexInstance#cpu}.
	Cpu *float64 `field:"required" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/postgresflex_instance#ram PostgresflexInstance#ram}.
	Ram *float64 `field:"required" json:"ram" yaml:"ram"`
}

