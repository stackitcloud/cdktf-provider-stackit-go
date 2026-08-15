package datastackitalbwafcustomrulegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitAlbWafCustomRuleGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/data-sources/alb_waf_custom_rule_group#name DataStackitAlbWafCustomRuleGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID associated with the ALB WAF Custom Rule Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/data-sources/alb_waf_custom_rule_group#project_id DataStackitAlbWafCustomRuleGroup#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// STACKIT region name the resource is located in. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/data-sources/alb_waf_custom_rule_group#region DataStackitAlbWafCustomRuleGroup#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

