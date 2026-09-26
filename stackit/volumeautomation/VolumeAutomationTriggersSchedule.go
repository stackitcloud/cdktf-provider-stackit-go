package volumeautomation


type VolumeAutomationTriggersSchedule struct {
	// An `rrule` (Recurrence Rule) is a standardized string format used in iCalendar (RFC 5545) to define repeating events, and you can generate one by using a dedicated library or by using online generator tools to specify parameters like frequency, interval, and end dates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.117.0/docs/resources/volume_automation#rrule VolumeAutomation#rrule}
	Rrule *string `field:"required" json:"rrule" yaml:"rrule"`
}

