package datastackitalbwafconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitAlbWafConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/data-sources/alb_waf_configuration#name DataStackitAlbWafConfiguration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the WAF Configuration is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/data-sources/alb_waf_configuration#project_id DataStackitAlbWafConfiguration#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The resource region (e.g. eu01). If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/data-sources/alb_waf_configuration#region DataStackitAlbWafConfiguration#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

