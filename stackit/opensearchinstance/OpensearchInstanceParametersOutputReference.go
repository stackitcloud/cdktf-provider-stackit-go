package opensearchinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/opensearchinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpensearchInstanceParametersOutputReference interface {
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
	JavaGarbageCollector() *string
	SetJavaGarbageCollector(val *string)
	JavaGarbageCollectorInput() *string
	JavaHeapspace() *float64
	SetJavaHeapspace(val *float64)
	JavaHeapspaceInput() *float64
	JavaMaxmetaspace() *float64
	SetJavaMaxmetaspace(val *float64)
	JavaMaxmetaspaceInput() *float64
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
	TlsProtocols() *string
	SetTlsProtocols(val *string)
	TlsProtocolsInput() *string
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
	ResetJavaGarbageCollector()
	ResetJavaHeapspace()
	ResetJavaMaxmetaspace()
	ResetMaxDiskThreshold()
	ResetMetricsFrequency()
	ResetMetricsPrefix()
	ResetMonitoringInstanceId()
	ResetPlugins()
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

// The jsii proxy struct for OpensearchInstanceParametersOutputReference
type jsiiProxy_OpensearchInstanceParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) EnableMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) EnableMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) Graphite() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) GraphiteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) JavaGarbageCollector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaGarbageCollector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) JavaGarbageCollectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaGarbageCollectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) JavaHeapspace() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"javaHeapspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) JavaHeapspaceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"javaHeapspaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) JavaMaxmetaspace() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"javaMaxmetaspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) JavaMaxmetaspaceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"javaMaxmetaspaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) MaxDiskThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDiskThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) MaxDiskThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDiskThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) MetricsFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) MetricsFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) MetricsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) MetricsPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) MonitoringInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) MonitoringInstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringInstanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) Plugins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"plugins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) PluginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pluginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) SgwAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sgwAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) SgwAclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sgwAclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) Syslog() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"syslog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) SyslogInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"syslogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) TlsCiphers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tlsCiphers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) TlsCiphersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tlsCiphersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) TlsProtocols() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference) TlsProtocolsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsProtocolsInput",
		&returns,
	)
	return returns
}


func NewOpensearchInstanceParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OpensearchInstanceParametersOutputReference {
	_init_.Initialize()

	if err := validateNewOpensearchInstanceParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpensearchInstanceParametersOutputReference{}

	_jsii_.Create(
		"stackit.opensearchInstance.OpensearchInstanceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOpensearchInstanceParametersOutputReference_Override(o OpensearchInstanceParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.opensearchInstance.OpensearchInstanceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference)SetEnableMonitoring(val interface{}) {
	if err := j.validateSetEnableMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMonitoring",
		val,
	)
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference)SetGraphite(val *string) {
	if err := j.validateSetGraphiteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"graphite",
		val,
	)
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference)SetJavaGarbageCollector(val *string) {
	if err := j.validateSetJavaGarbageCollectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javaGarbageCollector",
		val,
	)
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference)SetJavaHeapspace(val *float64) {
	if err := j.validateSetJavaHeapspaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javaHeapspace",
		val,
	)
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference)SetJavaMaxmetaspace(val *float64) {
	if err := j.validateSetJavaMaxmetaspaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javaMaxmetaspace",
		val,
	)
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference)SetMaxDiskThreshold(val *float64) {
	if err := j.validateSetMaxDiskThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDiskThreshold",
		val,
	)
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference)SetMetricsFrequency(val *float64) {
	if err := j.validateSetMetricsFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsFrequency",
		val,
	)
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference)SetMetricsPrefix(val *string) {
	if err := j.validateSetMetricsPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsPrefix",
		val,
	)
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference)SetMonitoringInstanceId(val *string) {
	if err := j.validateSetMonitoringInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringInstanceId",
		val,
	)
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference)SetPlugins(val *[]*string) {
	if err := j.validateSetPluginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"plugins",
		val,
	)
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference)SetSgwAcl(val *string) {
	if err := j.validateSetSgwAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sgwAcl",
		val,
	)
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference)SetSyslog(val *[]*string) {
	if err := j.validateSetSyslogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syslog",
		val,
	)
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference)SetTlsCiphers(val *[]*string) {
	if err := j.validateSetTlsCiphersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsCiphers",
		val,
	)
}

func (j *jsiiProxy_OpensearchInstanceParametersOutputReference)SetTlsProtocols(val *string) {
	if err := j.validateSetTlsProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsProtocols",
		val,
	)
}

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) ResetEnableMonitoring() {
	_jsii_.InvokeVoid(
		o,
		"resetEnableMonitoring",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) ResetGraphite() {
	_jsii_.InvokeVoid(
		o,
		"resetGraphite",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) ResetJavaGarbageCollector() {
	_jsii_.InvokeVoid(
		o,
		"resetJavaGarbageCollector",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) ResetJavaHeapspace() {
	_jsii_.InvokeVoid(
		o,
		"resetJavaHeapspace",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) ResetJavaMaxmetaspace() {
	_jsii_.InvokeVoid(
		o,
		"resetJavaMaxmetaspace",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) ResetMaxDiskThreshold() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxDiskThreshold",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) ResetMetricsFrequency() {
	_jsii_.InvokeVoid(
		o,
		"resetMetricsFrequency",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) ResetMetricsPrefix() {
	_jsii_.InvokeVoid(
		o,
		"resetMetricsPrefix",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) ResetMonitoringInstanceId() {
	_jsii_.InvokeVoid(
		o,
		"resetMonitoringInstanceId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) ResetPlugins() {
	_jsii_.InvokeVoid(
		o,
		"resetPlugins",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) ResetSgwAcl() {
	_jsii_.InvokeVoid(
		o,
		"resetSgwAcl",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) ResetSyslog() {
	_jsii_.InvokeVoid(
		o,
		"resetSyslog",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) ResetTlsCiphers() {
	_jsii_.InvokeVoid(
		o,
		"resetTlsCiphers",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) ResetTlsProtocols() {
	_jsii_.InvokeVoid(
		o,
		"resetTlsProtocols",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchInstanceParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

