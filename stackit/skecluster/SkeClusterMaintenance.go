package skecluster


type SkeClusterMaintenance struct {
	// Time for maintenance window end. E.g. `01:23:45Z`, `05:00:00+02:00`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.76.0/docs/resources/ske_cluster#end SkeCluster#end}
	End *string `field:"required" json:"end" yaml:"end"`
	// Time for maintenance window start. E.g. `01:23:45Z`, `05:00:00+02:00`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.76.0/docs/resources/ske_cluster#start SkeCluster#start}
	Start *string `field:"required" json:"start" yaml:"start"`
	// Flag to enable/disable auto-updates of the Kubernetes version.
	//
	// Defaults to `true`. SKE automatically updates the cluster Kubernetes version if you have set `maintenance.enable_kubernetes_version_updates` to true or if there is a mandatory update, as described in [General information for Kubernetes & OS updates](https://docs.stackit.cloud/products/runtime/kubernetes-engine/basics/version-updates/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.76.0/docs/resources/ske_cluster#enable_kubernetes_version_updates SkeCluster#enable_kubernetes_version_updates}
	EnableKubernetesVersionUpdates interface{} `field:"optional" json:"enableKubernetesVersionUpdates" yaml:"enableKubernetesVersionUpdates"`
	// Flag to enable/disable auto-updates of the OS image version.
	//
	// Defaults to `true`. SKE automatically updates the cluster Kubernetes version if you have set `maintenance.enable_kubernetes_version_updates` to true or if there is a mandatory update, as described in [General information for Kubernetes & OS updates](https://docs.stackit.cloud/products/runtime/kubernetes-engine/basics/version-updates/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.76.0/docs/resources/ske_cluster#enable_machine_image_version_updates SkeCluster#enable_machine_image_version_updates}
	EnableMachineImageVersionUpdates interface{} `field:"optional" json:"enableMachineImageVersionUpdates" yaml:"enableMachineImageVersionUpdates"`
}

