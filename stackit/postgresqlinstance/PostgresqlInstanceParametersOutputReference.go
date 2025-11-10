package postgresqlinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/postgresqlinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PostgresqlInstanceParametersOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnableMonitoring() interface{}
	SetEnableMonitoring(val interface{})
	EnableMonitoringInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricsFrequency() *float64
	SetMetricsFrequency(val *float64)
	MetricsFrequencyInput() *float64
	MetricsPrefix() *string
	SetMetricsPrefix(val *string)
	MetricsPrefixInput() *string
	MonitoringInstanceId() *string
	SetMonitoringInstanceId(val *string)
	MonitoringInstanceIdInput() *string
	Plugins() *[]*string
	SetPlugins(val *[]*string)
	PluginsInput() *[]*string
	SgwAcl() *string
	SetSgwAcl(val *string)
	SgwAclInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetEnableMonitoring()
	ResetMetricsFrequency()
	ResetMetricsPrefix()
	ResetMonitoringInstanceId()
	ResetPlugins()
	ResetSgwAcl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PostgresqlInstanceParametersOutputReference
type jsiiProxy_PostgresqlInstanceParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference) EnableMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference) EnableMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference) MetricsFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference) MetricsFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference) MetricsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference) MetricsPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference) MonitoringInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference) MonitoringInstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringInstanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference) Plugins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"plugins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference) PluginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pluginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference) SgwAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sgwAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference) SgwAclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sgwAclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPostgresqlInstanceParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PostgresqlInstanceParametersOutputReference {
	_init_.Initialize()

	if err := validateNewPostgresqlInstanceParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PostgresqlInstanceParametersOutputReference{}

	_jsii_.Create(
		"stackit.postgresqlInstance.PostgresqlInstanceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPostgresqlInstanceParametersOutputReference_Override(p PostgresqlInstanceParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.postgresqlInstance.PostgresqlInstanceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference)SetEnableMonitoring(val interface{}) {
	if err := j.validateSetEnableMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMonitoring",
		val,
	)
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference)SetMetricsFrequency(val *float64) {
	if err := j.validateSetMetricsFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsFrequency",
		val,
	)
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference)SetMetricsPrefix(val *string) {
	if err := j.validateSetMetricsPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsPrefix",
		val,
	)
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference)SetMonitoringInstanceId(val *string) {
	if err := j.validateSetMonitoringInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringInstanceId",
		val,
	)
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference)SetPlugins(val *[]*string) {
	if err := j.validateSetPluginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"plugins",
		val,
	)
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference)SetSgwAcl(val *string) {
	if err := j.validateSetSgwAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sgwAcl",
		val,
	)
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PostgresqlInstanceParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) ResetEnableMonitoring() {
	_jsii_.InvokeVoid(
		p,
		"resetEnableMonitoring",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) ResetMetricsFrequency() {
	_jsii_.InvokeVoid(
		p,
		"resetMetricsFrequency",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) ResetMetricsPrefix() {
	_jsii_.InvokeVoid(
		p,
		"resetMetricsPrefix",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) ResetMonitoringInstanceId() {
	_jsii_.InvokeVoid(
		p,
		"resetMonitoringInstanceId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) ResetPlugins() {
	_jsii_.InvokeVoid(
		p,
		"resetPlugins",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) ResetSgwAcl() {
	_jsii_.InvokeVoid(
		p,
		"resetSgwAcl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlInstanceParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

