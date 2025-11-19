package datastackitimagev2


type DataStackitImageV2Filter struct {
	// Filter images by operating system distribution. For example: `ubuntu`, `ubuntu-arm64`, `debian`, `rhel`, etc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.71.0/docs/data-sources/image_v2#distro DataStackitImageV2#distro}
	Distro *string `field:"optional" json:"distro" yaml:"distro"`
	// Filter images by operating system type, such as `linux` or `windows`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.71.0/docs/data-sources/image_v2#os DataStackitImageV2#os}
	Os *string `field:"optional" json:"os" yaml:"os"`
	// Filter images with Secure Boot support. Set to `true` to match images that support Secure Boot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.71.0/docs/data-sources/image_v2#secure_boot DataStackitImageV2#secure_boot}
	SecureBoot interface{} `field:"optional" json:"secureBoot" yaml:"secureBoot"`
	// Filter images based on UEFI support. Set to `true` to match images that support UEFI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.71.0/docs/data-sources/image_v2#uefi DataStackitImageV2#uefi}
	Uefi interface{} `field:"optional" json:"uefi" yaml:"uefi"`
	// Filter images by OS distribution version, such as `22.04`, `11`, or `9.1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.71.0/docs/data-sources/image_v2#version DataStackitImageV2#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

