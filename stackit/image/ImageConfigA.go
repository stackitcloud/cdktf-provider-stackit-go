package image


type ImageConfigA struct {
	// Enables the BIOS bootmenu.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/image#boot_menu Image#boot_menu}
	BootMenu interface{} `field:"optional" json:"bootMenu" yaml:"bootMenu"`
	// Sets CDROM bus controller type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/image#cdrom_bus Image#cdrom_bus}
	CdromBus *string `field:"optional" json:"cdromBus" yaml:"cdromBus"`
	// Sets Disk bus controller type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/image#disk_bus Image#disk_bus}
	DiskBus *string `field:"optional" json:"diskBus" yaml:"diskBus"`
	// Sets virtual network interface model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/image#nic_model Image#nic_model}
	NicModel *string `field:"optional" json:"nicModel" yaml:"nicModel"`
	// Enables operating system specific optimizations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/image#operating_system Image#operating_system}
	OperatingSystem *string `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// Operating system distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/image#operating_system_distro Image#operating_system_distro}
	OperatingSystemDistro *string `field:"optional" json:"operatingSystemDistro" yaml:"operatingSystemDistro"`
	// Version of the operating system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/image#operating_system_version Image#operating_system_version}
	OperatingSystemVersion *string `field:"optional" json:"operatingSystemVersion" yaml:"operatingSystemVersion"`
	// Sets the device bus when the image is used as a rescue image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/image#rescue_bus Image#rescue_bus}
	RescueBus *string `field:"optional" json:"rescueBus" yaml:"rescueBus"`
	// Sets the device when the image is used as a rescue image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/image#rescue_device Image#rescue_device}
	RescueDevice *string `field:"optional" json:"rescueDevice" yaml:"rescueDevice"`
	// Enables Secure Boot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/image#secure_boot Image#secure_boot}
	SecureBoot interface{} `field:"optional" json:"secureBoot" yaml:"secureBoot"`
	// Enables UEFI boot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/image#uefi Image#uefi}
	Uefi interface{} `field:"optional" json:"uefi" yaml:"uefi"`
	// Sets Graphic device model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/image#video_model Image#video_model}
	VideoModel *string `field:"optional" json:"videoModel" yaml:"videoModel"`
	// Enables the use of VirtIO SCSI to provide block device access. By default instances use VirtIO Block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/image#virtio_scsi Image#virtio_scsi}
	VirtioScsi interface{} `field:"optional" json:"virtioScsi" yaml:"virtioScsi"`
}

