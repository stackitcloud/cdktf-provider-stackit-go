package loadbalancer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerConfig struct {
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
	// List of all listeners which will accept traffic. Limited to 20.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.3/docs/resources/loadbalancer#listeners Loadbalancer#listeners}
	Listeners interface{} `field:"required" json:"listeners" yaml:"listeners"`
	// Load balancer name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.3/docs/resources/loadbalancer#name Loadbalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of networks that listeners and targets reside in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.3/docs/resources/loadbalancer#networks Loadbalancer#networks}
	Networks interface{} `field:"required" json:"networks" yaml:"networks"`
	// STACKIT project ID to which the Load Balancer is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.3/docs/resources/loadbalancer#project_id Loadbalancer#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// List of all target pools which will be used in the Load Balancer. Limited to 20.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.3/docs/resources/loadbalancer#target_pools Loadbalancer#target_pools}
	TargetPools interface{} `field:"required" json:"targetPools" yaml:"targetPools"`
	// External Load Balancer IP address where this Load Balancer is exposed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.3/docs/resources/loadbalancer#external_address Loadbalancer#external_address}
	ExternalAddress *string `field:"optional" json:"externalAddress" yaml:"externalAddress"`
	// Defines any optional functionality you want to have enabled on your load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.3/docs/resources/loadbalancer#options Loadbalancer#options}
	Options *LoadbalancerOptions `field:"optional" json:"options" yaml:"options"`
}

