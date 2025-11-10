package networkarea

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkAreaConfig struct {
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
	// The name of the network area.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.38.1/docs/resources/network_area#name NetworkArea#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of Network ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.38.1/docs/resources/network_area#network_ranges NetworkArea#network_ranges}
	NetworkRanges interface{} `field:"required" json:"networkRanges" yaml:"networkRanges"`
	// STACKIT organization ID to which the network area is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.38.1/docs/resources/network_area#organization_id NetworkArea#organization_id}
	OrganizationId *string `field:"required" json:"organizationId" yaml:"organizationId"`
	// Classless Inter-Domain Routing (CIDR).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.38.1/docs/resources/network_area#transfer_network NetworkArea#transfer_network}
	TransferNetwork *string `field:"required" json:"transferNetwork" yaml:"transferNetwork"`
	// List of DNS Servers/Nameservers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.38.1/docs/resources/network_area#default_nameservers NetworkArea#default_nameservers}
	DefaultNameservers *[]*string `field:"optional" json:"defaultNameservers" yaml:"defaultNameservers"`
	// The default prefix length for networks in the network area.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.38.1/docs/resources/network_area#default_prefix_length NetworkArea#default_prefix_length}
	DefaultPrefixLength *float64 `field:"optional" json:"defaultPrefixLength" yaml:"defaultPrefixLength"`
	// Labels are key-value string pairs which can be attached to a resource container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.38.1/docs/resources/network_area#labels NetworkArea#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The maximal prefix length for networks in the network area.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.38.1/docs/resources/network_area#max_prefix_length NetworkArea#max_prefix_length}
	MaxPrefixLength *float64 `field:"optional" json:"maxPrefixLength" yaml:"maxPrefixLength"`
	// The minimal prefix length for networks in the network area.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.38.1/docs/resources/network_area#min_prefix_length NetworkArea#min_prefix_length}
	MinPrefixLength *float64 `field:"optional" json:"minPrefixLength" yaml:"minPrefixLength"`
}

