package securitygrouprule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityGroupRuleConfig struct {
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
	// The direction of the traffic which the rule should match.
	//
	// Some of the possible values are: Supported values are: `ingress`, `egress`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/security_group_rule#direction SecurityGroupRule#direction}
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// STACKIT project ID to which the security group rule is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/security_group_rule#project_id SecurityGroupRule#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The security group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/security_group_rule#security_group_id SecurityGroupRule#security_group_id}
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// The rule description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/security_group_rule#description SecurityGroupRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ethertype which the rule should match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/security_group_rule#ether_type SecurityGroupRule#ether_type}
	EtherType *string `field:"optional" json:"etherType" yaml:"etherType"`
	// ICMP Parameters. These parameters should only be provided if the protocol is ICMP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/security_group_rule#icmp_parameters SecurityGroupRule#icmp_parameters}
	IcmpParameters *SecurityGroupRuleIcmpParameters `field:"optional" json:"icmpParameters" yaml:"icmpParameters"`
	// The remote IP range which the rule should match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/security_group_rule#ip_range SecurityGroupRule#ip_range}
	IpRange *string `field:"optional" json:"ipRange" yaml:"ipRange"`
	// The range of ports. This should only be provided if the protocol is not ICMP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/security_group_rule#port_range SecurityGroupRule#port_range}
	PortRange *SecurityGroupRulePortRange `field:"optional" json:"portRange" yaml:"portRange"`
	// The internet protocol which the rule should match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/security_group_rule#protocol SecurityGroupRule#protocol}
	Protocol *SecurityGroupRuleProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The remote security group which the rule should match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.35.0/docs/resources/security_group_rule#remote_security_group_id SecurityGroupRule#remote_security_group_id}
	RemoteSecurityGroupId *string `field:"optional" json:"remoteSecurityGroupId" yaml:"remoteSecurityGroupId"`
}

