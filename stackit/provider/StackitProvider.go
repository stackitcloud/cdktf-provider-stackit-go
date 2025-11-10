package provider

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/provider/internal"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs stackit}.
type StackitProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	ArgusCustomEndpoint() *string
	SetArgusCustomEndpoint(val *string)
	ArgusCustomEndpointInput() *string
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
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	LogmeCustomEndpoint() *string
	SetLogmeCustomEndpoint(val *string)
	LogmeCustomEndpointInput() *string
	MariadbCustomEndpoint() *string
	SetMariadbCustomEndpoint(val *string)
	MariadbCustomEndpointInput() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	OpensearchCustomEndpoint() *string
	SetOpensearchCustomEndpoint(val *string)
	OpensearchCustomEndpointInput() *string
	PostgresflexCustomEndpoint() *string
	SetPostgresflexCustomEndpoint(val *string)
	PostgresflexCustomEndpointInput() *string
	PostgresqlCustomEndpoint() *string
	SetPostgresqlCustomEndpoint(val *string)
	PostgresqlCustomEndpointInput() *string
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
	ServiceAccountEmail() *string
	SetServiceAccountEmail(val *string)
	ServiceAccountEmailInput() *string
	ServiceAccountToken() *string
	SetServiceAccountToken(val *string)
	ServiceAccountTokenInput() *string
	SkeCustomEndpoint() *string
	SetSkeCustomEndpoint(val *string)
	SkeCustomEndpointInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetArgusCustomEndpoint()
	ResetCredentialsPath()
	ResetDnsCustomEndpoint()
	ResetLogmeCustomEndpoint()
	ResetMariadbCustomEndpoint()
	ResetOpensearchCustomEndpoint()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPostgresflexCustomEndpoint()
	ResetPostgresqlCustomEndpoint()
	ResetRabbitmqCustomEndpoint()
	ResetRedisCustomEndpoint()
	ResetRegion()
	ResetResourcemanagerCustomEndpoint()
	ResetServiceAccountEmail()
	ResetServiceAccountToken()
	ResetSkeCustomEndpoint()
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

func (j *jsiiProxy_StackitProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
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


// Create a new {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs stackit} Resource.
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

// Create a new {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.2.0/docs stackit} Resource.
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

func (j *jsiiProxy_StackitProvider)SetServiceAccountEmail(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountEmail",
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

func (j *jsiiProxy_StackitProvider)SetSkeCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"skeCustomEndpoint",
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

func (s *jsiiProxy_StackitProvider) ResetServiceAccountEmail() {
	_jsii_.InvokeVoid(
		s,
		"resetServiceAccountEmail",
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

func (s *jsiiProxy_StackitProvider) ResetSkeCustomEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetSkeCustomEndpoint",
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

