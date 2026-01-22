package logsinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsInstanceConfig struct {
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
	// The displayed name of the Logs instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/logs_instance#display_name LogsInstance#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// STACKIT project ID associated with the Logs instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/logs_instance#project_id LogsInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The log retention time in days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/logs_instance#retention_days LogsInstance#retention_days}
	RetentionDays *float64 `field:"required" json:"retentionDays" yaml:"retentionDays"`
	// The access control list entries for the Logs instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/logs_instance#acl LogsInstance#acl}
	Acl *[]*string `field:"optional" json:"acl" yaml:"acl"`
	// The description of the Logs instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/logs_instance#description LogsInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// STACKIT region name the resource is located in. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/logs_instance#region LogsInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

