package skecluster


type SkeClusterExtensions struct {
	// Cluster access control configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.20.1/docs/resources/ske_cluster#acl SkeCluster#acl}
	Acl *SkeClusterExtensionsAcl `field:"optional" json:"acl" yaml:"acl"`
	// A single argus block as defined below.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.20.1/docs/resources/ske_cluster#argus SkeCluster#argus}
	Argus *SkeClusterExtensionsArgus `field:"optional" json:"argus" yaml:"argus"`
}

