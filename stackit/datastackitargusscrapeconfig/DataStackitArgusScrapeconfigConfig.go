package datastackitargusscrapeconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitArgusScrapeconfigConfig struct {
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
	// Argus instance ID to which the scraping job is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.26.3/docs/data-sources/argus_scrapeconfig#instance_id DataStackitArgusScrapeconfig#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Specifies the name of the scraping job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.26.3/docs/data-sources/argus_scrapeconfig#name DataStackitArgusScrapeconfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the scraping job is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.26.3/docs/data-sources/argus_scrapeconfig#project_id DataStackitArgusScrapeconfig#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

