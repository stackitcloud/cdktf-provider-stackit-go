package server


type ServerAgent struct {
	// Agent provisioning policy: `ALWAYS`, `NEVER`, or `INHERIT`. `INHERIT` follows the image default value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.96.0/docs/resources/server#provisioning_policy Server#provisioning_policy}
	ProvisioningPolicy *string `field:"optional" json:"provisioningPolicy" yaml:"provisioningPolicy"`
}

