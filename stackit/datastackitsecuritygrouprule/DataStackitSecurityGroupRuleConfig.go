package datastackitsecuritygrouprule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitSecurityGroupRuleConfig struct {
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
	// STACKIT project ID to which the security group rule is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.3/docs/data-sources/security_group_rule#project_id DataStackitSecurityGroupRule#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The security group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.3/docs/data-sources/security_group_rule#security_group_id DataStackitSecurityGroupRule#security_group_id}
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// The security group rule ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.3/docs/data-sources/security_group_rule#security_group_rule_id DataStackitSecurityGroupRule#security_group_rule_id}
	SecurityGroupRuleId *string `field:"required" json:"securityGroupRuleId" yaml:"securityGroupRuleId"`
}

