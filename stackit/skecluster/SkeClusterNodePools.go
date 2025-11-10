package skecluster


type SkeClusterNodePools struct {
	// Specify a list of availability zones. E.g. `eu01-m`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/ske_cluster#availability_zones SkeCluster#availability_zones}
	AvailabilityZones *[]*string `field:"required" json:"availabilityZones" yaml:"availabilityZones"`
	// The machine type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/ske_cluster#machine_type SkeCluster#machine_type}
	MachineType *string `field:"required" json:"machineType" yaml:"machineType"`
	// Maximum number of nodes in the pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/ske_cluster#maximum SkeCluster#maximum}
	Maximum *float64 `field:"required" json:"maximum" yaml:"maximum"`
	// Minimum number of nodes in the pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/ske_cluster#minimum SkeCluster#minimum}
	Minimum *float64 `field:"required" json:"minimum" yaml:"minimum"`
	// Specifies the name of the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/ske_cluster#name SkeCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Allow system components to run on this node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/ske_cluster#allow_system_components SkeCluster#allow_system_components}
	AllowSystemComponents interface{} `field:"optional" json:"allowSystemComponents" yaml:"allowSystemComponents"`
	// Specifies the container runtime. Defaults to `containerd`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/ske_cluster#cri SkeCluster#cri}
	Cri *string `field:"optional" json:"cri" yaml:"cri"`
	// Labels to add to each node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/ske_cluster#labels SkeCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Maximum number of additional VMs that are created during an update.
	//
	// If set (larger than 0), then it must be at least the amount of zones configured for the nodepool. The `max_surge` and `max_unavailable` fields cannot both be unset at the same time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/ske_cluster#max_surge SkeCluster#max_surge}
	MaxSurge *float64 `field:"optional" json:"maxSurge" yaml:"maxSurge"`
	// Maximum number of VMs that that can be unavailable during an update.
	//
	// If set (larger than 0), then it must be at least the amount of zones configured for the nodepool. The `max_surge` and `max_unavailable` fields cannot both be unset at the same time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/ske_cluster#max_unavailable SkeCluster#max_unavailable}
	MaxUnavailable *float64 `field:"optional" json:"maxUnavailable" yaml:"maxUnavailable"`
	// The name of the OS image. Defaults to `flatcar`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/ske_cluster#os_name SkeCluster#os_name}
	OsName *string `field:"optional" json:"osName" yaml:"osName"`
	// This field is deprecated, use `os_version_min` to configure the version and `os_version_used` to get the currently used version instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/ske_cluster#os_version SkeCluster#os_version}
	OsVersion *string `field:"optional" json:"osVersion" yaml:"osVersion"`
	// The minimum OS image version.
	//
	// This field will be used to set the minimum OS image version on creation/update of the cluster. If unset, the latest supported OS image version will be used. SKE automatically updates the cluster Kubernetes version if you have set `maintenance.enable_kubernetes_version_updates` to true or if there is a mandatory update, as described in [Updates for Kubernetes versions and Operating System versions in SKE](https://docs.stackit.cloud/stackit/en/version-updates-in-ske-10125631.html). To get the current OS image version being used for the node pool, use the read-only `os_version_used` field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/ske_cluster#os_version_min SkeCluster#os_version_min}
	OsVersionMin *string `field:"optional" json:"osVersionMin" yaml:"osVersionMin"`
	// Specifies a taint list as defined below.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/ske_cluster#taints SkeCluster#taints}
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
	// The volume size in GB. Defaults to `20`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/ske_cluster#volume_size SkeCluster#volume_size}
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// Specifies the volume type. Defaults to `storage_premium_perf1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.1/docs/resources/ske_cluster#volume_type SkeCluster#volume_type}
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

