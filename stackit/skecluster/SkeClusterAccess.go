package skecluster


type SkeClusterAccess struct {
	// Configure IDP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/ske_cluster#idp SkeCluster#idp}
	Idp *SkeClusterAccessIdp `field:"optional" json:"idp" yaml:"idp"`
}

