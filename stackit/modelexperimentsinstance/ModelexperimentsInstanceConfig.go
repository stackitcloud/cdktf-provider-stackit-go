package modelexperimentsinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelexperimentsInstanceConfig struct {
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
	// The display name is a short name chosen by the user to identify the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.112.0/docs/resources/modelexperiments_instance#name ModelexperimentsInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT Project ID to which the resource is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.112.0/docs/resources/modelexperiments_instance#project_id ModelexperimentsInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The deleted experiment retention time of the AI Model Experiments instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.112.0/docs/resources/modelexperiments_instance#deleted_experiment_retention ModelexperimentsInstance#deleted_experiment_retention}
	DeletedExperimentRetention *string `field:"optional" json:"deletedExperimentRetention" yaml:"deletedExperimentRetention"`
	// The description is a longer text chosen by the user to provide more context for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.112.0/docs/resources/modelexperiments_instance#description ModelexperimentsInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A map of arbitrary key/value pairs that can be attached to the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.112.0/docs/resources/modelexperiments_instance#labels ModelexperimentsInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The STACKIT region name the resource is located in. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.112.0/docs/resources/modelexperiments_instance#region ModelexperimentsInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

