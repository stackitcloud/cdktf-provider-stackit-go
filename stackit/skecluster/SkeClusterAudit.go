package skecluster


type SkeClusterAudit struct {
	// Enable cluster audit log forwarding to a Telemetry Router.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.112.0/docs/resources/ske_cluster#enabled SkeCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

