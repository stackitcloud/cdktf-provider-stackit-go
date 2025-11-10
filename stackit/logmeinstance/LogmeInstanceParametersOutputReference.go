package logmeinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/logmeinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogmeInstanceParametersOutputReference interface {
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
	FluentdTcp() *float64
	SetFluentdTcp(val *float64)
	FluentdTcpInput() *float64
	FluentdTls() *float64
	SetFluentdTls(val *float64)
	FluentdTlsCiphers() *string
	SetFluentdTlsCiphers(val *string)
	FluentdTlsCiphersInput() *string
	FluentdTlsInput() *float64
	FluentdTlsMaxVersion() *string
	SetFluentdTlsMaxVersion(val *string)
	FluentdTlsMaxVersionInput() *string
	FluentdTlsMinVersion() *string
	SetFluentdTlsMinVersion(val *string)
	FluentdTlsMinVersionInput() *string
	FluentdTlsVersion() *string
	SetFluentdTlsVersion(val *string)
	FluentdTlsVersionInput() *string
	FluentdUdp() *float64
	SetFluentdUdp(val *float64)
	FluentdUdpInput() *float64
	// Experimental.
	Fqn() *string
	Graphite() *string
	SetGraphite(val *string)
	GraphiteInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsmDeletionAfter() *string
	SetIsmDeletionAfter(val *string)
	IsmDeletionAfterInput() *string
	IsmJitter() *float64
	SetIsmJitter(val *float64)
	IsmJitterInput() *float64
	IsmJobInterval() *float64
	SetIsmJobInterval(val *float64)
	IsmJobIntervalInput() *float64
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
	OpensearchTlsCiphers() *[]*string
	SetOpensearchTlsCiphers(val *[]*string)
	OpensearchTlsCiphersInput() *[]*string
	OpensearchTlsProtocols() *[]*string
	SetOpensearchTlsProtocols(val *[]*string)
	OpensearchTlsProtocolsInput() *[]*string
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
	ResetFluentdTcp()
	ResetFluentdTls()
	ResetFluentdTlsCiphers()
	ResetFluentdTlsMaxVersion()
	ResetFluentdTlsMinVersion()
	ResetFluentdTlsVersion()
	ResetFluentdUdp()
	ResetGraphite()
	ResetIsmDeletionAfter()
	ResetIsmJitter()
	ResetIsmJobInterval()
	ResetJavaHeapspace()
	ResetJavaMaxmetaspace()
	ResetMaxDiskThreshold()
	ResetMetricsFrequency()
	ResetMetricsPrefix()
	ResetMonitoringInstanceId()
	ResetOpensearchTlsCiphers()
	ResetOpensearchTlsProtocols()
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

// The jsii proxy struct for LogmeInstanceParametersOutputReference
type jsiiProxy_LogmeInstanceParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) EnableMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) EnableMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) FluentdTcp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fluentdTcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) FluentdTcpInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fluentdTcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) FluentdTls() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fluentdTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) FluentdTlsCiphers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fluentdTlsCiphers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) FluentdTlsCiphersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fluentdTlsCiphersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) FluentdTlsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fluentdTlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) FluentdTlsMaxVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fluentdTlsMaxVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) FluentdTlsMaxVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fluentdTlsMaxVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) FluentdTlsMinVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fluentdTlsMinVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) FluentdTlsMinVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fluentdTlsMinVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) FluentdTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fluentdTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) FluentdTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fluentdTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) FluentdUdp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fluentdUdp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) FluentdUdpInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fluentdUdpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) Graphite() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) GraphiteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) IsmDeletionAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ismDeletionAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) IsmDeletionAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ismDeletionAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) IsmJitter() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismJitter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) IsmJitterInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismJitterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) IsmJobInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismJobInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) IsmJobIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismJobIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) JavaHeapspace() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"javaHeapspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) JavaHeapspaceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"javaHeapspaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) JavaMaxmetaspace() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"javaMaxmetaspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) JavaMaxmetaspaceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"javaMaxmetaspaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) MaxDiskThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDiskThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) MaxDiskThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDiskThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) MetricsFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) MetricsFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) MetricsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) MetricsPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) MonitoringInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) MonitoringInstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringInstanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) OpensearchTlsCiphers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"opensearchTlsCiphers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) OpensearchTlsCiphersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"opensearchTlsCiphersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) OpensearchTlsProtocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"opensearchTlsProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) OpensearchTlsProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"opensearchTlsProtocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) SgwAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sgwAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) SgwAclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sgwAclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) Syslog() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"syslog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) SyslogInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"syslogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLogmeInstanceParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LogmeInstanceParametersOutputReference {
	_init_.Initialize()

	if err := validateNewLogmeInstanceParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogmeInstanceParametersOutputReference{}

	_jsii_.Create(
		"stackit.logmeInstance.LogmeInstanceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLogmeInstanceParametersOutputReference_Override(l LogmeInstanceParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.logmeInstance.LogmeInstanceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetEnableMonitoring(val interface{}) {
	if err := j.validateSetEnableMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMonitoring",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetFluentdTcp(val *float64) {
	if err := j.validateSetFluentdTcpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fluentdTcp",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetFluentdTls(val *float64) {
	if err := j.validateSetFluentdTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fluentdTls",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetFluentdTlsCiphers(val *string) {
	if err := j.validateSetFluentdTlsCiphersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fluentdTlsCiphers",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetFluentdTlsMaxVersion(val *string) {
	if err := j.validateSetFluentdTlsMaxVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fluentdTlsMaxVersion",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetFluentdTlsMinVersion(val *string) {
	if err := j.validateSetFluentdTlsMinVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fluentdTlsMinVersion",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetFluentdTlsVersion(val *string) {
	if err := j.validateSetFluentdTlsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fluentdTlsVersion",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetFluentdUdp(val *float64) {
	if err := j.validateSetFluentdUdpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fluentdUdp",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetGraphite(val *string) {
	if err := j.validateSetGraphiteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"graphite",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetIsmDeletionAfter(val *string) {
	if err := j.validateSetIsmDeletionAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ismDeletionAfter",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetIsmJitter(val *float64) {
	if err := j.validateSetIsmJitterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ismJitter",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetIsmJobInterval(val *float64) {
	if err := j.validateSetIsmJobIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ismJobInterval",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetJavaHeapspace(val *float64) {
	if err := j.validateSetJavaHeapspaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javaHeapspace",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetJavaMaxmetaspace(val *float64) {
	if err := j.validateSetJavaMaxmetaspaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javaMaxmetaspace",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetMaxDiskThreshold(val *float64) {
	if err := j.validateSetMaxDiskThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDiskThreshold",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetMetricsFrequency(val *float64) {
	if err := j.validateSetMetricsFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsFrequency",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetMetricsPrefix(val *string) {
	if err := j.validateSetMetricsPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsPrefix",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetMonitoringInstanceId(val *string) {
	if err := j.validateSetMonitoringInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringInstanceId",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetOpensearchTlsCiphers(val *[]*string) {
	if err := j.validateSetOpensearchTlsCiphersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opensearchTlsCiphers",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetOpensearchTlsProtocols(val *[]*string) {
	if err := j.validateSetOpensearchTlsProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opensearchTlsProtocols",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetSgwAcl(val *string) {
	if err := j.validateSetSgwAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sgwAcl",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetSyslog(val *[]*string) {
	if err := j.validateSetSyslogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syslog",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LogmeInstanceParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetEnableMonitoring() {
	_jsii_.InvokeVoid(
		l,
		"resetEnableMonitoring",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetFluentdTcp() {
	_jsii_.InvokeVoid(
		l,
		"resetFluentdTcp",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetFluentdTls() {
	_jsii_.InvokeVoid(
		l,
		"resetFluentdTls",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetFluentdTlsCiphers() {
	_jsii_.InvokeVoid(
		l,
		"resetFluentdTlsCiphers",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetFluentdTlsMaxVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetFluentdTlsMaxVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetFluentdTlsMinVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetFluentdTlsMinVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetFluentdTlsVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetFluentdTlsVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetFluentdUdp() {
	_jsii_.InvokeVoid(
		l,
		"resetFluentdUdp",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetGraphite() {
	_jsii_.InvokeVoid(
		l,
		"resetGraphite",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetIsmDeletionAfter() {
	_jsii_.InvokeVoid(
		l,
		"resetIsmDeletionAfter",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetIsmJitter() {
	_jsii_.InvokeVoid(
		l,
		"resetIsmJitter",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetIsmJobInterval() {
	_jsii_.InvokeVoid(
		l,
		"resetIsmJobInterval",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetJavaHeapspace() {
	_jsii_.InvokeVoid(
		l,
		"resetJavaHeapspace",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetJavaMaxmetaspace() {
	_jsii_.InvokeVoid(
		l,
		"resetJavaMaxmetaspace",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetMaxDiskThreshold() {
	_jsii_.InvokeVoid(
		l,
		"resetMaxDiskThreshold",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetMetricsFrequency() {
	_jsii_.InvokeVoid(
		l,
		"resetMetricsFrequency",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetMetricsPrefix() {
	_jsii_.InvokeVoid(
		l,
		"resetMetricsPrefix",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetMonitoringInstanceId() {
	_jsii_.InvokeVoid(
		l,
		"resetMonitoringInstanceId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetOpensearchTlsCiphers() {
	_jsii_.InvokeVoid(
		l,
		"resetOpensearchTlsCiphers",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetOpensearchTlsProtocols() {
	_jsii_.InvokeVoid(
		l,
		"resetOpensearchTlsProtocols",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetSgwAcl() {
	_jsii_.InvokeVoid(
		l,
		"resetSgwAcl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ResetSyslog() {
	_jsii_.InvokeVoid(
		l,
		"resetSyslog",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogmeInstanceParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

