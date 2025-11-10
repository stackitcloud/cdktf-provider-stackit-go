package datastackitcdncustomdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitCdnCustomDomainConfig struct {
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
	// CDN distribution ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.56.0/docs/data-sources/cdn_custom_domain#distribution_id DataStackitCdnCustomDomain#distribution_id}
	DistributionId *string `field:"required" json:"distributionId" yaml:"distributionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.56.0/docs/data-sources/cdn_custom_domain#name DataStackitCdnCustomDomain#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID associated with the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.56.0/docs/data-sources/cdn_custom_domain#project_id DataStackitCdnCustomDomain#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

