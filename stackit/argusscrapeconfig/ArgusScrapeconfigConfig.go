package argusscrapeconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ArgusScrapeconfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.15.2/docs/resources/argus_scrapeconfig#instance_id ArgusScrapeconfig#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Specifies the job scraping url path. E.g. `/metrics`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.15.2/docs/resources/argus_scrapeconfig#metrics_path ArgusScrapeconfig#metrics_path}
	MetricsPath *string `field:"required" json:"metricsPath" yaml:"metricsPath"`
	// Specifies the name of the scraping job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.15.2/docs/resources/argus_scrapeconfig#name ArgusScrapeconfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the scraping job is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.15.2/docs/resources/argus_scrapeconfig#project_id ArgusScrapeconfig#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The targets list (specified by the static config).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.15.2/docs/resources/argus_scrapeconfig#targets ArgusScrapeconfig#targets}
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
	// A basic authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.15.2/docs/resources/argus_scrapeconfig#basic_auth ArgusScrapeconfig#basic_auth}
	BasicAuth *ArgusScrapeconfigBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	// A SAML2 configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.15.2/docs/resources/argus_scrapeconfig#saml2 ArgusScrapeconfig#saml2}
	Saml2 *ArgusScrapeconfigSaml2 `field:"optional" json:"saml2" yaml:"saml2"`
	// Specifies the scrape sample limit. Upper limit depends on the service plan. Default is `5000`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.15.2/docs/resources/argus_scrapeconfig#sample_limit ArgusScrapeconfig#sample_limit}
	SampleLimit *float64 `field:"optional" json:"sampleLimit" yaml:"sampleLimit"`
	// Specifies the http scheme. E.g. `https`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.15.2/docs/resources/argus_scrapeconfig#scheme ArgusScrapeconfig#scheme}
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// Specifies the scrape interval as duration string. E.g. `5m`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.15.2/docs/resources/argus_scrapeconfig#scrape_interval ArgusScrapeconfig#scrape_interval}
	ScrapeInterval *string `field:"optional" json:"scrapeInterval" yaml:"scrapeInterval"`
	// Specifies the scrape timeout as duration string. E.g.`2m`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.15.2/docs/resources/argus_scrapeconfig#scrape_timeout ArgusScrapeconfig#scrape_timeout}
	ScrapeTimeout *string `field:"optional" json:"scrapeTimeout" yaml:"scrapeTimeout"`
}

