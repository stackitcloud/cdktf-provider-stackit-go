package postgresflexinstance


type PostgresflexInstanceStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/postgresflex_instance#class PostgresflexInstance#class}.
	Class *string `field:"required" json:"class" yaml:"class"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/postgresflex_instance#size PostgresflexInstance#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
}

