package datastackitvpcroutingtablestaticroute


type DataStackitVpcRoutingTableStaticRouteTimeouts struct {
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/data-sources/vpc_routing_table_static_route#read DataStackitVpcRoutingTableStaticRoute#read}
	Read *string `field:"optional" json:"read" yaml:"read"`
}

