package datastackitapplicationloadbalancer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitApplicationLoadBalancerConfig struct {
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
	// Application Load balancer name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.85.0/docs/data-sources/application_load_balancer#name DataStackitApplicationLoadBalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the Application Load Balancer is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.85.0/docs/data-sources/application_load_balancer#project_id DataStackitApplicationLoadBalancer#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

