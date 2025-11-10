package skecluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SkeClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The cluster name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.45.0/docs/resources/ske_cluster#name SkeCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// One or more `node_pool` block as defined below.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.45.0/docs/resources/ske_cluster#node_pools SkeCluster#node_pools}
	NodePools interface{} `field:"required" json:"nodePools" yaml:"nodePools"`
	// STACKIT project ID to which the cluster is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.45.0/docs/resources/ske_cluster#project_id SkeCluster#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Flag to specify if privileged mode for containers is enabled or not.
	//
	// This should be used with care since it also disables a couple of other features like the use of some volume type (e.g. PVCs).
	// Deprecated as of Kubernetes 1.25 and later
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.45.0/docs/resources/ske_cluster#allow_privileged_containers SkeCluster#allow_privileged_containers}
	AllowPrivilegedContainers interface{} `field:"optional" json:"allowPrivilegedContainers" yaml:"allowPrivilegedContainers"`
	// A single extensions block as defined below.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.45.0/docs/resources/ske_cluster#extensions SkeCluster#extensions}
	Extensions *SkeClusterExtensions `field:"optional" json:"extensions" yaml:"extensions"`
	// One or more hibernation block as defined below.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.45.0/docs/resources/ske_cluster#hibernations SkeCluster#hibernations}
	Hibernations interface{} `field:"optional" json:"hibernations" yaml:"hibernations"`
	// Kubernetes version. Must only contain major and minor version (e.g. 1.22). This field is deprecated, use `kubernetes_version_min instead`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.45.0/docs/resources/ske_cluster#kubernetes_version SkeCluster#kubernetes_version}
	KubernetesVersion *string `field:"optional" json:"kubernetesVersion" yaml:"kubernetesVersion"`
	// The minimum Kubernetes version.
	//
	// This field will be used to set the minimum kubernetes version on creation/update of the cluster. If unset, the latest supported Kubernetes version will be used. SKE automatically updates the cluster Kubernetes version if you have set `maintenance.enable_kubernetes_version_updates` to true or if there is a mandatory update, as described in [Updates for Kubernetes versions and Operating System versions in SKE](https://docs.stackit.cloud/stackit/en/version-updates-in-ske-10125631.html). To get the current kubernetes version being used for your cluster, use the read-only `kubernetes_version_used` field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.45.0/docs/resources/ske_cluster#kubernetes_version_min SkeCluster#kubernetes_version_min}
	KubernetesVersionMin *string `field:"optional" json:"kubernetesVersionMin" yaml:"kubernetesVersionMin"`
	// A single maintenance block as defined below.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.45.0/docs/resources/ske_cluster#maintenance SkeCluster#maintenance}
	Maintenance *SkeClusterMaintenance `field:"optional" json:"maintenance" yaml:"maintenance"`
	// Network block as defined below.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.45.0/docs/resources/ske_cluster#network SkeCluster#network}
	Network *SkeClusterNetwork `field:"optional" json:"network" yaml:"network"`
}

