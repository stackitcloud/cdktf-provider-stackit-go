package mariadbinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/mariadbinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MariadbInstanceParametersOutputReference interface {
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
	Graphite() *string
	SetGraphite(val *string)
	GraphiteInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxDiskThreshold() *float64
	SetMaxDiskThreshold(val *float64)
	MaxDiskThresholdInput() *float64
	MetricsFrequency() *float64
	SetMetricsFrequency(val *float64)
	MetricsFrequencyInput() *float64
	MetricsPrefix() *string
	SetMetricsPrefix(val *string)
	MetricsPrefixInput() *string
	MonitoringInstanceId() *string
	SetMonitoringInstanceId(val *string)
	MonitoringInstanceIdInput() *string
	SgwAcl() *string
	SetSgwAcl(val *string)
	SgwAclInput() *string
	Syslog() *[]*string
	SetSyslog(val *[]*string)
	SyslogInput() *[]*string
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
	ResetGraphite()
	ResetMaxDiskThreshold()
	ResetMetricsFrequency()
	ResetMetricsPrefix()
	ResetMonitoringInstanceId()
	ResetSgwAcl()
	ResetSyslog()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MariadbInstanceParametersOutputReference
type jsiiProxy_MariadbInstanceParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) EnableMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) EnableMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) Graphite() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) GraphiteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) MaxDiskThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDiskThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) MaxDiskThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDiskThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) MetricsFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) MetricsFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) MetricsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) MetricsPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) MonitoringInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) MonitoringInstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringInstanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) SgwAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sgwAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) SgwAclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sgwAclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) Syslog() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"syslog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) SyslogInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"syslogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMariadbInstanceParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MariadbInstanceParametersOutputReference {
	_init_.Initialize()

	if err := validateNewMariadbInstanceParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MariadbInstanceParametersOutputReference{}

	_jsii_.Create(
		"stackit.mariadbInstance.MariadbInstanceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMariadbInstanceParametersOutputReference_Override(m MariadbInstanceParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.mariadbInstance.MariadbInstanceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference)SetEnableMonitoring(val interface{}) {
	if err := j.validateSetEnableMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMonitoring",
		val,
	)
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference)SetGraphite(val *string) {
	if err := j.validateSetGraphiteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"graphite",
		val,
	)
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference)SetMaxDiskThreshold(val *float64) {
	if err := j.validateSetMaxDiskThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDiskThreshold",
		val,
	)
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference)SetMetricsFrequency(val *float64) {
	if err := j.validateSetMetricsFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsFrequency",
		val,
	)
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference)SetMetricsPrefix(val *string) {
	if err := j.validateSetMetricsPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsPrefix",
		val,
	)
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference)SetMonitoringInstanceId(val *string) {
	if err := j.validateSetMonitoringInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringInstanceId",
		val,
	)
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference)SetSgwAcl(val *string) {
	if err := j.validateSetSgwAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sgwAcl",
		val,
	)
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference)SetSyslog(val *[]*string) {
	if err := j.validateSetSyslogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syslog",
		val,
	)
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MariadbInstanceParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) ResetEnableMonitoring() {
	_jsii_.InvokeVoid(
		m,
		"resetEnableMonitoring",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) ResetGraphite() {
	_jsii_.InvokeVoid(
		m,
		"resetGraphite",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) ResetMaxDiskThreshold() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxDiskThreshold",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) ResetMetricsFrequency() {
	_jsii_.InvokeVoid(
		m,
		"resetMetricsFrequency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) ResetMetricsPrefix() {
	_jsii_.InvokeVoid(
		m,
		"resetMetricsPrefix",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) ResetMonitoringInstanceId() {
	_jsii_.InvokeVoid(
		m,
		"resetMonitoringInstanceId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) ResetSgwAcl() {
	_jsii_.InvokeVoid(
		m,
		"resetSgwAcl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) ResetSyslog() {
	_jsii_.InvokeVoid(
		m,
		"resetSyslog",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MariadbInstanceParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

