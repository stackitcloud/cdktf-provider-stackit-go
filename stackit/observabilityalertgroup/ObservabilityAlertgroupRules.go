package observabilityalertgroup


type ObservabilityAlertgroupRules struct {
	// The PromQL expression to evaluate.
	//
	// Every evaluation cycle this is evaluated at the current time, and all resultant time series become pending/firing alerts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/observability_alertgroup#expression ObservabilityAlertgroup#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The name of the alert rule. Is the identifier and must be unique in the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/observability_alertgroup#alert ObservabilityAlertgroup#alert}
	Alert *string `field:"optional" json:"alert" yaml:"alert"`
	// A map of key:value. Annotations to add or overwrite for each alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/observability_alertgroup#annotations ObservabilityAlertgroup#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Alerts are considered firing once they have been returned for this long.
	//
	// Alerts which have not yet fired for long enough are considered pending. Default is 0s
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/observability_alertgroup#for ObservabilityAlertgroup#for}
	For *string `field:"optional" json:"for" yaml:"for"`
	// A map of key:value. Labels to add or overwrite for each alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/observability_alertgroup#labels ObservabilityAlertgroup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the metric. It's the identifier and must be unique in the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/observability_alertgroup#record ObservabilityAlertgroup#record}
	Record *string `field:"optional" json:"record" yaml:"record"`
}

