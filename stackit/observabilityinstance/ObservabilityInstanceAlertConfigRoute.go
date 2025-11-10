package observabilityinstance


type ObservabilityInstanceAlertConfigRoute struct {
	// The name of the receiver to route the alerts to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.39.1/docs/resources/observability_instance#receiver ObservabilityInstance#receiver}
	Receiver *string `field:"required" json:"receiver" yaml:"receiver"`
	// The labels by which incoming alerts are grouped together.
	//
	// For example, multiple alerts coming in for cluster=A and alertname=LatencyHigh would be batched into a single group. To aggregate by all possible labels use the special value '...' as the sole label name, for example: group_by: ['...']. This effectively disables aggregation entirely, passing through all alerts as-is. This is unlikely to be what you want, unless you have a very low alert volume or your upstream notification system performs its own grouping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.39.1/docs/resources/observability_instance#group_by ObservabilityInstance#group_by}
	GroupBy *[]*string `field:"optional" json:"groupBy" yaml:"groupBy"`
	// How long to wait before sending a notification about new alerts that are added to a group of alerts for which an initial notification has already been sent.
	//
	// (Usually ~5m or more.)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.39.1/docs/resources/observability_instance#group_interval ObservabilityInstance#group_interval}
	GroupInterval *string `field:"optional" json:"groupInterval" yaml:"groupInterval"`
	// How long to initially wait to send a notification for a group of alerts.
	//
	// Allows to wait for an inhibiting alert to arrive or collect more initial alerts for the same group. (Usually ~0s to few minutes.)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.39.1/docs/resources/observability_instance#group_wait ObservabilityInstance#group_wait}
	GroupWait *string `field:"optional" json:"groupWait" yaml:"groupWait"`
	// A set of equality matchers an alert has to fulfill to match the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.39.1/docs/resources/observability_instance#match ObservabilityInstance#match}
	Match *map[string]*string `field:"optional" json:"match" yaml:"match"`
	// A set of regex-matchers an alert has to fulfill to match the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.39.1/docs/resources/observability_instance#match_regex ObservabilityInstance#match_regex}
	MatchRegex *map[string]*string `field:"optional" json:"matchRegex" yaml:"matchRegex"`
	// How long to wait before sending a notification again if it has already been sent successfully for an alert.
	//
	// (Usually ~3h or more).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.39.1/docs/resources/observability_instance#repeat_interval ObservabilityInstance#repeat_interval}
	RepeatInterval *string `field:"optional" json:"repeatInterval" yaml:"repeatInterval"`
	// List of child routes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.39.1/docs/resources/observability_instance#routes ObservabilityInstance#routes}
	Routes interface{} `field:"optional" json:"routes" yaml:"routes"`
}

