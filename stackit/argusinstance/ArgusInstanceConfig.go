package argusinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ArgusInstanceConfig struct {
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
	// The name of the Argus instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.26.2/docs/resources/argus_instance#name ArgusInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the Argus plan. E.g. `Monitoring-Medium-EU01`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.26.2/docs/resources/argus_instance#plan_name ArgusInstance#plan_name}
	PlanName *string `field:"required" json:"planName" yaml:"planName"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.26.2/docs/resources/argus_instance#project_id ArgusInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The access control list for this instance.
	//
	// Each entry is an IP address range that is permitted to access, in CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.26.2/docs/resources/argus_instance#acl ArgusInstance#acl}
	Acl *[]*string `field:"optional" json:"acl" yaml:"acl"`
	// Alert configuration for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.26.2/docs/resources/argus_instance#alert_config ArgusInstance#alert_config}
	AlertConfig *ArgusInstanceAlertConfig `field:"optional" json:"alertConfig" yaml:"alertConfig"`
	// Specifies for how many days the raw metrics are kept.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.26.2/docs/resources/argus_instance#metrics_retention_days ArgusInstance#metrics_retention_days}
	MetricsRetentionDays *float64 `field:"optional" json:"metricsRetentionDays" yaml:"metricsRetentionDays"`
	// Specifies for how many days the 1h downsampled metrics are kept.
	//
	// must be less than the value of the 5m downsampling retention. Default is set to `0` (disabled).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.26.2/docs/resources/argus_instance#metrics_retention_days_1h_downsampling ArgusInstance#metrics_retention_days_1h_downsampling}
	MetricsRetentionDays1HDownsampling *float64 `field:"optional" json:"metricsRetentionDays1HDownsampling" yaml:"metricsRetentionDays1HDownsampling"`
	// Specifies for how many days the 5m downsampled metrics are kept.
	//
	// must be less than the value of the general retention. Default is set to `0` (disabled).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.26.2/docs/resources/argus_instance#metrics_retention_days_5m_downsampling ArgusInstance#metrics_retention_days_5m_downsampling}
	MetricsRetentionDays5MDownsampling *float64 `field:"optional" json:"metricsRetentionDays5MDownsampling" yaml:"metricsRetentionDays5MDownsampling"`
	// Additional parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.26.2/docs/resources/argus_instance#parameters ArgusInstance#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

