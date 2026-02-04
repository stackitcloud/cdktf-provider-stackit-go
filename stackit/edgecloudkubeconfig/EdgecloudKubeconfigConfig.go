package edgecloudkubeconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EdgecloudKubeconfigConfig struct {
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
	// STACKIT project ID to which the Edge Cloud instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.80.0/docs/resources/edgecloud_kubeconfig#project_id EdgecloudKubeconfig#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Expiration time of the kubeconfig, in seconds. Minimum is 600, Maximum is 15552000. Defaults to `3600`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.80.0/docs/resources/edgecloud_kubeconfig#expiration EdgecloudKubeconfig#expiration}
	Expiration *float64 `field:"optional" json:"expiration" yaml:"expiration"`
	// ID of the Edge Cloud instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.80.0/docs/resources/edgecloud_kubeconfig#instance_id EdgecloudKubeconfig#instance_id}
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// Name of the Edge Cloud instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.80.0/docs/resources/edgecloud_kubeconfig#instance_name EdgecloudKubeconfig#instance_name}
	InstanceName *string `field:"optional" json:"instanceName" yaml:"instanceName"`
	// Number of seconds before expiration to trigger recreation of the kubeconfig at.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.80.0/docs/resources/edgecloud_kubeconfig#recreate_before EdgecloudKubeconfig#recreate_before}
	RecreateBefore *float64 `field:"optional" json:"recreateBefore" yaml:"recreateBefore"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.80.0/docs/resources/edgecloud_kubeconfig#region EdgecloudKubeconfig#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

