package datastackitserverupdateschedule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitServerUpdateScheduleConfig struct {
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
	// STACKIT Project ID to which the server is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.36.0/docs/data-sources/server_update_schedule#project_id DataStackitServerUpdateSchedule#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Server ID for the update schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.36.0/docs/data-sources/server_update_schedule#server_id DataStackitServerUpdateSchedule#server_id}
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// Update schedule ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.36.0/docs/data-sources/server_update_schedule#update_schedule_id DataStackitServerUpdateSchedule#update_schedule_id}
	UpdateScheduleId *float64 `field:"required" json:"updateScheduleId" yaml:"updateScheduleId"`
}

