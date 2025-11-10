package observabilityscrapeconfig

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/observabilityscrapeconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.2/docs/resources/observability_scrapeconfig stackit_observability_scrapeconfig}.
type ObservabilityScrapeconfig interface {
	cdktf.TerraformResource
	BasicAuth() ObservabilityScrapeconfigBasicAuthOutputReference
	BasicAuthInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	InstanceId() *string
	SetInstanceId(val *string)
	InstanceIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MetricsPath() *string
	SetMetricsPath(val *string)
	MetricsPathInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Saml2() ObservabilityScrapeconfigSaml2OutputReference
	Saml2Input() interface{}
	SampleLimit() *float64
	SetSampleLimit(val *float64)
	SampleLimitInput() *float64
	Scheme() *string
	SetScheme(val *string)
	SchemeInput() *string
	ScrapeInterval() *string
	SetScrapeInterval(val *string)
	ScrapeIntervalInput() *string
	ScrapeTimeout() *string
	SetScrapeTimeout(val *string)
	ScrapeTimeoutInput() *string
	Targets() ObservabilityScrapeconfigTargetsList
	TargetsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutBasicAuth(value *ObservabilityScrapeconfigBasicAuth)
	PutSaml2(value *ObservabilityScrapeconfigSaml2)
	PutTargets(value interface{})
	ResetBasicAuth()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSaml2()
	ResetSampleLimit()
	ResetScheme()
	ResetScrapeInterval()
	ResetScrapeTimeout()
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

// The jsii proxy struct for ObservabilityScrapeconfig
type jsiiProxy_ObservabilityScrapeconfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ObservabilityScrapeconfig) BasicAuth() ObservabilityScrapeconfigBasicAuthOutputReference {
	var returns ObservabilityScrapeconfigBasicAuthOutputReference
	_jsii_.Get(
		j,
		"basicAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) BasicAuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"basicAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) MetricsPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) MetricsPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) Saml2() ObservabilityScrapeconfigSaml2OutputReference {
	var returns ObservabilityScrapeconfigSaml2OutputReference
	_jsii_.Get(
		j,
		"saml2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) Saml2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"saml2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) SampleLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) SampleLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) Scheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) SchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) ScrapeInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scrapeInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) ScrapeIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scrapeIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) ScrapeTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scrapeTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) ScrapeTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scrapeTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) Targets() ObservabilityScrapeconfigTargetsList {
	var returns ObservabilityScrapeconfigTargetsList
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) TargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityScrapeconfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.2/docs/resources/observability_scrapeconfig stackit_observability_scrapeconfig} Resource.
func NewObservabilityScrapeconfig(scope constructs.Construct, id *string, config *ObservabilityScrapeconfigConfig) ObservabilityScrapeconfig {
	_init_.Initialize()

	if err := validateNewObservabilityScrapeconfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityScrapeconfig{}

	_jsii_.Create(
		"stackit.observabilityScrapeconfig.ObservabilityScrapeconfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.58.2/docs/resources/observability_scrapeconfig stackit_observability_scrapeconfig} Resource.
func NewObservabilityScrapeconfig_Override(o ObservabilityScrapeconfig, scope constructs.Construct, id *string, config *ObservabilityScrapeconfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.observabilityScrapeconfig.ObservabilityScrapeconfig",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_ObservabilityScrapeconfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ObservabilityScrapeconfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ObservabilityScrapeconfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ObservabilityScrapeconfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ObservabilityScrapeconfig)SetInstanceId(val *string) {
	if err := j.validateSetInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_ObservabilityScrapeconfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ObservabilityScrapeconfig)SetMetricsPath(val *string) {
	if err := j.validateSetMetricsPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsPath",
		val,
	)
}

func (j *jsiiProxy_ObservabilityScrapeconfig)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ObservabilityScrapeconfig)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_ObservabilityScrapeconfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ObservabilityScrapeconfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ObservabilityScrapeconfig)SetSampleLimit(val *float64) {
	if err := j.validateSetSampleLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleLimit",
		val,
	)
}

func (j *jsiiProxy_ObservabilityScrapeconfig)SetScheme(val *string) {
	if err := j.validateSetSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheme",
		val,
	)
}

func (j *jsiiProxy_ObservabilityScrapeconfig)SetScrapeInterval(val *string) {
	if err := j.validateSetScrapeIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scrapeInterval",
		val,
	)
}

func (j *jsiiProxy_ObservabilityScrapeconfig)SetScrapeTimeout(val *string) {
	if err := j.validateSetScrapeTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scrapeTimeout",
		val,
	)
}

// Generates CDKTF code for importing a ObservabilityScrapeconfig resource upon running "cdktf plan <stack-name>".
func ObservabilityScrapeconfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateObservabilityScrapeconfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"stackit.observabilityScrapeconfig.ObservabilityScrapeconfig",
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
func ObservabilityScrapeconfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateObservabilityScrapeconfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.observabilityScrapeconfig.ObservabilityScrapeconfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ObservabilityScrapeconfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateObservabilityScrapeconfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.observabilityScrapeconfig.ObservabilityScrapeconfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ObservabilityScrapeconfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateObservabilityScrapeconfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.observabilityScrapeconfig.ObservabilityScrapeconfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ObservabilityScrapeconfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"stackit.observabilityScrapeconfig.ObservabilityScrapeconfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_ObservabilityScrapeconfig) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_ObservabilityScrapeconfig) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_ObservabilityScrapeconfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityScrapeconfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityScrapeconfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityScrapeconfig) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityScrapeconfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityScrapeconfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityScrapeconfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityScrapeconfig) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityScrapeconfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityScrapeconfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityScrapeconfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_ObservabilityScrapeconfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityScrapeconfig) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_ObservabilityScrapeconfig) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_ObservabilityScrapeconfig) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_ObservabilityScrapeconfig) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_ObservabilityScrapeconfig) PutBasicAuth(value *ObservabilityScrapeconfigBasicAuth) {
	if err := o.validatePutBasicAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putBasicAuth",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityScrapeconfig) PutSaml2(value *ObservabilityScrapeconfigSaml2) {
	if err := o.validatePutSaml2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSaml2",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityScrapeconfig) PutTargets(value interface{}) {
	if err := o.validatePutTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTargets",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityScrapeconfig) ResetBasicAuth() {
	_jsii_.InvokeVoid(
		o,
		"resetBasicAuth",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityScrapeconfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityScrapeconfig) ResetSaml2() {
	_jsii_.InvokeVoid(
		o,
		"resetSaml2",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityScrapeconfig) ResetSampleLimit() {
	_jsii_.InvokeVoid(
		o,
		"resetSampleLimit",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityScrapeconfig) ResetScheme() {
	_jsii_.InvokeVoid(
		o,
		"resetScheme",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityScrapeconfig) ResetScrapeInterval() {
	_jsii_.InvokeVoid(
		o,
		"resetScrapeInterval",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityScrapeconfig) ResetScrapeTimeout() {
	_jsii_.InvokeVoid(
		o,
		"resetScrapeTimeout",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityScrapeconfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityScrapeconfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityScrapeconfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityScrapeconfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityScrapeconfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityScrapeconfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

