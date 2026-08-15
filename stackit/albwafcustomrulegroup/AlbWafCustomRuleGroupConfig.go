package albwafcustomrulegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbWafCustomRuleGroupConfig struct {
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
	// Custom rule group configuration name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/alb_waf_custom_rule_group#name AlbWafCustomRuleGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID associated with the ALB WAF Custom Rule Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/alb_waf_custom_rule_group#project_id AlbWafCustomRuleGroup#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Enriched rules containing auto-generated IDs and computed severity values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/alb_waf_custom_rule_group#rules AlbWafCustomRuleGroup#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// STACKIT region name the resource is located in. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/alb_waf_custom_rule_group#region AlbWafCustomRuleGroup#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

