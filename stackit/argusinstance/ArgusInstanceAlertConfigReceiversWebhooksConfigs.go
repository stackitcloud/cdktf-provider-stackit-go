package argusinstance


type ArgusInstanceAlertConfigReceiversWebhooksConfigs struct {
	// Microsoft Teams webhooks require special handling, set this to true if the webhook is for Microsoft Teams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.1/docs/resources/argus_instance#ms_teams ArgusInstance#ms_teams}
	MsTeams interface{} `field:"optional" json:"msTeams" yaml:"msTeams"`
	// The endpoint to send HTTP POST requests to. Must be a valid URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.43.1/docs/resources/argus_instance#url ArgusInstance#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

