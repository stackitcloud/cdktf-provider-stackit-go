package skecluster


type SkeClusterAccessIdp struct {
	// Enable IDP integration for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/ske_cluster#enabled SkeCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The IDP type. Possible values: 'stackit'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.104.0/docs/resources/ske_cluster#type SkeCluster#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

