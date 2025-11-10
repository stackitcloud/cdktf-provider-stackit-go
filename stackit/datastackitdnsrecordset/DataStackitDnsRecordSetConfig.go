package datastackitdnsrecordset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitDnsRecordSetConfig struct {
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
	// STACKIT project ID to which the dns record set is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/data-sources/dns_record_set#project_id DataStackitDnsRecordSet#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The rr set id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/data-sources/dns_record_set#record_set_id DataStackitDnsRecordSet#record_set_id}
	RecordSetId *string `field:"required" json:"recordSetId" yaml:"recordSetId"`
	// The zone ID to which is dns record set is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.60.0/docs/data-sources/dns_record_set#zone_id DataStackitDnsRecordSet#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

