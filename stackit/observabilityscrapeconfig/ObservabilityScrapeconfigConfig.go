package observabilityscrapeconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObservabilityScrapeconfigConfig struct {
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
	// Observability instance ID to which the scraping job is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/observability_scrapeconfig#instance_id ObservabilityScrapeconfig#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Specifies the job scraping url path. E.g. `/metrics`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/observability_scrapeconfig#metrics_path ObservabilityScrapeconfig#metrics_path}
	MetricsPath *string `field:"required" json:"metricsPath" yaml:"metricsPath"`
	// Specifies the name of the scraping job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/observability_scrapeconfig#name ObservabilityScrapeconfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the scraping job is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/observability_scrapeconfig#project_id ObservabilityScrapeconfig#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The targets list (specified by the static config).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/observability_scrapeconfig#targets ObservabilityScrapeconfig#targets}
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
	// A basic authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/observability_scrapeconfig#basic_auth ObservabilityScrapeconfig#basic_auth}
	BasicAuth *ObservabilityScrapeconfigBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	// A SAML2 configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/observability_scrapeconfig#saml2 ObservabilityScrapeconfig#saml2}
	Saml2 *ObservabilityScrapeconfigSaml2 `field:"optional" json:"saml2" yaml:"saml2"`
	// Specifies the scrape sample limit. Upper limit depends on the service plan. Defaults to `5000`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/observability_scrapeconfig#sample_limit ObservabilityScrapeconfig#sample_limit}
	SampleLimit *float64 `field:"optional" json:"sampleLimit" yaml:"sampleLimit"`
	// Specifies the http scheme. Defaults to `https`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/observability_scrapeconfig#scheme ObservabilityScrapeconfig#scheme}
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// Specifies the scrape interval as duration string. Defaults to `5m`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/observability_scrapeconfig#scrape_interval ObservabilityScrapeconfig#scrape_interval}
	ScrapeInterval *string `field:"optional" json:"scrapeInterval" yaml:"scrapeInterval"`
	// Specifies the scrape timeout as duration string. Defaults to `2m`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/observability_scrapeconfig#scrape_timeout ObservabilityScrapeconfig#scrape_timeout}
	ScrapeTimeout *string `field:"optional" json:"scrapeTimeout" yaml:"scrapeTimeout"`
}

