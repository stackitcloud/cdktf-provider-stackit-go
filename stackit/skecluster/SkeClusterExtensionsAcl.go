package skecluster


type SkeClusterExtensionsAcl struct {
	// Is ACL enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.21.1/docs/resources/ske_cluster#enabled SkeCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specify a list of CIDRs to whitelist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.21.1/docs/resources/ske_cluster#allowed_cidrs SkeCluster#allowed_cidrs}
	AllowedCidrs *[]*string `field:"optional" json:"allowedCidrs" yaml:"allowedCidrs"`
}

