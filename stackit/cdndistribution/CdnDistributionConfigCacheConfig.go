package cdndistribution


type CdnDistributionConfigCacheConfig struct {
	// HTTP request header names whose values participate in the cache key.
	//
	// Each entry multiplies the number of cache variants by the number of distinct values observed for that header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.123.0/docs/resources/cdn_distribution#cache_key_headers CdnDistribution#cache_key_headers}
	CacheKeyHeaders *[]*string `field:"optional" json:"cacheKeyHeaders" yaml:"cacheKeyHeaders"`
	// When true, the CDN varies its cache by the request query string.
	//
	// If query_string_vary_parameters is empty, every unique query-string combination produces its own cache entry; if non-empty, only the listed parameters influence the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.123.0/docs/resources/cdn_distribution#query_string_vary_enabled CdnDistribution#query_string_vary_enabled}
	QueryStringVaryEnabled interface{} `field:"optional" json:"queryStringVaryEnabled" yaml:"queryStringVaryEnabled"`
	// Allowlist of query-string parameter names that participate in the cache key when query_string_vary_enabled is true.
	//
	// Ignored while query_string_vary_enabled is false, but still stored so it can be re-activated without losing the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.123.0/docs/resources/cdn_distribution#query_string_vary_parameters CdnDistribution#query_string_vary_parameters}
	QueryStringVaryParameters *[]*string `field:"optional" json:"queryStringVaryParameters" yaml:"queryStringVaryParameters"`
}

