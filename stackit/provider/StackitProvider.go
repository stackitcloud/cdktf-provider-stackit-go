package provider

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/provider/internal"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.30.0/docs stackit}.
type StackitProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	ArgusCustomEndpoint() *string
	SetArgusCustomEndpoint(val *string)
	ArgusCustomEndpointInput() *string
	AuthorizationCustomEndpoint() *string
	SetAuthorizationCustomEndpoint(val *string)
	AuthorizationCustomEndpointInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CredentialsPath() *string
	SetCredentialsPath(val *string)
	CredentialsPathInput() *string
	DnsCustomEndpoint() *string
	SetDnsCustomEndpoint(val *string)
	DnsCustomEndpointInput() *string
	EnableBetaResources() interface{}
	SetEnableBetaResources(val interface{})
	EnableBetaResourcesInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IaasCustomEndpoint() *string
	SetIaasCustomEndpoint(val *string)
	IaasCustomEndpointInput() *string
	JwksCustomEndpoint() *string
	SetJwksCustomEndpoint(val *string)
	JwksCustomEndpointInput() *string
	LoadbalancerCustomEndpoint() *string
	SetLoadbalancerCustomEndpoint(val *string)
	LoadbalancerCustomEndpointInput() *string
	LogmeCustomEndpoint() *string
	SetLogmeCustomEndpoint(val *string)
	LogmeCustomEndpointInput() *string
	MariadbCustomEndpoint() *string
	SetMariadbCustomEndpoint(val *string)
	MariadbCustomEndpointInput() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	MongodbflexCustomEndpoint() *string
	SetMongodbflexCustomEndpoint(val *string)
	MongodbflexCustomEndpointInput() *string
	// The tree node.
	Node() constructs.Node
	ObjectstorageCustomEndpoint() *string
	SetObjectstorageCustomEndpoint(val *string)
	ObjectstorageCustomEndpointInput() *string
	ObservabilityCustomEndpoint() *string
	SetObservabilityCustomEndpoint(val *string)
	ObservabilityCustomEndpointInput() *string
	OpensearchCustomEndpoint() *string
	SetOpensearchCustomEndpoint(val *string)
	OpensearchCustomEndpointInput() *string
	PostgresflexCustomEndpoint() *string
	SetPostgresflexCustomEndpoint(val *string)
	PostgresflexCustomEndpointInput() *string
	PostgresqlCustomEndpoint() *string
	SetPostgresqlCustomEndpoint(val *string)
	PostgresqlCustomEndpointInput() *string
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyInput() *string
	PrivateKeyPath() *string
	SetPrivateKeyPath(val *string)
	PrivateKeyPathInput() *string
	RabbitmqCustomEndpoint() *string
	SetRabbitmqCustomEndpoint(val *string)
	RabbitmqCustomEndpointInput() *string
	// Experimental.
	RawOverrides() interface{}
	RedisCustomEndpoint() *string
	SetRedisCustomEndpoint(val *string)
	RedisCustomEndpointInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ResourcemanagerCustomEndpoint() *string
	SetResourcemanagerCustomEndpoint(val *string)
	ResourcemanagerCustomEndpointInput() *string
	SecretsmanagerCustomEndpoint() *string
	SetSecretsmanagerCustomEndpoint(val *string)
	SecretsmanagerCustomEndpointInput() *string
	ServerBackupCustomEndpoint() *string
	SetServerBackupCustomEndpoint(val *string)
	ServerBackupCustomEndpointInput() *string
	ServiceAccountEmail() *string
	SetServiceAccountEmail(val *string)
	ServiceAccountEmailInput() *string
	ServiceAccountKey() *string
	SetServiceAccountKey(val *string)
	ServiceAccountKeyInput() *string
	ServiceAccountKeyPath() *string
	SetServiceAccountKeyPath(val *string)
	ServiceAccountKeyPathInput() *string
	ServiceAccountToken() *string
	SetServiceAccountToken(val *string)
	ServiceAccountTokenInput() *string
	ServiceEnablementCustomEndpoint() *string
	SetServiceEnablementCustomEndpoint(val *string)
	ServiceEnablementCustomEndpointInput() *string
	SkeCustomEndpoint() *string
	SetSkeCustomEndpoint(val *string)
	SkeCustomEndpointInput() *string
	SqlserverflexCustomEndpoint() *string
	SetSqlserverflexCustomEndpoint(val *string)
	SqlserverflexCustomEndpointInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	TokenCustomEndpoint() *string
	SetTokenCustomEndpoint(val *string)
	TokenCustomEndpointInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetArgusCustomEndpoint()
	ResetAuthorizationCustomEndpoint()
	ResetCredentialsPath()
	ResetDnsCustomEndpoint()
	ResetEnableBetaResources()
	ResetIaasCustomEndpoint()
	ResetJwksCustomEndpoint()
	ResetLoadbalancerCustomEndpoint()
	ResetLogmeCustomEndpoint()
	ResetMariadbCustomEndpoint()
	ResetMongodbflexCustomEndpoint()
	ResetObjectstorageCustomEndpoint()
	ResetObservabilityCustomEndpoint()
	ResetOpensearchCustomEndpoint()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPostgresflexCustomEndpoint()
	ResetPostgresqlCustomEndpoint()
	ResetPrivateKey()
	ResetPrivateKeyPath()
	ResetRabbitmqCustomEndpoint()
	ResetRedisCustomEndpoint()
	ResetRegion()
	ResetResourcemanagerCustomEndpoint()
	ResetSecretsmanagerCustomEndpoint()
	ResetServerBackupCustomEndpoint()
	ResetServiceAccountEmail()
	ResetServiceAccountKey()
	ResetServiceAccountKeyPath()
	ResetServiceAccountToken()
	ResetServiceEnablementCustomEndpoint()
	ResetSkeCustomEndpoint()
	ResetSqlserverflexCustomEndpoint()
	ResetTokenCustomEndpoint()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for StackitProvider
