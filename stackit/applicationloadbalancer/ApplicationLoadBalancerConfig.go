package applicationloadbalancer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationLoadBalancerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#listeners ApplicationLoadBalancer#listeners}
	Listeners interface{} `field:"required" json:"listeners" yaml:"listeners"`
	// Application Load balancer name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#name ApplicationLoadBalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of networks that listeners and targets reside in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#networks ApplicationLoadBalancer#networks}
	Networks interface{} `field:"required" json:"networks" yaml:"networks"`
	// Service Plan configures the size of the Application Load Balancer e.g. 'p10'. See available plans via STACKIT CLI 'stackit beta alb plans' or API https://docs.api.stackit.cloud/documentation/alb/version/v2#tag/Project/operation/APIService_ListPlans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#plan_id ApplicationLoadBalancer#plan_id}
	PlanId *string `field:"required" json:"planId" yaml:"planId"`
	// STACKIT project ID to which the Application Load Balancer is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#project_id ApplicationLoadBalancer#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// List of all target pools which will be used in the Application Load Balancer. Limited to 20.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#target_pools ApplicationLoadBalancer#target_pools}
	TargetPools interface{} `field:"required" json:"targetPools" yaml:"targetPools"`
	// Disable target security group assignemt to allow targets outside of the given network.
	//
	// Connectivity to targets need to be ensured by the customer, including routing and Security Groups (targetSecurityGroup can be assigned). Not changeable after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#disable_target_security_group_assignment ApplicationLoadBalancer#disable_target_security_group_assignment}
	DisableTargetSecurityGroupAssignment interface{} `field:"optional" json:"disableTargetSecurityGroupAssignment" yaml:"disableTargetSecurityGroupAssignment"`
	// The external IP address where this Application Load Balancer is exposed. Not changeable after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#external_address ApplicationLoadBalancer#external_address}
	ExternalAddress *string `field:"optional" json:"externalAddress" yaml:"externalAddress"`
	// Labels represent user-defined metadata as key-value pairs. Label count cannot exceed 64 per ALB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#labels ApplicationLoadBalancer#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Defines any optional functionality you want to have enabled on your Application Load Balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#options ApplicationLoadBalancer#options}
	Options *ApplicationLoadBalancerOptions `field:"optional" json:"options" yaml:"options"`
	// The resource region (e.g. eu01). If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/resources/application_load_balancer#region ApplicationLoadBalancer#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

