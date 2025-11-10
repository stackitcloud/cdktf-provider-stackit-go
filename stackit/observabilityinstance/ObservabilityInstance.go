package observabilityinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/observabilityinstance/internal"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.67.0/docs/resources/observability_instance stackit_observability_instance}.
type ObservabilityInstance interface {
	cdktf.TerraformResource
	Acl() *[]*string
	SetAcl(val *[]*string)
	AclInput() *[]*string
	AlertConfig() ObservabilityInstanceAlertConfigOutputReference
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
	PutAlertConfig(value *ObservabilityInstanceAlertConfig)
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

// The jsii proxy struct for ObservabilityInstance
type jsiiProxy_ObservabilityInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ObservabilityInstance) Acl() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) AclInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) AlertConfig() ObservabilityInstanceAlertConfigOutputReference {
	var returns ObservabilityInstanceAlertConfigOutputReference
	_jsii_.Get(
		j,
		"alertConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) AlertConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) AlertingUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertingUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) DashboardUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) GrafanaInitialAdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaInitialAdminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) GrafanaInitialAdminUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaInitialAdminUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) GrafanaPublicReadAccess() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"grafanaPublicReadAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) GrafanaUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) IsUpdatable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isUpdatable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) JaegerTracesUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jaegerTracesUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) JaegerUiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jaegerUiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) LogsPushUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logsPushUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) LogsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) MetricsPushUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPushUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) MetricsRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) MetricsRetentionDays1HDownsampling() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsRetentionDays1HDownsampling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) MetricsRetentionDays1HDownsamplingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsRetentionDays1HDownsamplingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) MetricsRetentionDays5MDownsampling() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsRetentionDays5MDownsampling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) MetricsRetentionDays5MDownsamplingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsRetentionDays5MDownsamplingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) MetricsRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) MetricsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) OtlpTracesUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"otlpTracesUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) PlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) PlanName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) PlanNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) TargetsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstance) ZipkinSpansUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipkinSpansUrl",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.67.0/docs/resources/observability_instance stackit_observability_instance} Resource.
func NewObservabilityInstance(scope constructs.Construct, id *string, config *ObservabilityInstanceConfig) ObservabilityInstance {
	_init_.Initialize()

	if err := validateNewObservabilityInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityInstance{}

	_jsii_.Create(
		"stackit.observabilityInstance.ObservabilityInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.67.0/docs/resources/observability_instance stackit_observability_instance} Resource.
func NewObservabilityInstance_Override(o ObservabilityInstance, scope constructs.Construct, id *string, config *ObservabilityInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.observabilityInstance.ObservabilityInstance",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_ObservabilityInstance)SetAcl(val *[]*string) {
	if err := j.validateSetAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acl",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstance)SetMetricsRetentionDays(val *float64) {
	if err := j.validateSetMetricsRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsRetentionDays",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstance)SetMetricsRetentionDays1HDownsampling(val *float64) {
	if err := j.validateSetMetricsRetentionDays1HDownsamplingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsRetentionDays1HDownsampling",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstance)SetMetricsRetentionDays5MDownsampling(val *float64) {
	if err := j.validateSetMetricsRetentionDays5MDownsamplingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsRetentionDays5MDownsampling",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstance)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstance)SetPlanName(val *string) {
	if err := j.validateSetPlanNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"planName",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstance)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a ObservabilityInstance resource upon running "cdktf plan <stack-name>".
func ObservabilityInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateObservabilityInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"stackit.observabilityInstance.ObservabilityInstance",
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
func ObservabilityInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateObservabilityInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.observabilityInstance.ObservabilityInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ObservabilityInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateObservabilityInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.observabilityInstance.ObservabilityInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ObservabilityInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateObservabilityInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.observabilityInstance.ObservabilityInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ObservabilityInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"stackit.observabilityInstance.ObservabilityInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_ObservabilityInstance) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_ObservabilityInstance) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_ObservabilityInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityInstance) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_ObservabilityInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityInstance) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_ObservabilityInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_ObservabilityInstance) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_ObservabilityInstance) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_ObservabilityInstance) PutAlertConfig(value *ObservabilityInstanceAlertConfig) {
	if err := o.validatePutAlertConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAlertConfig",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityInstance) ResetAcl() {
	_jsii_.InvokeVoid(
		o,
		"resetAcl",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstance) ResetAlertConfig() {
	_jsii_.InvokeVoid(
		o,
		"resetAlertConfig",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstance) ResetMetricsRetentionDays() {
	_jsii_.InvokeVoid(
		o,
		"resetMetricsRetentionDays",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstance) ResetMetricsRetentionDays1HDownsampling() {
	_jsii_.InvokeVoid(
		o,
		"resetMetricsRetentionDays1HDownsampling",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstance) ResetMetricsRetentionDays5MDownsampling() {
	_jsii_.InvokeVoid(
		o,
		"resetMetricsRetentionDays5MDownsampling",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstance) ResetParameters() {
	_jsii_.InvokeVoid(
		o,
		"resetParameters",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

