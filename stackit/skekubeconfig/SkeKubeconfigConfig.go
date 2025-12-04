package skekubeconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SkeKubeconfigConfig struct {
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
	// Name of the SKE cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.73.0/docs/resources/ske_kubeconfig#cluster_name SkeKubeconfig#cluster_name}
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// STACKIT project ID to which the cluster is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.73.0/docs/resources/ske_kubeconfig#project_id SkeKubeconfig#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Expiration time of the kubeconfig, in seconds. Defaults to `3600`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.73.0/docs/resources/ske_kubeconfig#expiration SkeKubeconfig#expiration}
	Expiration *float64 `field:"optional" json:"expiration" yaml:"expiration"`
	// If set to true, the provider will check if the kubeconfig has expired and will generated a new valid one in-place.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.73.0/docs/resources/ske_kubeconfig#refresh SkeKubeconfig#refresh}
	Refresh interface{} `field:"optional" json:"refresh" yaml:"refresh"`
	// Number of seconds before expiration to trigger refresh of the kubeconfig at.
	//
	// Only used if refresh is set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.73.0/docs/resources/ske_kubeconfig#refresh_before SkeKubeconfig#refresh_before}
	RefreshBefore *float64 `field:"optional" json:"refreshBefore" yaml:"refreshBefore"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.73.0/docs/resources/ske_kubeconfig#region SkeKubeconfig#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

