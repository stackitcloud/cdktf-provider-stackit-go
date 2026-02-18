package dnszone

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsZoneConfig struct {
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
	// The zone name. E.g. `example.com`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/dns_zone#dns_name DnsZone#dns_name}
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// The user given name of the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/dns_zone#name DnsZone#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the dns zone is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/dns_zone#project_id DnsZone#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The access control list. E.g. `0.0.0.0/0,::/0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/dns_zone#acl DnsZone#acl}
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/dns_zone#active DnsZone#active}.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// A contact e-mail for the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/dns_zone#contact_email DnsZone#contact_email}
	ContactEmail *string `field:"optional" json:"contactEmail" yaml:"contactEmail"`
	// Default time to live. E.g. 3600.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/dns_zone#default_ttl DnsZone#default_ttl}
	DefaultTtl *float64 `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// Description of the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/dns_zone#description DnsZone#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Expire time. E.g. 1209600.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/dns_zone#expire_time DnsZone#expire_time}
	ExpireTime *float64 `field:"optional" json:"expireTime" yaml:"expireTime"`
	// Specifies, if the zone is a reverse zone or not. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/dns_zone#is_reverse_zone DnsZone#is_reverse_zone}
	IsReverseZone interface{} `field:"optional" json:"isReverseZone" yaml:"isReverseZone"`
	// Negative caching. E.g. 60.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/dns_zone#negative_cache DnsZone#negative_cache}
	NegativeCache *float64 `field:"optional" json:"negativeCache" yaml:"negativeCache"`
	// Primary name server for secondary zone. E.g. ["1.2.3.4"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/dns_zone#primaries DnsZone#primaries}
	Primaries *[]*string `field:"optional" json:"primaries" yaml:"primaries"`
	// Refresh time. E.g. 3600.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/dns_zone#refresh_time DnsZone#refresh_time}
	RefreshTime *float64 `field:"optional" json:"refreshTime" yaml:"refreshTime"`
	// Retry time. E.g. 600.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/dns_zone#retry_time DnsZone#retry_time}
	RetryTime *float64 `field:"optional" json:"retryTime" yaml:"retryTime"`
	// Zone type. Defaults to `primary`. Possible values are: `primary`, `secondary`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/dns_zone#type DnsZone#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

