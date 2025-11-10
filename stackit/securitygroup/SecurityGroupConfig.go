package securitygroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityGroupConfig struct {
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
	// The name of the security group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.62.0/docs/resources/security_group#name SecurityGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the security group is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.62.0/docs/resources/security_group#project_id SecurityGroup#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The description of the security group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.62.0/docs/resources/security_group#description SecurityGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Labels are key-value string pairs which can be attached to a resource container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.62.0/docs/resources/security_group#labels SecurityGroup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Configures if a security group is stateful or stateless.
	//
	// There can only be one type of security groups per network interface/server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.62.0/docs/resources/security_group#stateful SecurityGroup#stateful}
	Stateful interface{} `field:"optional" json:"stateful" yaml:"stateful"`
}

