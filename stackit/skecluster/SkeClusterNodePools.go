package skecluster


type SkeClusterNodePools struct {
	// Specify a list of availability zones. E.g. `eu01-m`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.12.0/docs/resources/ske_cluster#availability_zones SkeCluster#availability_zones}
	AvailabilityZones *[]*string `field:"required" json:"availabilityZones" yaml:"availabilityZones"`
	// The machine type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.12.0/docs/resources/ske_cluster#machine_type SkeCluster#machine_type}
	MachineType *string `field:"required" json:"machineType" yaml:"machineType"`
	// Maximum number of nodes in the pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.12.0/docs/resources/ske_cluster#maximum SkeCluster#maximum}
	Maximum *float64 `field:"required" json:"maximum" yaml:"maximum"`
	// Minimum number of nodes in the pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.12.0/docs/resources/ske_cluster#minimum SkeCluster#minimum}
	Minimum *float64 `field:"required" json:"minimum" yaml:"minimum"`
	// Specifies the name of the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.12.0/docs/resources/ske_cluster#name SkeCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The OS image version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.12.0/docs/resources/ske_cluster#os_version SkeCluster#os_version}
	OsVersion *string `field:"required" json:"osVersion" yaml:"osVersion"`
	// Specifies the container runtime. E.g. `containerd`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.12.0/docs/resources/ske_cluster#cri SkeCluster#cri}
	Cri *string `field:"optional" json:"cri" yaml:"cri"`
	// Labels to add to each node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.12.0/docs/resources/ske_cluster#labels SkeCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Maximum number of additional VMs that are created during an update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.12.0/docs/resources/ske_cluster#max_surge SkeCluster#max_surge}
	MaxSurge *float64 `field:"optional" json:"maxSurge" yaml:"maxSurge"`
	// Maximum number of VMs that that can be unavailable during an update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.12.0/docs/resources/ske_cluster#max_unavailable SkeCluster#max_unavailable}
	MaxUnavailable *float64 `field:"optional" json:"maxUnavailable" yaml:"maxUnavailable"`
	// The name of the OS image. E.g. `flatcar`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.12.0/docs/resources/ske_cluster#os_name SkeCluster#os_name}
	OsName *string `field:"optional" json:"osName" yaml:"osName"`
	// Specifies a taint list as defined below.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.12.0/docs/resources/ske_cluster#taints SkeCluster#taints}
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
	// The volume size in GB. E.g. `20`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.12.0/docs/resources/ske_cluster#volume_size SkeCluster#volume_size}
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// Specifies the volume type. E.g. `storage_premium_perf1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.12.0/docs/resources/ske_cluster#volume_type SkeCluster#volume_type}
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

