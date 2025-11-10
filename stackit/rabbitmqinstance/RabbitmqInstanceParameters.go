package rabbitmqinstance


type RabbitmqInstanceParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.6.8/docs/resources/rabbitmq_instance#sgw_acl RabbitmqInstance#sgw_acl}.
	SgwAcl *string `field:"optional" json:"sgwAcl" yaml:"sgwAcl"`
}

