package cdncustomdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnCustomDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.80.0/docs/resources/cdn_custom_domain#distribution_id CdnCustomDomain#distribution_id}
	DistributionId *string `field:"required" json:"distributionId" yaml:"distributionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.80.0/docs/resources/cdn_custom_domain#name CdnCustomDomain#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID associated with the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.80.0/docs/resources/cdn_custom_domain#project_id CdnCustomDomain#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The TLS certificate for the custom domain.
	//
	// If omitted, a managed certificate will be used. If the block is specified, a custom certificate is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.80.0/docs/resources/cdn_custom_domain#certificate CdnCustomDomain#certificate}
	Certificate *CdnCustomDomainCertificate `field:"optional" json:"certificate" yaml:"certificate"`
}

