package skecluster


type SkeClusterNodePoolsTaints struct {
	// The taint effect. E.g `PreferNoSchedule`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.65.0/docs/resources/ske_cluster#effect SkeCluster#effect}
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Taint key to be applied to a node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.65.0/docs/resources/ske_cluster#key SkeCluster#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Taint value corresponding to the taint key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.65.0/docs/resources/ske_cluster#value SkeCluster#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

