package skecluster


type SkeClusterExtensionsAcl struct {
	// Specify a list of CIDRs to whitelist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.2/docs/resources/ske_cluster#allowed_cidrs SkeCluster#allowed_cidrs}
	AllowedCidrs *[]*string `field:"required" json:"allowedCidrs" yaml:"allowedCidrs"`
	// Is ACL enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.2/docs/resources/ske_cluster#enabled SkeCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

