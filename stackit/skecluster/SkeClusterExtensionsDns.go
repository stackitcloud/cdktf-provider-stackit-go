package skecluster


type SkeClusterExtensionsDns struct {
	// Flag to enable/disable DNS extensions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.113.0/docs/resources/ske_cluster#enabled SkeCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Enables Gateway API support for ExternalDNS.
	//
	// The CRDs must be installed by the user. Once installed, ExternalDNS will be configured at the next cluster reconcile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.113.0/docs/resources/ske_cluster#gateway_api SkeCluster#gateway_api}
	GatewayApi interface{} `field:"optional" json:"gatewayApi" yaml:"gatewayApi"`
	// Specify a list of domain filters for externalDNS (e.g., `foo.runs.onstackit.cloud`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.113.0/docs/resources/ske_cluster#zones SkeCluster#zones}
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

