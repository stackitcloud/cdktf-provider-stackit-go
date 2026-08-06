package albwafconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbWafConfigurationConfig struct {
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
	// The name of the WAF Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.110.0/docs/resources/alb_waf_configuration#name AlbWafConfiguration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the WAF Configuration is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.110.0/docs/resources/alb_waf_configuration#project_id AlbWafConfiguration#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Name of the custom rule group for this WAF Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.110.0/docs/resources/alb_waf_configuration#custom_rule_group_name AlbWafConfiguration#custom_rule_group_name}
	CustomRuleGroupName *string `field:"optional" json:"customRuleGroupName" yaml:"customRuleGroupName"`
	// User-defined metadata as key-value pairs. Should not exceed 64 entries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.110.0/docs/resources/alb_waf_configuration#labels AlbWafConfiguration#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Name of the managed rule set configuration for this WAF Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.110.0/docs/resources/alb_waf_configuration#managed_rule_set_name AlbWafConfiguration#managed_rule_set_name}
	ManagedRuleSetName *string `field:"optional" json:"managedRuleSetName" yaml:"managedRuleSetName"`
	// The resource region (e.g. eu01). If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.110.0/docs/resources/alb_waf_configuration#region AlbWafConfiguration#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

