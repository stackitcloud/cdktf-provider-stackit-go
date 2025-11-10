package skecluster


type SkeClusterMaintenance struct {
	// Flag to enable/disable auto-updates of the Kubernetes version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.13.0/docs/resources/ske_cluster#enable_kubernetes_version_updates SkeCluster#enable_kubernetes_version_updates}
	EnableKubernetesVersionUpdates interface{} `field:"required" json:"enableKubernetesVersionUpdates" yaml:"enableKubernetesVersionUpdates"`
	// Flag to enable/disable auto-updates of the OS image version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.13.0/docs/resources/ske_cluster#enable_machine_image_version_updates SkeCluster#enable_machine_image_version_updates}
	EnableMachineImageVersionUpdates interface{} `field:"required" json:"enableMachineImageVersionUpdates" yaml:"enableMachineImageVersionUpdates"`
	// Time for maintenance window end. E.g. `01:23:45Z`, `05:00:00+02:00`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.13.0/docs/resources/ske_cluster#end SkeCluster#end}
	End *string `field:"required" json:"end" yaml:"end"`
	// Time for maintenance window start. E.g. `01:23:45Z`, `05:00:00+02:00`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.13.0/docs/resources/ske_cluster#start SkeCluster#start}
	Start *string `field:"required" json:"start" yaml:"start"`
}

