package loadbalancer


type LoadbalancerTargetPoolsTargets struct {
	// Target display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.6.3/docs/resources/loadbalancer#display_name Loadbalancer#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Target IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.6.3/docs/resources/loadbalancer#ip Loadbalancer#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
}

