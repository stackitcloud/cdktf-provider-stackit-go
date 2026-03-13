package skecluster


type SkeClusterNetworkControlPlane struct {
	// Access scope of the control plane.
	//
	// It defines if the Kubernetes control plane is public or only available inside a STACKIT Network Area.Possible values are: `PUBLIC`, `SNA`. The field is immutable!
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/ske_cluster#access_scope SkeCluster#access_scope}
	AccessScope *string `field:"optional" json:"accessScope" yaml:"accessScope"`
}

