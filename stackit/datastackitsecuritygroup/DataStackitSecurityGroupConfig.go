package datastackitsecuritygroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitSecurityGroupConfig struct {
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
	// STACKIT project ID to which the security group is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.38.0/docs/data-sources/security_group#project_id DataStackitSecurityGroup#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The security group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.38.0/docs/data-sources/security_group#security_group_id DataStackitSecurityGroup#security_group_id}
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
}

