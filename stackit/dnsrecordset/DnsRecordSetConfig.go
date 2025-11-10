package dnsrecordset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsRecordSetConfig struct {
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
	// Name of the record which should be a valid domain according to rfc1035 Section 2.3.4. E.g. `example.com`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.21.1/docs/resources/dns_record_set#name DnsRecordSet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the dns record set is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.21.1/docs/resources/dns_record_set#project_id DnsRecordSet#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Records.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.21.1/docs/resources/dns_record_set#records DnsRecordSet#records}
	Records *[]*string `field:"required" json:"records" yaml:"records"`
	// The zone ID to which is dns record set is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.21.1/docs/resources/dns_record_set#zone_id DnsRecordSet#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Specifies if the record set is active or not. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.21.1/docs/resources/dns_record_set#active DnsRecordSet#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// Comment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.21.1/docs/resources/dns_record_set#comment DnsRecordSet#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Time to live. E.g. 3600.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.21.1/docs/resources/dns_record_set#ttl DnsRecordSet#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
	// The record set type. E.g. `A` or `CNAME`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.21.1/docs/resources/dns_record_set#type DnsRecordSet#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

