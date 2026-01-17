package datastackitedgecloudinstances

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitEdgecloudInstancesConfig struct {
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
	// STACKIT project ID to which the Edge Cloud instances are associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/data-sources/edgecloud_instances#project_id DataStackitEdgecloudInstances#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.78.1/docs/data-sources/edgecloud_instances#region DataStackitEdgecloudInstances#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

