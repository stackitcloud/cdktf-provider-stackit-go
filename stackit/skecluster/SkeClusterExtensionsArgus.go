package skecluster


type SkeClusterExtensionsArgus struct {
	// Argus instance ID to choose which Argus instance is used. Required when enabled is set to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.27.0/docs/resources/ske_cluster#argus_instance_id SkeCluster#argus_instance_id}
	ArgusInstanceId *string `field:"required" json:"argusInstanceId" yaml:"argusInstanceId"`
	// Flag to enable/disable Argus extensions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.27.0/docs/resources/ske_cluster#enabled SkeCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

