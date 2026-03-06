package applicationloadbalancer


type ApplicationLoadBalancerOptionsAccessControl struct {
	// Application Load Balancer is accessible only from an IP address in this range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.85.0/docs/resources/application_load_balancer#allowed_source_ranges ApplicationLoadBalancer#allowed_source_ranges}
	AllowedSourceRanges *[]*string `field:"required" json:"allowedSourceRanges" yaml:"allowedSourceRanges"`
}

