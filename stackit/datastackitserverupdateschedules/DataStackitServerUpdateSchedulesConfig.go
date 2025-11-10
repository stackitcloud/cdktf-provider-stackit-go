package datastackitserverupdateschedules

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitServerUpdateSchedulesConfig struct {
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
	// STACKIT Project ID (UUID) to which the server is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/data-sources/server_update_schedules#project_id DataStackitServerUpdateSchedules#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Server ID (UUID) to which the update schedule is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/data-sources/server_update_schedules#server_id DataStackitServerUpdateSchedules#server_id}
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/data-sources/server_update_schedules#region DataStackitServerUpdateSchedules#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

