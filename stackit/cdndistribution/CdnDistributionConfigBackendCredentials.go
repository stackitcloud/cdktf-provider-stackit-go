package cdndistribution


type CdnDistributionConfigBackendCredentials struct {
	// The access key for the bucket. Required if type is 'bucket'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/cdn_distribution#access_key_id CdnDistribution#access_key_id}
	AccessKeyId *string `field:"required" json:"accessKeyId" yaml:"accessKeyId"`
	// The access key for the bucket. Required if type is 'bucket'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.82.0/docs/resources/cdn_distribution#secret_access_key CdnDistribution#secret_access_key}
	SecretAccessKey *string `field:"required" json:"secretAccessKey" yaml:"secretAccessKey"`
}

