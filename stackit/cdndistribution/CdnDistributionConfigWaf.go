package cdndistribution


type CdnDistributionConfigWaf struct {
	// Restricts which HTTP methods the distribution accepts.
	//
	// If provided, the set must contain at least one item. Case you removed waf will retain the last known state and if omitted, the API applies the following defaults: `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#allowed_http_methods CdnDistribution#allowed_http_methods}
	AllowedHttpMethods *[]*string `field:"optional" json:"allowedHttpMethods" yaml:"allowedHttpMethods"`
	// Restricts which HTTP protocol versions are accepted.
	//
	// If provided, the set must contain at least one item. If omitted, the API applies the following defaults: `HTTP/1.0`, `HTTP/1.1`, `HTTP/2`, `HTTP/2.0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#allowed_http_versions CdnDistribution#allowed_http_versions}
	AllowedHttpVersions *[]*string `field:"optional" json:"allowedHttpVersions" yaml:"allowedHttpVersions"`
	// Restricts which Content-Type headers are accepted in request bodies.
	//
	// If provided, the set must contain at least one item. Case you removed waf will retain the last known state and if omitted, the API applies the following defaults: `application/x-www-form-urlencoded`, `multipart/form-data`, `multipart/related`, `text/xml`, `application/xml`, `application/soap+xml`, `application/x-amf`, `application/json`, `application/octet-stream`, `application/csp-report`, `application/xss-auditor-report`, `text/plain`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#allowed_request_content_types CdnDistribution#allowed_request_content_types}
	AllowedRequestContentTypes *[]*string `field:"optional" json:"allowedRequestContentTypes" yaml:"allowedRequestContentTypes"`
	// Set of WAF Collection IDs explicitly disabled.
	//
	// Can be set to an empty set to clear previously set rules. Case you removed waf will retain the last known state. To view available rule collections, please consult the API documentation: https://docs.api.eu01.stackit.cloud/documentation/cdn/version/v1#tag/WAF/operation/ListWafCollections
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#disabled_rule_collection_ids CdnDistribution#disabled_rule_collection_ids}
	DisabledRuleCollectionIds *[]*string `field:"optional" json:"disabledRuleCollectionIds" yaml:"disabledRuleCollectionIds"`
	// Set of WAF Rule Group IDs explicitly disabled.
	//
	// Can be set to an empty set to clear previously set rules. Case you removed waf will retain the last known state. Precedence hierarchy: Groups override Collections. To view available rule groups, please consult the API documentation: https://docs.api.eu01.stackit.cloud/documentation/cdn/version/v1#tag/WAF/operation/ListWafCollections
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#disabled_rule_group_ids CdnDistribution#disabled_rule_group_ids}
	DisabledRuleGroupIds *[]*string `field:"optional" json:"disabledRuleGroupIds" yaml:"disabledRuleGroupIds"`
	// Set of WAF rule IDs explicitly disabled.
	//
	// Can be set to an empty set to clear previously set rules. Case you removed waf will retain the last known state. Precedence hierarchy: Specific Rules override Groups. For example, an explicitly disabled Rule ID takes precedence over an enabled Group ID. To view available rules, please consult the API documentation: https://docs.api.eu01.stackit.cloud/documentation/cdn/version/v1#tag/WAF/operation/ListWafCollections
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#disabled_rule_ids CdnDistribution#disabled_rule_ids}
	DisabledRuleIds *[]*string `field:"optional" json:"disabledRuleIds" yaml:"disabledRuleIds"`
	// Set of WAF Collection IDs explicitly enabled.
	//
	// Can be set to an empty set to clear previously set rules. Case you removed waf will retain the last known state. To view available rule collections, please consult the API documentation: https://docs.api.eu01.stackit.cloud/documentation/cdn/version/v1#tag/WAF/operation/ListWafCollections
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#enabled_rule_collection_ids CdnDistribution#enabled_rule_collection_ids}
	EnabledRuleCollectionIds *[]*string `field:"optional" json:"enabledRuleCollectionIds" yaml:"enabledRuleCollectionIds"`
	// Set of WAF Rule Group IDs explicitly enabled.
	//
	// Can be set to an empty set to clear previously set rules. Case you removed waf will retain the last known state. Precedence hierarchy: Groups override Collections. To view available rule groups, please consult the API documentation: https://docs.api.eu01.stackit.cloud/documentation/cdn/version/v1#tag/WAF/operation/ListWafCollections
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#enabled_rule_group_ids CdnDistribution#enabled_rule_group_ids}
	EnabledRuleGroupIds *[]*string `field:"optional" json:"enabledRuleGroupIds" yaml:"enabledRuleGroupIds"`
	// Set of WAF rule IDs explicitly enabled.
	//
	// Can be set to an empty set to clear previously set rules. Case you removed waf will retain the last known state. Precedence hierarchy: Specific Rules override Groups. For example, an explicitly enabled Rule ID takes precedence over a disabled Group ID. To view available rules, please consult the API documentation: https://docs.api.eu01.stackit.cloud/documentation/cdn/version/v1#tag/WAF/operation/ListWafCollections
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#enabled_rule_ids CdnDistribution#enabled_rule_ids}
	EnabledRuleIds *[]*string `field:"optional" json:"enabledRuleIds" yaml:"enabledRuleIds"`
	// Set of WAF Collection IDs explicitly marked as Log Only.
	//
	// Can be set to an empty set to clear previously set rules. Case you removed waf will retain the last known state. To view available rule collections, please consult the API documentation: https://docs.api.eu01.stackit.cloud/documentation/cdn/version/v1#tag/WAF/operation/ListWafCollections
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#log_only_rule_collection_ids CdnDistribution#log_only_rule_collection_ids}
	LogOnlyRuleCollectionIds *[]*string `field:"optional" json:"logOnlyRuleCollectionIds" yaml:"logOnlyRuleCollectionIds"`
	// Set of WAF Rule Group IDs explicitly marked as Log Only.
	//
	// Can be set to an empty set to clear previously set rules. Case you removed waf will retain the last known state. Precedence hierarchy: Groups override Collections. To view available rule groups, please consult the API documentation: https://docs.api.eu01.stackit.cloud/documentation/cdn/version/v1#tag/WAF/operation/ListWafCollections
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#log_only_rule_group_ids CdnDistribution#log_only_rule_group_ids}
	LogOnlyRuleGroupIds *[]*string `field:"optional" json:"logOnlyRuleGroupIds" yaml:"logOnlyRuleGroupIds"`
	// Set of WAF rule IDs explicitly marked as Log Only.
	//
	// Can be set to an empty set to clear previously set rules. Case you removed waf will retain the last known state. Precedence hierarchy: Specific Rules override Groups. To view available rules, please consult the API documentation: https://docs.api.eu01.stackit.cloud/documentation/cdn/version/v1#tag/WAF/operation/ListWafCollections
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#log_only_rule_ids CdnDistribution#log_only_rule_ids}
	LogOnlyRuleIds *[]*string `field:"optional" json:"logOnlyRuleIds" yaml:"logOnlyRuleIds"`
	// The operating mode of the WAF.
	//
	// 'ENABLED' actively blocks threats, 'LOG_ONLY' logs matches without blocking, and 'DISABLED' completely turns off inspection. Defaults to 'DISABLED'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#mode CdnDistribution#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Defines how aggressively the WAF should act on requests.
	//
	// Valid values are 'L1' to 'L4'. Case you removed waf will retain the last known state and if omitted, The API applies the following default 'L1'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#paranoia_level CdnDistribution#paranoia_level}
	ParanoiaLevel *string `field:"optional" json:"paranoiaLevel" yaml:"paranoiaLevel"`
	// The tier of the WAF. Valid values are 'FREE' or 'PREMIUM'. Defaults to 'FREE'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.97.0/docs/resources/cdn_distribution#type CdnDistribution#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

