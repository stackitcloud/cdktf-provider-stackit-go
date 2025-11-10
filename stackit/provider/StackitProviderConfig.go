package provider


type StackitProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs#alias StackitProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Custom endpoint for the Argus service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs#argus_custom_endpoint StackitProvider#argus_custom_endpoint}
	ArgusCustomEndpoint *string `field:"optional" json:"argusCustomEndpoint" yaml:"argusCustomEndpoint"`
	// Path of JSON from where the credentials are read.
	//
	// Takes precedence over the env var `STACKIT_CREDENTIALS_PATH`. Default value is `~/.stackit/credentials.json`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs#credentials_path StackitProvider#credentials_path}
	CredentialsPath *string `field:"optional" json:"credentialsPath" yaml:"credentialsPath"`
	// Custom endpoint for the DNS service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs#dns_custom_endpoint StackitProvider#dns_custom_endpoint}
	DnsCustomEndpoint *string `field:"optional" json:"dnsCustomEndpoint" yaml:"dnsCustomEndpoint"`
	// Custom endpoint for the LogMe service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs#logme_custom_endpoint StackitProvider#logme_custom_endpoint}
	LogmeCustomEndpoint *string `field:"optional" json:"logmeCustomEndpoint" yaml:"logmeCustomEndpoint"`
	// Custom endpoint for the MariaDB service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs#mariadb_custom_endpoint StackitProvider#mariadb_custom_endpoint}
	MariadbCustomEndpoint *string `field:"optional" json:"mariadbCustomEndpoint" yaml:"mariadbCustomEndpoint"`
	// Custom endpoint for the OpenSearch service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs#opensearch_custom_endpoint StackitProvider#opensearch_custom_endpoint}
	OpensearchCustomEndpoint *string `field:"optional" json:"opensearchCustomEndpoint" yaml:"opensearchCustomEndpoint"`
	// Custom endpoint for the PostgresFlex service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs#postgresflex_custom_endpoint StackitProvider#postgresflex_custom_endpoint}
	PostgresflexCustomEndpoint *string `field:"optional" json:"postgresflexCustomEndpoint" yaml:"postgresflexCustomEndpoint"`
	// Custom endpoint for the PostgreSQL service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs#postgresql_custom_endpoint StackitProvider#postgresql_custom_endpoint}
	PostgresqlCustomEndpoint *string `field:"optional" json:"postgresqlCustomEndpoint" yaml:"postgresqlCustomEndpoint"`
	// Custom endpoint for the RabbitMQ service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs#rabbitmq_custom_endpoint StackitProvider#rabbitmq_custom_endpoint}
	RabbitmqCustomEndpoint *string `field:"optional" json:"rabbitmqCustomEndpoint" yaml:"rabbitmqCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs#redis_custom_endpoint StackitProvider#redis_custom_endpoint}.
	RedisCustomEndpoint *string `field:"optional" json:"redisCustomEndpoint" yaml:"redisCustomEndpoint"`
	// Region will be used as the default location for regional services.
	//
	// Not all services require a region, some are global
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs#region StackitProvider#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Custom endpoint for the Resource Manager service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs#resourcemanager_custom_endpoint StackitProvider#resourcemanager_custom_endpoint}
	ResourcemanagerCustomEndpoint *string `field:"optional" json:"resourcemanagerCustomEndpoint" yaml:"resourcemanagerCustomEndpoint"`
	// Service account email. It can also be set using the environment variable STACKIT_SERVICE_ACCOUNT_EMAIL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs#service_account_email StackitProvider#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// Token used for authentication. If set, the token flow will be used to authenticate all operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs#service_account_token StackitProvider#service_account_token}
	ServiceAccountToken *string `field:"optional" json:"serviceAccountToken" yaml:"serviceAccountToken"`
	// Custom endpoint for the Kubernetes Engine (SKE) service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs#ske_custom_endpoint StackitProvider#ske_custom_endpoint}
	SkeCustomEndpoint *string `field:"optional" json:"skeCustomEndpoint" yaml:"skeCustomEndpoint"`
}

