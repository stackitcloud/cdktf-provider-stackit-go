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
	// Kubernetes version. Must only contain major and minor version (e.g. 1.22).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.4.0/docs/resources/ske_cluster#kubernetes_version SkeCluster#kubernetes_version}
	KubernetesVersion *string `field:"required" json:"kubernetesVersion" yaml:"kubernetesVersion"`
	// The cluster name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.4.0/docs/resources/ske_cluster#name SkeCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// One or more `node_pool` block as defined below.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.4.0/docs/resources/ske_cluster#node_pools SkeCluster#node_pools}
	NodePools interface{} `field:"required" json:"nodePools" yaml:"nodePools"`
	// STACKIT project ID to which the cluster is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.4.0/docs/resources/ske_cluster#project_id SkeCluster#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Flag to specify if privileged mode for containers is enabled or not.
	//
	// This should be used with care since it also disables a couple of other features like the use of some volume type (e.g. PVCs).
	// Deprecated as of Kubernetes 1.25 and later
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.4.0/docs/resources/ske_cluster#allow_privileged_containers SkeCluster#allow_privileged_containers}
	AllowPrivilegedContainers interface{} `field:"optional" json:"allowPrivilegedContainers" yaml:"allowPrivilegedContainers"`
	// A single extensions block as defined below.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.4.0/docs/resources/ske_cluster#extensions SkeCluster#extensions}
	Extensions *SkeClusterExtensions `field:"optional" json:"extensions" yaml:"extensions"`
	// One or more hibernation block as defined below.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.4.0/docs/resources/ske_cluster#hibernations SkeCluster#hibernations}
	Hibernations interface{} `field:"optional" json:"hibernations" yaml:"hibernations"`
	// A single maintenance block as defined below.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.4.0/docs/resources/ske_cluster#maintenance SkeCluster#maintenance}
	Maintenance *SkeClusterMaintenance `field:"optional" json:"maintenance" yaml:"maintenance"`
}

