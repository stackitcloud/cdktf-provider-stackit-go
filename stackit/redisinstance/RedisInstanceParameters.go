package redisinstance


type RedisInstanceParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.14.1/docs/resources/redis_instance#sgw_acl RedisInstance#sgw_acl}.
	SgwAcl *string `field:"optional" json:"sgwAcl" yaml:"sgwAcl"`
}

