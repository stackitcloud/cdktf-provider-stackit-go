package datastackitmodelexperimentsinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitModelexperimentsInstanceConfig struct {
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
	// The AI Model Experiments instance ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/data-sources/modelexperiments_instance#instance_id DataStackitModelexperimentsInstance#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// STACKIT Project ID to which the resource is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/data-sources/modelexperiments_instance#project_id DataStackitModelexperimentsInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The STACKIT region name the resource is located in. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/data-sources/modelexperiments_instance#region DataStackitModelexperimentsInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

