package redisinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/redisinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedisInstanceParametersOutputReference interface {
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
	DownAfterMilliseconds() *float64
	SetDownAfterMilliseconds(val *float64)
	DownAfterMillisecondsInput() *float64
	EnableMonitoring() interface{}
	SetEnableMonitoring(val interface{})
	EnableMonitoringInput() interface{}
	FailoverTimeout() *float64
	SetFailoverTimeout(val *float64)
	FailoverTimeoutInput() *float64
	// Experimental.
	Fqn() *string
	Graphite() *string
	SetGraphite(val *string)
	GraphiteInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LazyfreeLazyEviction() *string
	SetLazyfreeLazyEviction(val *string)
	LazyfreeLazyEvictionInput() *string
	LazyfreeLazyExpire() *string
	SetLazyfreeLazyExpire(val *string)
	LazyfreeLazyExpireInput() *string
	LuaTimeLimit() *float64
	SetLuaTimeLimit(val *float64)
	LuaTimeLimitInput() *float64
	Maxclients() *float64
	SetMaxclients(val *float64)
	MaxclientsInput() *float64
	MaxDiskThreshold() *float64
	SetMaxDiskThreshold(val *float64)
	MaxDiskThresholdInput() *float64
	MaxmemoryPolicy() *string
	SetMaxmemoryPolicy(val *string)
	MaxmemoryPolicyInput() *string
	MaxmemorySamples() *float64
	SetMaxmemorySamples(val *float64)
	MaxmemorySamplesInput() *float64
	MetricsFrequency() *float64
	SetMetricsFrequency(val *float64)
	MetricsFrequencyInput() *float64
	MetricsPrefix() *string
	SetMetricsPrefix(val *string)
	MetricsPrefixInput() *string
	MinReplicasMaxLag() *float64
	SetMinReplicasMaxLag(val *float64)
	MinReplicasMaxLagInput() *float64
	MonitoringInstanceId() *string
	SetMonitoringInstanceId(val *string)
	MonitoringInstanceIdInput() *string
	NotifyKeyspaceEvents() *string
	SetNotifyKeyspaceEvents(val *string)
	NotifyKeyspaceEventsInput() *string
	SgwAcl() *string
	SetSgwAcl(val *string)
	SgwAclInput() *string
	Snapshot() *string
	SetSnapshot(val *string)
	SnapshotInput() *string
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
	TlsCiphersuites() *string
	SetTlsCiphersuites(val *string)
	TlsCiphersuitesInput() *string
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
	ResetDownAfterMilliseconds()
	ResetEnableMonitoring()
	ResetFailoverTimeout()
	ResetGraphite()
	ResetLazyfreeLazyEviction()
	ResetLazyfreeLazyExpire()
	ResetLuaTimeLimit()
	ResetMaxclients()
	ResetMaxDiskThreshold()
	ResetMaxmemoryPolicy()
	ResetMaxmemorySamples()
	ResetMetricsFrequency()
	ResetMetricsPrefix()
	ResetMinReplicasMaxLag()
	ResetMonitoringInstanceId()
	ResetNotifyKeyspaceEvents()
	ResetSgwAcl()
	ResetSnapshot()
	ResetSyslog()
	ResetTlsCiphers()
	ResetTlsCiphersuites()
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

// The jsii proxy struct for RedisInstanceParametersOutputReference
type jsiiProxy_RedisInstanceParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) DownAfterMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"downAfterMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) DownAfterMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"downAfterMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) EnableMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) EnableMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) FailoverTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failoverTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) FailoverTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failoverTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) Graphite() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) GraphiteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) LazyfreeLazyEviction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lazyfreeLazyEviction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) LazyfreeLazyEvictionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lazyfreeLazyEvictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) LazyfreeLazyExpire() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lazyfreeLazyExpire",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) LazyfreeLazyExpireInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lazyfreeLazyExpireInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) LuaTimeLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"luaTimeLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) LuaTimeLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"luaTimeLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) Maxclients() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxclients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) MaxclientsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxclientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) MaxDiskThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDiskThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) MaxDiskThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDiskThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) MaxmemoryPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxmemoryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) MaxmemoryPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxmemoryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) MaxmemorySamples() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxmemorySamples",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) MaxmemorySamplesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxmemorySamplesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) MetricsFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) MetricsFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) MetricsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) MetricsPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) MinReplicasMaxLag() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicasMaxLag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) MinReplicasMaxLagInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicasMaxLagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) MonitoringInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) MonitoringInstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringInstanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) NotifyKeyspaceEvents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyKeyspaceEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) NotifyKeyspaceEventsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyKeyspaceEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) SgwAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sgwAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) SgwAclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sgwAclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) Snapshot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) SnapshotInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) Syslog() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"syslog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) SyslogInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"syslogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) TlsCiphers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tlsCiphers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) TlsCiphersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tlsCiphersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) TlsCiphersuites() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCiphersuites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) TlsCiphersuitesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCiphersuitesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) TlsProtocols() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference) TlsProtocolsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsProtocolsInput",
		&returns,
	)
	return returns
}


func NewRedisInstanceParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RedisInstanceParametersOutputReference {
	_init_.Initialize()

	if err := validateNewRedisInstanceParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedisInstanceParametersOutputReference{}

	_jsii_.Create(
		"stackit.redisInstance.RedisInstanceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRedisInstanceParametersOutputReference_Override(r RedisInstanceParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.redisInstance.RedisInstanceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetDownAfterMilliseconds(val *float64) {
	if err := j.validateSetDownAfterMillisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"downAfterMilliseconds",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetEnableMonitoring(val interface{}) {
	if err := j.validateSetEnableMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMonitoring",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetFailoverTimeout(val *float64) {
	if err := j.validateSetFailoverTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failoverTimeout",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetGraphite(val *string) {
	if err := j.validateSetGraphiteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"graphite",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetLazyfreeLazyEviction(val *string) {
	if err := j.validateSetLazyfreeLazyEvictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lazyfreeLazyEviction",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetLazyfreeLazyExpire(val *string) {
	if err := j.validateSetLazyfreeLazyExpireParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lazyfreeLazyExpire",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetLuaTimeLimit(val *float64) {
	if err := j.validateSetLuaTimeLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"luaTimeLimit",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetMaxclients(val *float64) {
	if err := j.validateSetMaxclientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxclients",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetMaxDiskThreshold(val *float64) {
	if err := j.validateSetMaxDiskThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDiskThreshold",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetMaxmemoryPolicy(val *string) {
	if err := j.validateSetMaxmemoryPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxmemoryPolicy",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetMaxmemorySamples(val *float64) {
	if err := j.validateSetMaxmemorySamplesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxmemorySamples",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetMetricsFrequency(val *float64) {
	if err := j.validateSetMetricsFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsFrequency",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetMetricsPrefix(val *string) {
	if err := j.validateSetMetricsPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsPrefix",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetMinReplicasMaxLag(val *float64) {
	if err := j.validateSetMinReplicasMaxLagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minReplicasMaxLag",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetMonitoringInstanceId(val *string) {
	if err := j.validateSetMonitoringInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringInstanceId",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetNotifyKeyspaceEvents(val *string) {
	if err := j.validateSetNotifyKeyspaceEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyKeyspaceEvents",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetSgwAcl(val *string) {
	if err := j.validateSetSgwAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sgwAcl",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetSnapshot(val *string) {
	if err := j.validateSetSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshot",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetSyslog(val *[]*string) {
	if err := j.validateSetSyslogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syslog",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetTlsCiphers(val *[]*string) {
	if err := j.validateSetTlsCiphersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsCiphers",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetTlsCiphersuites(val *string) {
	if err := j.validateSetTlsCiphersuitesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsCiphersuites",
		val,
	)
}

func (j *jsiiProxy_RedisInstanceParametersOutputReference)SetTlsProtocols(val *string) {
	if err := j.validateSetTlsProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsProtocols",
		val,
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RedisInstanceParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RedisInstanceParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RedisInstanceParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RedisInstanceParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RedisInstanceParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RedisInstanceParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RedisInstanceParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RedisInstanceParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RedisInstanceParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetDownAfterMilliseconds() {
	_jsii_.InvokeVoid(
		r,
		"resetDownAfterMilliseconds",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetEnableMonitoring() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableMonitoring",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetFailoverTimeout() {
	_jsii_.InvokeVoid(
		r,
		"resetFailoverTimeout",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetGraphite() {
	_jsii_.InvokeVoid(
		r,
		"resetGraphite",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetLazyfreeLazyEviction() {
	_jsii_.InvokeVoid(
		r,
		"resetLazyfreeLazyEviction",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetLazyfreeLazyExpire() {
	_jsii_.InvokeVoid(
		r,
		"resetLazyfreeLazyExpire",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetLuaTimeLimit() {
	_jsii_.InvokeVoid(
		r,
		"resetLuaTimeLimit",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetMaxclients() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxclients",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetMaxDiskThreshold() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxDiskThreshold",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetMaxmemoryPolicy() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxmemoryPolicy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetMaxmemorySamples() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxmemorySamples",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetMetricsFrequency() {
	_jsii_.InvokeVoid(
		r,
		"resetMetricsFrequency",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetMetricsPrefix() {
	_jsii_.InvokeVoid(
		r,
		"resetMetricsPrefix",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetMinReplicasMaxLag() {
	_jsii_.InvokeVoid(
		r,
		"resetMinReplicasMaxLag",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetMonitoringInstanceId() {
	_jsii_.InvokeVoid(
		r,
		"resetMonitoringInstanceId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetNotifyKeyspaceEvents() {
	_jsii_.InvokeVoid(
		r,
		"resetNotifyKeyspaceEvents",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetSgwAcl() {
	_jsii_.InvokeVoid(
		r,
		"resetSgwAcl",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetSnapshot() {
	_jsii_.InvokeVoid(
		r,
		"resetSnapshot",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetSyslog() {
	_jsii_.InvokeVoid(
		r,
		"resetSyslog",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetTlsCiphers() {
	_jsii_.InvokeVoid(
		r,
		"resetTlsCiphers",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetTlsCiphersuites() {
	_jsii_.InvokeVoid(
		r,
		"resetTlsCiphersuites",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ResetTlsProtocols() {
	_jsii_.InvokeVoid(
		r,
		"resetTlsProtocols",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisInstanceParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_RedisInstanceParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

