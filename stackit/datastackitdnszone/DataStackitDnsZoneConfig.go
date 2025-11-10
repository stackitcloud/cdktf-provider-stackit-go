package datastackitdnszone

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitDnsZoneConfig struct {
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
	// STACKIT project ID to which the dns zone is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/data-sources/dns_zone#project_id DataStackitDnsZone#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The zone ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/data-sources/dns_zone#zone_id DataStackitDnsZone#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

