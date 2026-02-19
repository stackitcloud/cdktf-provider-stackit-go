package cdndistribution


type CdnDistributionConfigBackend struct {
	// The configured backend type. Possible values are: `http`, `bucket`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/cdn_distribution#type CdnDistribution#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The URL of the bucket (e.g. https://s3.example.com). Required if type is 'bucket'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/cdn_distribution#bucket_url CdnDistribution#bucket_url}
	BucketUrl *string `field:"optional" json:"bucketUrl" yaml:"bucketUrl"`
	// The credentials for the bucket. Required if type is 'bucket'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/cdn_distribution#credentials CdnDistribution#credentials}
	Credentials *CdnDistributionConfigBackendCredentials `field:"optional" json:"credentials" yaml:"credentials"`
	// The configured type http to configure countries where content is allowed.
	//
	// A map of URLs to a list of countries
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/cdn_distribution#geofencing CdnDistribution#geofencing}
	Geofencing interface{} `field:"optional" json:"geofencing" yaml:"geofencing"`
	// The configured type http origin request headers for the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/cdn_distribution#origin_request_headers CdnDistribution#origin_request_headers}
	OriginRequestHeaders *map[string]*string `field:"optional" json:"originRequestHeaders" yaml:"originRequestHeaders"`
	// The configured backend type http for the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/cdn_distribution#origin_url CdnDistribution#origin_url}
	OriginUrl *string `field:"optional" json:"originUrl" yaml:"originUrl"`
	// The region where the bucket is hosted. Required if type is 'bucket'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/cdn_distribution#region CdnDistribution#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

