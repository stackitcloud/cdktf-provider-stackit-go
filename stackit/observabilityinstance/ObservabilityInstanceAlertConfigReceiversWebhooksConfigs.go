package observabilityinstance


type ObservabilityInstanceAlertConfigReceiversWebhooksConfigs struct {
	// Microsoft Teams webhooks require special handling, set this to true if the webhook is for Microsoft Teams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.34.0/docs/resources/observability_instance#ms_teams ObservabilityInstance#ms_teams}
	MsTeams interface{} `field:"optional" json:"msTeams" yaml:"msTeams"`
	// The endpoint to send HTTP POST requests to. Must be a valid URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.34.0/docs/resources/observability_instance#url ObservabilityInstance#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

