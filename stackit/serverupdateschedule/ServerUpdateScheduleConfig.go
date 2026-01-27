package serverupdateschedule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServerUpdateScheduleConfig struct {
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
	// Is the update schedule enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.1/docs/resources/server_update_schedule#enabled ServerUpdateSchedule#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Maintenance window [1..24].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.1/docs/resources/server_update_schedule#maintenance_window ServerUpdateSchedule#maintenance_window}
	MaintenanceWindow *float64 `field:"required" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// The schedule name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.1/docs/resources/server_update_schedule#name ServerUpdateSchedule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT Project ID to which the server is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.1/docs/resources/server_update_schedule#project_id ServerUpdateSchedule#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Update schedule described in `rrule` (recurrence rule) format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.1/docs/resources/server_update_schedule#rrule ServerUpdateSchedule#rrule}
	Rrule *string `field:"required" json:"rrule" yaml:"rrule"`
	// Server ID for the update schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.1/docs/resources/server_update_schedule#server_id ServerUpdateSchedule#server_id}
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.1/docs/resources/server_update_schedule#region ServerUpdateSchedule#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

