package argusscrapeconfig

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/argusscrapeconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.4.0/docs/resources/argus_scrapeconfig stackit_argus_scrapeconfig}.
type ArgusScrapeconfig interface {
	cdktf.TerraformResource
	BasicAuth() ArgusScrapeconfigBasicAuthOutputReference
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
	Saml2() ArgusScrapeconfigSaml2OutputReference
	Saml2Input() interface{}
	Scheme() *string
	SetScheme(val *string)
	SchemeInput() *string
	ScrapeInterval() *string
	SetScrapeInterval(val *string)
	ScrapeIntervalInput() *string
	ScrapeTimeout() *string
	SetScrapeTimeout(val *string)
	ScrapeTimeoutInput() *string
	Targets() ArgusScrapeconfigTargetsList
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
	PutBasicAuth(value *ArgusScrapeconfigBasicAuth)
	PutSaml2(value *ArgusScrapeconfigSaml2)
	PutTargets(value interface{})
	ResetBasicAuth()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSaml2()
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

// The jsii proxy struct for ArgusScrapeconfig
type jsiiProxy_ArgusScrapeconfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ArgusScrapeconfig) BasicAuth() ArgusScrapeconfigBasicAuthOutputReference {
	var returns ArgusScrapeconfigBasicAuthOutputReference
	_jsii_.Get(
		j,
		"basicAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) BasicAuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"basicAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) MetricsPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) MetricsPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) Saml2() ArgusScrapeconfigSaml2OutputReference {
	var returns ArgusScrapeconfigSaml2OutputReference
	_jsii_.Get(
		j,
		"saml2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) Saml2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"saml2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) Scheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) SchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) ScrapeInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scrapeInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) ScrapeIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scrapeIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) ScrapeTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scrapeTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) ScrapeTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scrapeTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) Targets() ArgusScrapeconfigTargetsList {
	var returns ArgusScrapeconfigTargetsList
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) TargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusScrapeconfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.4.0/docs/resources/argus_scrapeconfig stackit_argus_scrapeconfig} Resource.
func NewArgusScrapeconfig(scope constructs.Construct, id *string, config *ArgusScrapeconfigConfig) ArgusScrapeconfig {
	_init_.Initialize()

	if err := validateNewArgusScrapeconfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArgusScrapeconfig{}

	_jsii_.Create(
		"stackit.argusScrapeconfig.ArgusScrapeconfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.4.0/docs/resources/argus_scrapeconfig stackit_argus_scrapeconfig} Resource.
func NewArgusScrapeconfig_Override(a ArgusScrapeconfig, scope constructs.Construct, id *string, config *ArgusScrapeconfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.argusScrapeconfig.ArgusScrapeconfig",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ArgusScrapeconfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ArgusScrapeconfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ArgusScrapeconfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ArgusScrapeconfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ArgusScrapeconfig)SetInstanceId(val *string) {
	if err := j.validateSetInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_ArgusScrapeconfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ArgusScrapeconfig)SetMetricsPath(val *string) {
	if err := j.validateSetMetricsPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsPath",
		val,
	)
}

func (j *jsiiProxy_ArgusScrapeconfig)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ArgusScrapeconfig)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_ArgusScrapeconfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ArgusScrapeconfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ArgusScrapeconfig)SetScheme(val *string) {
	if err := j.validateSetSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheme",
		val,
	)
}

func (j *jsiiProxy_ArgusScrapeconfig)SetScrapeInterval(val *string) {
	if err := j.validateSetScrapeIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scrapeInterval",
		val,
	)
}

func (j *jsiiProxy_ArgusScrapeconfig)SetScrapeTimeout(val *string) {
	if err := j.validateSetScrapeTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scrapeTimeout",
		val,
	)
}

// Generates CDKTF code for importing a ArgusScrapeconfig resource upon running "cdktf plan <stack-name>".
func ArgusScrapeconfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateArgusScrapeconfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"stackit.argusScrapeconfig.ArgusScrapeconfig",
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
func ArgusScrapeconfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateArgusScrapeconfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.argusScrapeconfig.ArgusScrapeconfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ArgusScrapeconfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateArgusScrapeconfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.argusScrapeconfig.ArgusScrapeconfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ArgusScrapeconfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateArgusScrapeconfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.argusScrapeconfig.ArgusScrapeconfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ArgusScrapeconfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"stackit.argusScrapeconfig.ArgusScrapeconfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ArgusScrapeconfig) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ArgusScrapeconfig) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ArgusScrapeconfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusScrapeconfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusScrapeconfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusScrapeconfig) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusScrapeconfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusScrapeconfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusScrapeconfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusScrapeconfig) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusScrapeconfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusScrapeconfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusScrapeconfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ArgusScrapeconfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusScrapeconfig) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ArgusScrapeconfig) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ArgusScrapeconfig) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ArgusScrapeconfig) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ArgusScrapeconfig) PutBasicAuth(value *ArgusScrapeconfigBasicAuth) {
	if err := a.validatePutBasicAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBasicAuth",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArgusScrapeconfig) PutSaml2(value *ArgusScrapeconfigSaml2) {
	if err := a.validatePutSaml2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSaml2",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArgusScrapeconfig) PutTargets(value interface{}) {
	if err := a.validatePutTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTargets",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArgusScrapeconfig) ResetBasicAuth() {
	_jsii_.InvokeVoid(
		a,
		"resetBasicAuth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusScrapeconfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusScrapeconfig) ResetSaml2() {
	_jsii_.InvokeVoid(
		a,
		"resetSaml2",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusScrapeconfig) ResetScheme() {
	_jsii_.InvokeVoid(
		a,
		"resetScheme",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusScrapeconfig) ResetScrapeInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetScrapeInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusScrapeconfig) ResetScrapeTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetScrapeTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusScrapeconfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusScrapeconfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusScrapeconfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusScrapeconfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusScrapeconfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusScrapeconfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

