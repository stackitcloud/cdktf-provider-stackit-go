package skecluster


type SkeClusterExtensions struct {
	// Cluster access control configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/ske_cluster#acl SkeCluster#acl}
	Acl *SkeClusterExtensionsAcl `field:"optional" json:"acl" yaml:"acl"`
	// A single argus block as defined below. This field is deprecated and will be removed 06 January 2026.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/ske_cluster#argus SkeCluster#argus}
	Argus *SkeClusterExtensionsArgus `field:"optional" json:"argus" yaml:"argus"`
	// DNS extension configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/ske_cluster#dns SkeCluster#dns}
	Dns *SkeClusterExtensionsDns `field:"optional" json:"dns" yaml:"dns"`
	// A single observability block as defined below.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/ske_cluster#observability SkeCluster#observability}
	Observability *SkeClusterExtensionsObservability `field:"optional" json:"observability" yaml:"observability"`
}

