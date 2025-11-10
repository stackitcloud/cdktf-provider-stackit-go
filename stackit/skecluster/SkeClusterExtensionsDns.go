package skecluster


type SkeClusterExtensionsDns struct {
	// Flag to enable/disable DNS extensions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/ske_cluster#enabled SkeCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specify a list of domain filters for externalDNS (e.g., `foo.runs.onstackit.cloud`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/resources/ske_cluster#zones SkeCluster#zones}
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

