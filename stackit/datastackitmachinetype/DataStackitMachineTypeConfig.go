package datastackitmachinetype

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitMachineTypeConfig struct {
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
	// Expr-lang filter for filtering machine types.
	//
	// Examples:
	// - vcpus == 2
	// - ram >= 2048
	// - extraSpecs.cpu == "intel-icelake-generic"
	// - extraSpecs.cpu == "intel-icelake-generic" && vcpus == 2
	//
	// Syntax reference: https://expr-lang.org/docs/language-definition
	//
	// You can also list available machine-types using the [STACKIT CLI](https://github.com/stackitcloud/stackit-cli):
	//
	// ```bash
	// stackit server machine-type list
	// ```
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/data-sources/machine_type#filter DataStackitMachineType#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// STACKIT Project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/data-sources/machine_type#project_id DataStackitMachineType#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/data-sources/machine_type#region DataStackitMachineType#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Sort machine types by name ascending (`true`) or descending (`false`). Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/data-sources/machine_type#sort_ascending DataStackitMachineType#sort_ascending}
	SortAscending interface{} `field:"optional" json:"sortAscending" yaml:"sortAscending"`
}

