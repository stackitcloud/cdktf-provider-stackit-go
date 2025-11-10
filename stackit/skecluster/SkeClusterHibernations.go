package skecluster


type SkeClusterHibernations struct {
	// End time of hibernation in crontab syntax.
	//
	// E.g. `0 8 * * *` for waking up the cluster at 8am.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.6.0/docs/resources/ske_cluster#end SkeCluster#end}
	End *string `field:"required" json:"end" yaml:"end"`
	// Start time of cluster hibernation in crontab syntax. E.g. `0 18 * * *` for starting everyday at 6pm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.6.0/docs/resources/ske_cluster#start SkeCluster#start}
	Start *string `field:"required" json:"start" yaml:"start"`
	// Timezone name corresponding to a file in the IANA Time Zone database. i.e. `Europe/Berlin`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.6.0/docs/resources/ske_cluster#timezone SkeCluster#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

