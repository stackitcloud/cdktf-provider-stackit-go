package postgresflexinstance


type PostgresflexInstanceStorage struct {
	// The storage class. You can list available storage classes using the [STACKIT CLI](https://github.com/stackitcloud/stackit-cli): ```bash stackit postgresflex options --storages --flavor-id FLAVOR_ID ```.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.99.0/docs/resources/postgresflex_instance#class PostgresflexInstance#class}
	Class *string `field:"required" json:"class" yaml:"class"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.99.0/docs/resources/postgresflex_instance#size PostgresflexInstance#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
}

