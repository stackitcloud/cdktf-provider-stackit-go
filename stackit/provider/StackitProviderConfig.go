package provider


type StackitProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#alias StackitProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Custom endpoint for the Argus service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#argus_custom_endpoint StackitProvider#argus_custom_endpoint}
	ArgusCustomEndpoint *string `field:"optional" json:"argusCustomEndpoint" yaml:"argusCustomEndpoint"`
	// Custom endpoint for the Membership service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#authorization_custom_endpoint StackitProvider#authorization_custom_endpoint}
	AuthorizationCustomEndpoint *string `field:"optional" json:"authorizationCustomEndpoint" yaml:"authorizationCustomEndpoint"`
	// Path of JSON from where the credentials are read.
	//
	// Takes precedence over the env var `STACKIT_CREDENTIALS_PATH`. Default value is `~/.stackit/credentials.json`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#credentials_path StackitProvider#credentials_path}
	CredentialsPath *string `field:"optional" json:"credentialsPath" yaml:"credentialsPath"`
	// Region will be used as the default location for regional services.
	//
	// Not all services require a region, some are global
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#default_region StackitProvider#default_region}
	DefaultRegion *string `field:"optional" json:"defaultRegion" yaml:"defaultRegion"`
	// Custom endpoint for the DNS service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#dns_custom_endpoint StackitProvider#dns_custom_endpoint}
	DnsCustomEndpoint *string `field:"optional" json:"dnsCustomEndpoint" yaml:"dnsCustomEndpoint"`
	// Enable beta resources. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#enable_beta_resources StackitProvider#enable_beta_resources}
	EnableBetaResources interface{} `field:"optional" json:"enableBetaResources" yaml:"enableBetaResources"`
	// Enables experiments.
	//
	// These are unstable features without official support. More information can be found in the README. Available Experiments: [iam]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#experiments StackitProvider#experiments}
	Experiments *[]*string `field:"optional" json:"experiments" yaml:"experiments"`
	// Custom endpoint for the IaaS service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#iaas_custom_endpoint StackitProvider#iaas_custom_endpoint}
	IaasCustomEndpoint *string `field:"optional" json:"iaasCustomEndpoint" yaml:"iaasCustomEndpoint"`
	// Custom endpoint for the Load Balancer service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#loadbalancer_custom_endpoint StackitProvider#loadbalancer_custom_endpoint}
	LoadbalancerCustomEndpoint *string `field:"optional" json:"loadbalancerCustomEndpoint" yaml:"loadbalancerCustomEndpoint"`
	// Custom endpoint for the LogMe service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#logme_custom_endpoint StackitProvider#logme_custom_endpoint}
	LogmeCustomEndpoint *string `field:"optional" json:"logmeCustomEndpoint" yaml:"logmeCustomEndpoint"`
	// Custom endpoint for the MariaDB service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#mariadb_custom_endpoint StackitProvider#mariadb_custom_endpoint}
	MariadbCustomEndpoint *string `field:"optional" json:"mariadbCustomEndpoint" yaml:"mariadbCustomEndpoint"`
	// Custom endpoint for the MongoDB Flex service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#mongodbflex_custom_endpoint StackitProvider#mongodbflex_custom_endpoint}
	MongodbflexCustomEndpoint *string `field:"optional" json:"mongodbflexCustomEndpoint" yaml:"mongodbflexCustomEndpoint"`
	// Custom endpoint for the Object Storage service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#objectstorage_custom_endpoint StackitProvider#objectstorage_custom_endpoint}
	ObjectstorageCustomEndpoint *string `field:"optional" json:"objectstorageCustomEndpoint" yaml:"objectstorageCustomEndpoint"`
	// Custom endpoint for the Observability service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#observability_custom_endpoint StackitProvider#observability_custom_endpoint}
	ObservabilityCustomEndpoint *string `field:"optional" json:"observabilityCustomEndpoint" yaml:"observabilityCustomEndpoint"`
	// Custom endpoint for the OpenSearch service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#opensearch_custom_endpoint StackitProvider#opensearch_custom_endpoint}
	OpensearchCustomEndpoint *string `field:"optional" json:"opensearchCustomEndpoint" yaml:"opensearchCustomEndpoint"`
	// Custom endpoint for the PostgresFlex service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#postgresflex_custom_endpoint StackitProvider#postgresflex_custom_endpoint}
	PostgresflexCustomEndpoint *string `field:"optional" json:"postgresflexCustomEndpoint" yaml:"postgresflexCustomEndpoint"`
	// Private RSA key used for authentication, relevant for the key flow.
	//
	// It takes precedence over the private key that is included in the service account key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#private_key StackitProvider#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Path for the private RSA key used for authentication, relevant for the key flow.
	//
	// It takes precedence over the private key that is included in the service account key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#private_key_path StackitProvider#private_key_path}
	PrivateKeyPath *string `field:"optional" json:"privateKeyPath" yaml:"privateKeyPath"`
	// Custom endpoint for the RabbitMQ service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#rabbitmq_custom_endpoint StackitProvider#rabbitmq_custom_endpoint}
	RabbitmqCustomEndpoint *string `field:"optional" json:"rabbitmqCustomEndpoint" yaml:"rabbitmqCustomEndpoint"`
	// Custom endpoint for the Redis service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#redis_custom_endpoint StackitProvider#redis_custom_endpoint}
	RedisCustomEndpoint *string `field:"optional" json:"redisCustomEndpoint" yaml:"redisCustomEndpoint"`
	// Region will be used as the default location for regional services.
	//
	// Not all services require a region, some are global
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#region StackitProvider#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Custom endpoint for the Resource Manager service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#resourcemanager_custom_endpoint StackitProvider#resourcemanager_custom_endpoint}
	ResourcemanagerCustomEndpoint *string `field:"optional" json:"resourcemanagerCustomEndpoint" yaml:"resourcemanagerCustomEndpoint"`
	// Custom endpoint for the Secrets Manager service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#secretsmanager_custom_endpoint StackitProvider#secretsmanager_custom_endpoint}
	SecretsmanagerCustomEndpoint *string `field:"optional" json:"secretsmanagerCustomEndpoint" yaml:"secretsmanagerCustomEndpoint"`
	// Custom endpoint for the Server Backup service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#server_backup_custom_endpoint StackitProvider#server_backup_custom_endpoint}
	ServerBackupCustomEndpoint *string `field:"optional" json:"serverBackupCustomEndpoint" yaml:"serverBackupCustomEndpoint"`
	// Custom endpoint for the Server Update service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#server_update_custom_endpoint StackitProvider#server_update_custom_endpoint}
	ServerUpdateCustomEndpoint *string `field:"optional" json:"serverUpdateCustomEndpoint" yaml:"serverUpdateCustomEndpoint"`
	// Custom endpoint for the Service Account service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#service_account_custom_endpoint StackitProvider#service_account_custom_endpoint}
	ServiceAccountCustomEndpoint *string `field:"optional" json:"serviceAccountCustomEndpoint" yaml:"serviceAccountCustomEndpoint"`
	// Service account email.
	//
	// It can also be set using the environment variable STACKIT_SERVICE_ACCOUNT_EMAIL. It is required if you want to use the resource manager project resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#service_account_email StackitProvider#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// Service account key used for authentication. If set, the key flow will be used to authenticate all operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#service_account_key StackitProvider#service_account_key}
	ServiceAccountKey *string `field:"optional" json:"serviceAccountKey" yaml:"serviceAccountKey"`
	// Path for the service account key used for authentication.
	//
	// If set, the key flow will be used to authenticate all operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#service_account_key_path StackitProvider#service_account_key_path}
	ServiceAccountKeyPath *string `field:"optional" json:"serviceAccountKeyPath" yaml:"serviceAccountKeyPath"`
	// Token used for authentication. If set, the token flow will be used to authenticate all operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#service_account_token StackitProvider#service_account_token}
	ServiceAccountToken *string `field:"optional" json:"serviceAccountToken" yaml:"serviceAccountToken"`
	// Custom endpoint for the Service Enablement API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#service_enablement_custom_endpoint StackitProvider#service_enablement_custom_endpoint}
	ServiceEnablementCustomEndpoint *string `field:"optional" json:"serviceEnablementCustomEndpoint" yaml:"serviceEnablementCustomEndpoint"`
	// Custom endpoint for the Kubernetes Engine (SKE) service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#ske_custom_endpoint StackitProvider#ske_custom_endpoint}
	SkeCustomEndpoint *string `field:"optional" json:"skeCustomEndpoint" yaml:"skeCustomEndpoint"`
	// Custom endpoint for the SQL Server Flex service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#sqlserverflex_custom_endpoint StackitProvider#sqlserverflex_custom_endpoint}
	SqlserverflexCustomEndpoint *string `field:"optional" json:"sqlserverflexCustomEndpoint" yaml:"sqlserverflexCustomEndpoint"`
	// Custom endpoint for the token API, which is used to request access tokens when using the key flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.46.0/docs#token_custom_endpoint StackitProvider#token_custom_endpoint}
	TokenCustomEndpoint *string `field:"optional" json:"tokenCustomEndpoint" yaml:"tokenCustomEndpoint"`
}

