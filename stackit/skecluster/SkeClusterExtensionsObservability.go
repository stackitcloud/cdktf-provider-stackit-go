package skecluster


type SkeClusterExtensionsObservability struct {
	// Flag to enable/disable Observability extensions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.69.0/docs/resources/ske_cluster#enabled SkeCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Observability instance ID to choose which Observability instance is used. Required when enabled is set to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.69.0/docs/resources/ske_cluster#instance_id SkeCluster#instance_id}
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
}

