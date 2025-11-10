package argusinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/argusinstance/internal"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/argus_instance stackit_argus_instance}.
type ArgusInstance interface {
	cdktf.TerraformResource
	Acl() *[]*string
	SetAcl(val *[]*string)
	AclInput() *[]*string
	AlertConfig() ArgusInstanceAlertConfigOutputReference
	AlertConfigInput() interface{}
	AlertingUrl() *string
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
	DashboardUrl() *string
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
	GrafanaInitialAdminPassword() *string
	GrafanaInitialAdminUser() *string
	GrafanaPublicReadAccess() cdktf.IResolvable
	GrafanaUrl() *string
	Id() *string
	InstanceId() *string
	IsUpdatable() cdktf.IResolvable
	JaegerTracesUrl() *string
	JaegerUiUrl() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogsPushUrl() *string
	LogsUrl() *string
	MetricsPushUrl() *string
	MetricsRetentionDays() *float64
	SetMetricsRetentionDays(val *float64)
	MetricsRetentionDays1HDownsampling() *float64
	SetMetricsRetentionDays1HDownsampling(val *float64)
	MetricsRetentionDays1HDownsamplingInput() *float64
	MetricsRetentionDays5MDownsampling() *float64
	SetMetricsRetentionDays5MDownsampling(val *float64)
	MetricsRetentionDays5MDownsamplingInput() *float64
	MetricsRetentionDaysInput() *float64
	MetricsUrl() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OtlpTracesUrl() *string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	PlanId() *string
	PlanName() *string
	SetPlanName(val *string)
	PlanNameInput() *string
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
	TargetsUrl() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ZipkinSpansUrl() *string
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
	PutAlertConfig(value *ArgusInstanceAlertConfig)
	ResetAcl()
	ResetAlertConfig()
	ResetMetricsRetentionDays()
	ResetMetricsRetentionDays1HDownsampling()
	ResetMetricsRetentionDays5MDownsampling()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
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

// The jsii proxy struct for ArgusInstance
type jsiiProxy_ArgusInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ArgusInstance) Acl() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) AclInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) AlertConfig() ArgusInstanceAlertConfigOutputReference {
	var returns ArgusInstanceAlertConfigOutputReference
	_jsii_.Get(
		j,
		"alertConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) AlertConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) AlertingUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertingUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) DashboardUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) GrafanaInitialAdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaInitialAdminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) GrafanaInitialAdminUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaInitialAdminUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) GrafanaPublicReadAccess() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"grafanaPublicReadAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) GrafanaUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) IsUpdatable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isUpdatable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) JaegerTracesUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jaegerTracesUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) JaegerUiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jaegerUiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) LogsPushUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logsPushUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) LogsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) MetricsPushUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPushUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) MetricsRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) MetricsRetentionDays1HDownsampling() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsRetentionDays1HDownsampling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) MetricsRetentionDays1HDownsamplingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsRetentionDays1HDownsamplingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) MetricsRetentionDays5MDownsampling() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsRetentionDays5MDownsampling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) MetricsRetentionDays5MDownsamplingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsRetentionDays5MDownsamplingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) MetricsRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) MetricsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) OtlpTracesUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"otlpTracesUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) PlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) PlanName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) PlanNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) TargetsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstance) ZipkinSpansUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipkinSpansUrl",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/argus_instance stackit_argus_instance} Resource.
func NewArgusInstance(scope constructs.Construct, id *string, config *ArgusInstanceConfig) ArgusInstance {
	_init_.Initialize()

	if err := validateNewArgusInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArgusInstance{}

	_jsii_.Create(
		"stackit.argusInstance.ArgusInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.0/docs/resources/argus_instance stackit_argus_instance} Resource.
func NewArgusInstance_Override(a ArgusInstance, scope constructs.Construct, id *string, config *ArgusInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.argusInstance.ArgusInstance",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ArgusInstance)SetAcl(val *[]*string) {
	if err := j.validateSetAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acl",
		val,
	)
}

func (j *jsiiProxy_ArgusInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ArgusInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ArgusInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ArgusInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ArgusInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ArgusInstance)SetMetricsRetentionDays(val *float64) {
	if err := j.validateSetMetricsRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsRetentionDays",
		val,
	)
}

func (j *jsiiProxy_ArgusInstance)SetMetricsRetentionDays1HDownsampling(val *float64) {
	if err := j.validateSetMetricsRetentionDays1HDownsamplingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsRetentionDays1HDownsampling",
		val,
	)
}

func (j *jsiiProxy_ArgusInstance)SetMetricsRetentionDays5MDownsampling(val *float64) {
	if err := j.validateSetMetricsRetentionDays5MDownsamplingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsRetentionDays5MDownsampling",
		val,
	)
}

func (j *jsiiProxy_ArgusInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ArgusInstance)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_ArgusInstance)SetPlanName(val *string) {
	if err := j.validateSetPlanNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"planName",
		val,
	)
}

func (j *jsiiProxy_ArgusInstance)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_ArgusInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ArgusInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a ArgusInstance resource upon running "cdktf plan <stack-name>".
func ArgusInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateArgusInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"stackit.argusInstance.ArgusInstance",
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
func ArgusInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateArgusInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.argusInstance.ArgusInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ArgusInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateArgusInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.argusInstance.ArgusInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ArgusInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateArgusInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.argusInstance.ArgusInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ArgusInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"stackit.argusInstance.ArgusInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ArgusInstance) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ArgusInstance) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ArgusInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ArgusInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ArgusInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ArgusInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ArgusInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ArgusInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ArgusInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ArgusInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ArgusInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ArgusInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstance) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ArgusInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ArgusInstance) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ArgusInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ArgusInstance) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ArgusInstance) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ArgusInstance) PutAlertConfig(value *ArgusInstanceAlertConfig) {
	if err := a.validatePutAlertConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAlertConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArgusInstance) ResetAcl() {
	_jsii_.InvokeVoid(
		a,
		"resetAcl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstance) ResetAlertConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAlertConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstance) ResetMetricsRetentionDays() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricsRetentionDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstance) ResetMetricsRetentionDays1HDownsampling() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricsRetentionDays1HDownsampling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstance) ResetMetricsRetentionDays5MDownsampling() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricsRetentionDays5MDownsampling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstance) ResetParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

