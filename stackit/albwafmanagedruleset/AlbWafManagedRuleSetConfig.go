package albwafmanagedruleset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbWafManagedRuleSetConfig struct {
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
	// Managed Rule Set configuration name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.113.0/docs/resources/alb_waf_managed_rule_set#name AlbWafManagedRuleSet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID associated with the ALB WAF Managed Rule Set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.113.0/docs/resources/alb_waf_managed_rule_set#project_id AlbWafManagedRuleSet#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Type of the Managed Rule Set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.113.0/docs/resources/alb_waf_managed_rule_set#type AlbWafManagedRuleSet#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// STACKIT region name the resource is located in. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.113.0/docs/resources/alb_waf_managed_rule_set#region AlbWafManagedRuleSet#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

