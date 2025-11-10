package datastackitobservabilityinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/datastackitobservabilityinstance/internal"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.56.0/docs/data-sources/observability_instance stackit_observability_instance}.
type DataStackitObservabilityInstance interface {
	cdktf.TerraformDataSource
	Acl() *[]*string
	AlertConfig() DataStackitObservabilityInstanceAlertConfigOutputReference
	AlertingUrl() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	SetInstanceId(val *string)
	InstanceIdInput() *string
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
	MetricsRetentionDays1HDownsampling() *float64
	MetricsRetentionDays5MDownsampling() *float64
	MetricsUrl() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	OtlpTracesUrl() *string
	Parameters() cdktf.StringMap
	PlanId() *string
	PlanName() *string
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataStackitObservabilityInstance
type jsiiProxy_DataStackitObservabilityInstance struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataStackitObservabilityInstance) Acl() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) AlertConfig() DataStackitObservabilityInstanceAlertConfigOutputReference {
	var returns DataStackitObservabilityInstanceAlertConfigOutputReference
	_jsii_.Get(
		j,
		"alertConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) AlertingUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertingUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) DashboardUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) GrafanaInitialAdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaInitialAdminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) GrafanaInitialAdminUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaInitialAdminUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) GrafanaPublicReadAccess() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"grafanaPublicReadAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) GrafanaUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grafanaUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) IsUpdatable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isUpdatable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) JaegerTracesUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jaegerTracesUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) JaegerUiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jaegerUiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) LogsPushUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logsPushUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) LogsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) MetricsPushUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPushUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) MetricsRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) MetricsRetentionDays1HDownsampling() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsRetentionDays1HDownsampling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) MetricsRetentionDays5MDownsampling() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsRetentionDays5MDownsampling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) MetricsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) OtlpTracesUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"otlpTracesUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) Parameters() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) PlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) PlanName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) TargetsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstance) ZipkinSpansUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipkinSpansUrl",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.56.0/docs/data-sources/observability_instance stackit_observability_instance} Data Source.
func NewDataStackitObservabilityInstance(scope constructs.Construct, id *string, config *DataStackitObservabilityInstanceConfig) DataStackitObservabilityInstance {
	_init_.Initialize()

	if err := validateNewDataStackitObservabilityInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataStackitObservabilityInstance{}

	_jsii_.Create(
		"stackit.dataStackitObservabilityInstance.DataStackitObservabilityInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.56.0/docs/data-sources/observability_instance stackit_observability_instance} Data Source.
func NewDataStackitObservabilityInstance_Override(d DataStackitObservabilityInstance, scope constructs.Construct, id *string, config *DataStackitObservabilityInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.dataStackitObservabilityInstance.DataStackitObservabilityInstance",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataStackitObservabilityInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataStackitObservabilityInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataStackitObservabilityInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataStackitObservabilityInstance)SetInstanceId(val *string) {
	if err := j.validateSetInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_DataStackitObservabilityInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataStackitObservabilityInstance)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_DataStackitObservabilityInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataStackitObservabilityInstance resource upon running "cdktf plan <stack-name>".
func DataStackitObservabilityInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataStackitObservabilityInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"stackit.dataStackitObservabilityInstance.DataStackitObservabilityInstance",
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
func DataStackitObservabilityInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataStackitObservabilityInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.dataStackitObservabilityInstance.DataStackitObservabilityInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataStackitObservabilityInstance_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataStackitObservabilityInstance_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.dataStackitObservabilityInstance.DataStackitObservabilityInstance",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataStackitObservabilityInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataStackitObservabilityInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.dataStackitObservabilityInstance.DataStackitObservabilityInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataStackitObservabilityInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"stackit.dataStackitObservabilityInstance.DataStackitObservabilityInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataStackitObservabilityInstance) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataStackitObservabilityInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitObservabilityInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitObservabilityInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitObservabilityInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitObservabilityInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitObservabilityInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitObservabilityInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitObservabilityInstance) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitObservabilityInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitObservabilityInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitObservabilityInstance) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataStackitObservabilityInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataStackitObservabilityInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitObservabilityInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitObservabilityInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitObservabilityInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitObservabilityInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitObservabilityInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

