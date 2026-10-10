package datastackitmachinetypes


type DataStackitMachineTypesFilter struct {
	// Exact disk size in GB to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.119.0/docs/data-sources/machine_types#disk DataStackitMachineTypes#disk}
	Disk *float64 `field:"optional" json:"disk" yaml:"disk"`
	// Extra specs to match by key and exact value (e.g., cpu or overcommit).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.119.0/docs/data-sources/machine_types#extra_specs DataStackitMachineTypes#extra_specs}
	ExtraSpecs *map[string]*string `field:"optional" json:"extraSpecs" yaml:"extraSpecs"`
	// Exact machine type name to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.119.0/docs/data-sources/machine_types#name DataStackitMachineTypes#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Exact RAM size in MB to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.119.0/docs/data-sources/machine_types#ram DataStackitMachineTypes#ram}
	Ram *float64 `field:"optional" json:"ram" yaml:"ram"`
	// Exact number of vCPUs to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.119.0/docs/data-sources/machine_types#vcpu DataStackitMachineTypes#vcpu}
	Vcpu *float64 `field:"optional" json:"vcpu" yaml:"vcpu"`
}

