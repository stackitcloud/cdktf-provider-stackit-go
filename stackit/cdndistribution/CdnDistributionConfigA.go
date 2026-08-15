package cdndistribution


type CdnDistributionConfigA struct {
	// The configured backend for the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/cdn_distribution#backend CdnDistribution#backend}
	Backend *CdnDistributionConfigBackend `field:"required" json:"backend" yaml:"backend"`
	// The configured regions where content will be hosted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/cdn_distribution#regions CdnDistribution#regions}
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
	// The configured countries where distribution of content is blocked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/cdn_distribution#blocked_countries CdnDistribution#blocked_countries}
	BlockedCountries *[]*string `field:"optional" json:"blockedCountries" yaml:"blockedCountries"`
	// Restricts access to your content by specifying a list of blocked IPv4 addresses.
	//
	// This feature enhances security and privacy by preventing these addresses from accessing your distribution. Note: once a value is set, removing the attribute from your configuration will retain the last known value in state; to clear it explicitly, set it to an empty list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/cdn_distribution#blocked_ips CdnDistribution#blocked_ips}
	BlockedIps *[]*string `field:"optional" json:"blockedIps" yaml:"blockedIps"`
	// Sets the default cache duration for the distribution.
	//
	// The default cache duration is applied when a 'Cache-Control' header is not presented in the origin's response. We use ISO8601 duration format for cache duration (e.g. P1DT2H30M). Note: once a value is set, removing the attribute from your configuration will retain the last known value in state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/cdn_distribution#default_cache_duration CdnDistribution#default_cache_duration}
	DefaultCacheDuration *string `field:"optional" json:"defaultCacheDuration" yaml:"defaultCacheDuration"`
	// Enable this allows the 'Host' header to be passed through to the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/cdn_distribution#forward_host_header CdnDistribution#forward_host_header}
	ForwardHostHeader interface{} `field:"optional" json:"forwardHostHeader" yaml:"forwardHostHeader"`
	// Sets the monthly limit of bandwidth in bytes that the pullzone is allowed to use.
	//
	// Note: once a value is set, removing the attribute from your configuration will retain the last known value in state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/cdn_distribution#monthly_limit_bytes CdnDistribution#monthly_limit_bytes}
	MonthlyLimitBytes *float64 `field:"optional" json:"monthlyLimitBytes" yaml:"monthlyLimitBytes"`
	// Configuration for the Image Optimizer.
	//
	// This is a paid feature that automatically optimizes images to reduce their file size for faster delivery, leading to improved website performance and a better user experience.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/cdn_distribution#optimizer CdnDistribution#optimizer}
	Optimizer *CdnDistributionConfigOptimizer `field:"optional" json:"optimizer" yaml:"optimizer"`
	// A wrapper for a list of redirect rules that allows for redirect settings on a distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/cdn_distribution#redirects CdnDistribution#redirects}
	Redirects *CdnDistributionConfigRedirects `field:"optional" json:"redirects" yaml:"redirects"`
	// Enable this to prevent origin-level cookies from being forwarded to the end user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/cdn_distribution#strip_response_cookies CdnDistribution#strip_response_cookies}
	StripResponseCookies interface{} `field:"optional" json:"stripResponseCookies" yaml:"stripResponseCookies"`
	// Configuration for TLS protocol versions. Note: Enabling older TLS versions (1.0, 1.1) is generally discouraged for security reasons.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/cdn_distribution#tls CdnDistribution#tls}
	Tls *CdnDistributionConfigTls `field:"optional" json:"tls" yaml:"tls"`
	// Configures the Web Application Firewall (WAF) for the distribution.
	//
	// If this block is undefined or removed from your configuration, the WAF mode will default to DISABLED and the type to FREE. All other WAF properties will retain their last known state in the API; if they were never defined, the API will apply its default settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.111.0/docs/resources/cdn_distribution#waf CdnDistribution#waf}
	Waf *CdnDistributionConfigWaf `field:"optional" json:"waf" yaml:"waf"`
}

