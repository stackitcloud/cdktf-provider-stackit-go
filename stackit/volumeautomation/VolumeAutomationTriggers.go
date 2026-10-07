package volumeautomation


type VolumeAutomationTriggers struct {
	// Runs the automation on a recurring schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.118.0/docs/resources/volume_automation#schedule VolumeAutomation#schedule}
	Schedule *VolumeAutomationTriggersSchedule `field:"optional" json:"schedule" yaml:"schedule"`
}

