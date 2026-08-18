package skecluster


type SkeClusterExtensionsApplicationLoadBalancer struct {
	// Enables the application load balancer extension.
	//
	// Note: This feature is in private preview. Enabling application load balancer extension is only possible for enabled accounts. Otherwise the request will be rejected. Default value will change to true once the private preview phase is over.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.112.0/docs/resources/ske_cluster#enabled SkeCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

