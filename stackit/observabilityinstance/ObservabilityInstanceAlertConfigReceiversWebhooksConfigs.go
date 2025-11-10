package observabilityinstance


type ObservabilityInstanceAlertConfigReceiversWebhooksConfigs struct {
	// Google Chat webhooks require special handling, set this to true if the webhook is for Google Chat.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.62.0/docs/resources/observability_instance#google_chat ObservabilityInstance#google_chat}
	GoogleChat interface{} `field:"optional" json:"googleChat" yaml:"googleChat"`
	// Microsoft Teams webhooks require special handling, set this to true if the webhook is for Microsoft Teams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.62.0/docs/resources/observability_instance#ms_teams ObservabilityInstance#ms_teams}
	MsTeams interface{} `field:"optional" json:"msTeams" yaml:"msTeams"`
	// The endpoint to send HTTP POST requests to. Must be a valid URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.62.0/docs/resources/observability_instance#url ObservabilityInstance#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

