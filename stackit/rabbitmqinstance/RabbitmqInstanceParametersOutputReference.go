package rabbitmqinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/rabbitmqinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RabbitmqInstanceParametersOutputReference interface {
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
	ConsumerTimeout() *float64
	SetConsumerTimeout(val *float64)
	ConsumerTimeoutInput() *float64
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
	Plugins() *[]*string
	SetPlugins(val *[]*string)
	PluginsInput() *[]*string
	Roles() *[]*string
	SetRoles(val *[]*string)
	RolesInput() *[]*string
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
	TlsCiphers() *[]*string
	SetTlsCiphers(val *[]*string)
	TlsCiphersInput() *[]*string
	TlsProtocols() *[]*string
	SetTlsProtocols(val *[]*string)
	TlsProtocolsInput() *[]*string
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
	ResetConsumerTimeout()
	ResetEnableMonitoring()
	ResetGraphite()
	ResetMaxDiskThreshold()
	ResetMetricsFrequency()
	ResetMetricsPrefix()
	ResetMonitoringInstanceId()
	ResetPlugins()
	ResetRoles()
	ResetSgwAcl()
	ResetSyslog()
	ResetTlsCiphers()
	ResetTlsProtocols()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RabbitmqInstanceParametersOutputReference
type jsiiProxy_RabbitmqInstanceParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) ConsumerTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consumerTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) ConsumerTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consumerTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) EnableMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) EnableMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) Graphite() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) GraphiteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) MaxDiskThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDiskThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) MaxDiskThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDiskThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) MetricsFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) MetricsFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) MetricsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) MetricsPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) MonitoringInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) MonitoringInstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringInstanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) Plugins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"plugins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) PluginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pluginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) Roles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"roles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) RolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) SgwAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sgwAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) SgwAclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sgwAclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) Syslog() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"syslog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) SyslogInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"syslogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) TlsCiphers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tlsCiphers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) TlsCiphersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tlsCiphersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) TlsProtocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tlsProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference) TlsProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tlsProtocolsInput",
		&returns,
	)
	return returns
}


func NewRabbitmqInstanceParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RabbitmqInstanceParametersOutputReference {
	_init_.Initialize()

	if err := validateNewRabbitmqInstanceParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RabbitmqInstanceParametersOutputReference{}

	_jsii_.Create(
		"stackit.rabbitmqInstance.RabbitmqInstanceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRabbitmqInstanceParametersOutputReference_Override(r RabbitmqInstanceParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.rabbitmqInstance.RabbitmqInstanceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference)SetConsumerTimeout(val *float64) {
	if err := j.validateSetConsumerTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerTimeout",
		val,
	)
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference)SetEnableMonitoring(val interface{}) {
	if err := j.validateSetEnableMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMonitoring",
		val,
	)
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference)SetGraphite(val *string) {
	if err := j.validateSetGraphiteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"graphite",
		val,
	)
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference)SetMaxDiskThreshold(val *float64) {
	if err := j.validateSetMaxDiskThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDiskThreshold",
		val,
	)
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference)SetMetricsFrequency(val *float64) {
	if err := j.validateSetMetricsFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsFrequency",
		val,
	)
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference)SetMetricsPrefix(val *string) {
	if err := j.validateSetMetricsPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsPrefix",
		val,
	)
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference)SetMonitoringInstanceId(val *string) {
	if err := j.validateSetMonitoringInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringInstanceId",
		val,
	)
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference)SetPlugins(val *[]*string) {
	if err := j.validateSetPluginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"plugins",
		val,
	)
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference)SetRoles(val *[]*string) {
	if err := j.validateSetRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roles",
		val,
	)
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference)SetSgwAcl(val *string) {
	if err := j.validateSetSgwAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sgwAcl",
		val,
	)
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference)SetSyslog(val *[]*string) {
	if err := j.validateSetSyslogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syslog",
		val,
	)
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference)SetTlsCiphers(val *[]*string) {
	if err := j.validateSetTlsCiphersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsCiphers",
		val,
	)
}

func (j *jsiiProxy_RabbitmqInstanceParametersOutputReference)SetTlsProtocols(val *[]*string) {
	if err := j.validateSetTlsProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsProtocols",
		val,
	)
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) ResetConsumerTimeout() {
	_jsii_.InvokeVoid(
		r,
		"resetConsumerTimeout",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) ResetEnableMonitoring() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableMonitoring",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) ResetGraphite() {
	_jsii_.InvokeVoid(
		r,
		"resetGraphite",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) ResetMaxDiskThreshold() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxDiskThreshold",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) ResetMetricsFrequency() {
	_jsii_.InvokeVoid(
		r,
		"resetMetricsFrequency",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) ResetMetricsPrefix() {
	_jsii_.InvokeVoid(
		r,
		"resetMetricsPrefix",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) ResetMonitoringInstanceId() {
	_jsii_.InvokeVoid(
		r,
		"resetMonitoringInstanceId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) ResetPlugins() {
	_jsii_.InvokeVoid(
		r,
		"resetPlugins",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) ResetRoles() {
	_jsii_.InvokeVoid(
		r,
		"resetRoles",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) ResetSgwAcl() {
	_jsii_.InvokeVoid(
		r,
		"resetSgwAcl",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) ResetSyslog() {
	_jsii_.InvokeVoid(
		r,
		"resetSyslog",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) ResetTlsCiphers() {
	_jsii_.InvokeVoid(
		r,
		"resetTlsCiphers",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) ResetTlsProtocols() {
	_jsii_.InvokeVoid(
		r,
		"resetTlsProtocols",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RabbitmqInstanceParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