type jsiiProxy_StackitProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_StackitProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ArgusCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"argusCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ArgusCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"argusCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) AuthorizationCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) AuthorizationCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) CredentialsPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) CredentialsPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) DnsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) DnsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) EnableBetaResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBetaResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) EnableBetaResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBetaResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) IaasCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iaasCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) IaasCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iaasCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) JwksCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) JwksCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) LoadbalancerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadbalancerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) LoadbalancerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadbalancerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) LogmeCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logmeCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) LogmeCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logmeCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) MariadbCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mariadbCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) MariadbCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mariadbCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) MongodbflexCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongodbflexCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) MongodbflexCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongodbflexCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ObjectstorageCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectstorageCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ObjectstorageCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectstorageCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ObservabilityCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"observabilityCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ObservabilityCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"observabilityCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) OpensearchCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opensearchCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) OpensearchCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opensearchCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) PostgresflexCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresflexCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) PostgresflexCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresflexCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) PostgresqlCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresqlCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) PostgresqlCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresqlCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) PrivateKeyPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) PrivateKeyPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) RabbitmqCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rabbitmqCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) RabbitmqCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rabbitmqCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) RedisCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) RedisCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ResourcemanagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcemanagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ResourcemanagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcemanagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) SecretsmanagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsmanagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) SecretsmanagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsmanagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ServerBackupCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverBackupCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ServerBackupCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverBackupCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ServiceAccountEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ServiceAccountKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ServiceAccountKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ServiceAccountKeyPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountKeyPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ServiceAccountKeyPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountKeyPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ServiceAccountToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ServiceAccountTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ServiceEnablementCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceEnablementCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) ServiceEnablementCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceEnablementCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) SkeCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skeCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) SkeCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skeCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) SqlserverflexCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlserverflexCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) SqlserverflexCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlserverflexCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) TokenCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackitProvider) TokenCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenCustomEndpointInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.30.0/docs stackit} Resource.
func NewStackitProvider(scope constructs.Construct, id *string, config *StackitProviderConfig) StackitProvider {
	_init_.Initialize()

	if err := validateNewStackitProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StackitProvider{}

	_jsii_.Create(
		"stackit.provider.StackitProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.30.0/docs stackit} Resource.
func NewStackitProvider_Override(s StackitProvider, scope constructs.Construct, id *string, config *StackitProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.provider.StackitProvider",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StackitProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetArgusCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"argusCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetAuthorizationCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"authorizationCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetCredentialsPath(val *string) {
	_jsii_.Set(
		j,
		"credentialsPath",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetDnsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dnsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetEnableBetaResources(val interface{}) {
	if err := j.validateSetEnableBetaResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBetaResources",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetIaasCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"iaasCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetJwksCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"jwksCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetLoadbalancerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"loadbalancerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetLogmeCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"logmeCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetMariadbCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"mariadbCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetMongodbflexCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"mongodbflexCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetObjectstorageCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"objectstorageCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetObservabilityCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"observabilityCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetOpensearchCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"opensearchCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetPostgresflexCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"postgresflexCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetPostgresqlCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"postgresqlCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetPrivateKey(val *string) {
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetPrivateKeyPath(val *string) {
	_jsii_.Set(
		j,
		"privateKeyPath",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetRabbitmqCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"rabbitmqCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetRedisCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"redisCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetResourcemanagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"resourcemanagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetSecretsmanagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"secretsmanagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetServerBackupCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"serverBackupCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetServiceAccountEmail(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountEmail",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetServiceAccountKey(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountKey",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetServiceAccountKeyPath(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountKeyPath",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetServiceAccountToken(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountToken",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetServiceEnablementCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"serviceEnablementCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetSkeCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"skeCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetSqlserverflexCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"sqlserverflexCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_StackitProvider)SetTokenCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"tokenCustomEndpoint",
		val,
	)
}

// Generates CDKTF code for importing a StackitProvider resource upon running "cdktf plan <stack-name>".
func StackitProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateStackitProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"stackit.provider.StackitProvider",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func StackitProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStackitProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.provider.StackitProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StackitProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStackitProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.provider.StackitProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StackitProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStackitProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.provider.StackitProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StackitProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"stackit.provider.StackitProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StackitProvider) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StackitProvider) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StackitProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		s,
		"resetAlias",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetArgusCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetArgusCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetAuthorizationCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetAuthorizationCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetCredentialsPath() {
	_jsii_.InvokeVoid(
		s,
		"resetCredentialsPath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetDnsCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetDnsCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetEnableBetaResources() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableBetaResources",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetIaasCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetIaasCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetJwksCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetJwksCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetLoadbalancerCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetLoadbalancerCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetLogmeCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetLogmeCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetMariadbCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetMariadbCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetMongodbflexCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetMongodbflexCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetObjectstorageCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectstorageCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetObservabilityCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetObservabilityCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetOpensearchCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetOpensearchCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetPostgresflexCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetPostgresflexCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetPostgresqlCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetPostgresqlCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetPrivateKey() {
	_jsii_.InvokeVoid(
		s,
		"resetPrivateKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetPrivateKeyPath() {
	_jsii_.InvokeVoid(
		s,
		"resetPrivateKeyPath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetRabbitmqCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetRabbitmqCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetRedisCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetRedisCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetRegion() {
	_jsii_.InvokeVoid(
		s,
		"resetRegion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetResourcemanagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetResourcemanagerCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetSecretsmanagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetSecretsmanagerCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetServerBackupCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetServerBackupCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetServiceAccountEmail() {
	_jsii_.InvokeVoid(
		s,
		"resetServiceAccountEmail",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetServiceAccountKey() {
	_jsii_.InvokeVoid(
		s,
		"resetServiceAccountKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetServiceAccountKeyPath() {
	_jsii_.InvokeVoid(
		s,
		"resetServiceAccountKeyPath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetServiceAccountToken() {
	_jsii_.InvokeVoid(
		s,
		"resetServiceAccountToken",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetServiceEnablementCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetServiceEnablementCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetSkeCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetSkeCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetSqlserverflexCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetSqlserverflexCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) ResetTokenCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetTokenCustomEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackitProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackitProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackitProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackitProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackitProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackitProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

