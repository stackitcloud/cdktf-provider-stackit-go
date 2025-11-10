package datastackitloadbalancer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitLoadbalancerConfig struct {
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
	// Load balancer name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.23.0/docs/data-sources/loadbalancer#name DataStackitLoadbalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the Load Balancer is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.23.0/docs/data-sources/loadbalancer#project_id DataStackitLoadbalancer#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

