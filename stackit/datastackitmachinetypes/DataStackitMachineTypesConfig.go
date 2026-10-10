package datastackitmachinetypes

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitMachineTypesConfig struct {
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
	// STACKIT project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.119.0/docs/data-sources/machine_types#project_id DataStackitMachineTypes#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Experimental API-side equality filters, which may be subject to breaking changes.
	//
	// All configured attributes and extra spec entries are combined with AND. Omit this object or use an empty object to list all machine types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.119.0/docs/data-sources/machine_types#filter DataStackitMachineTypes#filter}
	Filter *DataStackitMachineTypesFilter `field:"optional" json:"filter" yaml:"filter"`
	// Region override. Uses the provider region when omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.119.0/docs/data-sources/machine_types#region DataStackitMachineTypes#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.119.0/docs/data-sources/machine_types#timeouts DataStackitMachineTypes#timeouts}.
	Timeouts *DataStackitMachineTypesTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

