package sqlserverflexinstance


type SqlserverflexInstanceFlavor struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.32.0/docs/resources/sqlserverflex_instance#cpu SqlserverflexInstance#cpu}.
	Cpu *float64 `field:"required" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.32.0/docs/resources/sqlserverflex_instance#ram SqlserverflexInstance#ram}.
	Ram *float64 `field:"required" json:"ram" yaml:"ram"`
}

