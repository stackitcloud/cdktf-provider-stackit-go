package datastackitdnszone


type DataStackitDnsZoneTimeouts struct {
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.91.0/docs/data-sources/dns_zone#read DataStackitDnsZone#read}
	Read *string `field:"optional" json:"read" yaml:"read"`
}

