package valkeyinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/valkeyinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ValkeyInstanceParametersOutputReference interface {
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
	MinReplicasToWrite() *float64
	SetMinReplicasToWrite(val *float64)
	MinReplicasToWriteInput() *float64
	MonitoringInstanceId() *string
	SetMonitoringInstanceId(val *string)
	MonitoringInstanceIdInput() *string
	NotifyKeyspaceEvents() *string
	SetNotifyKeyspaceEvents(val *string)
	NotifyKeyspaceEventsInput() *string
	ReplBacklogSize() *string
	SetReplBacklogSize(val *string)
	ReplBacklogSizeInput() *string
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
	ResetMinReplicasToWrite()
	ResetMonitoringInstanceId()
	ResetNotifyKeyspaceEvents()
	ResetReplBacklogSize()
	ResetSgwAcl()
	ResetSnapshot()
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

// The jsii proxy struct for ValkeyInstanceParametersOutputReference
type jsiiProxy_ValkeyInstanceParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) DownAfterMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"downAfterMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) DownAfterMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"downAfterMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) EnableMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) EnableMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) FailoverTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failoverTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) FailoverTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failoverTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) Graphite() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) GraphiteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) LazyfreeLazyEviction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lazyfreeLazyEviction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) LazyfreeLazyEvictionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lazyfreeLazyEvictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) LazyfreeLazyExpire() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lazyfreeLazyExpire",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) LazyfreeLazyExpireInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lazyfreeLazyExpireInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) LuaTimeLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"luaTimeLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) LuaTimeLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"luaTimeLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) Maxclients() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxclients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) MaxclientsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxclientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) MaxDiskThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDiskThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) MaxDiskThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDiskThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) MaxmemoryPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxmemoryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) MaxmemoryPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxmemoryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) MaxmemorySamples() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxmemorySamples",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) MaxmemorySamplesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxmemorySamplesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) MetricsFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) MetricsFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) MetricsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) MetricsPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) MinReplicasMaxLag() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicasMaxLag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) MinReplicasMaxLagInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicasMaxLagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) MinReplicasToWrite() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicasToWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) MinReplicasToWriteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicasToWriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) MonitoringInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) MonitoringInstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringInstanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) NotifyKeyspaceEvents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyKeyspaceEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) NotifyKeyspaceEventsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyKeyspaceEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) ReplBacklogSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replBacklogSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) ReplBacklogSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replBacklogSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) SgwAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sgwAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) SgwAclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sgwAclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) Snapshot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) SnapshotInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) Syslog() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"syslog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) SyslogInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"syslogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewValkeyInstanceParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ValkeyInstanceParametersOutputReference {
	_init_.Initialize()

	if err := validateNewValkeyInstanceParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ValkeyInstanceParametersOutputReference{}

	_jsii_.Create(
		"stackit.valkeyInstance.ValkeyInstanceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewValkeyInstanceParametersOutputReference_Override(v ValkeyInstanceParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.valkeyInstance.ValkeyInstanceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetDownAfterMilliseconds(val *float64) {
	if err := j.validateSetDownAfterMillisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"downAfterMilliseconds",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetEnableMonitoring(val interface{}) {
	if err := j.validateSetEnableMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMonitoring",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetFailoverTimeout(val *float64) {
	if err := j.validateSetFailoverTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failoverTimeout",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetGraphite(val *string) {
	if err := j.validateSetGraphiteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"graphite",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetLazyfreeLazyEviction(val *string) {
	if err := j.validateSetLazyfreeLazyEvictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lazyfreeLazyEviction",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetLazyfreeLazyExpire(val *string) {
	if err := j.validateSetLazyfreeLazyExpireParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lazyfreeLazyExpire",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetLuaTimeLimit(val *float64) {
	if err := j.validateSetLuaTimeLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"luaTimeLimit",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetMaxclients(val *float64) {
	if err := j.validateSetMaxclientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxclients",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetMaxDiskThreshold(val *float64) {
	if err := j.validateSetMaxDiskThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDiskThreshold",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetMaxmemoryPolicy(val *string) {
	if err := j.validateSetMaxmemoryPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxmemoryPolicy",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetMaxmemorySamples(val *float64) {
	if err := j.validateSetMaxmemorySamplesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxmemorySamples",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetMetricsFrequency(val *float64) {
	if err := j.validateSetMetricsFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsFrequency",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetMetricsPrefix(val *string) {
	if err := j.validateSetMetricsPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsPrefix",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetMinReplicasMaxLag(val *float64) {
	if err := j.validateSetMinReplicasMaxLagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minReplicasMaxLag",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetMinReplicasToWrite(val *float64) {
	if err := j.validateSetMinReplicasToWriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minReplicasToWrite",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetMonitoringInstanceId(val *string) {
	if err := j.validateSetMonitoringInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringInstanceId",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetNotifyKeyspaceEvents(val *string) {
	if err := j.validateSetNotifyKeyspaceEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyKeyspaceEvents",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetReplBacklogSize(val *string) {
	if err := j.validateSetReplBacklogSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replBacklogSize",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetSgwAcl(val *string) {
	if err := j.validateSetSgwAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sgwAcl",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetSnapshot(val *string) {
	if err := j.validateSetSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshot",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetSyslog(val *[]*string) {
	if err := j.validateSetSyslogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syslog",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ValkeyInstanceParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetDownAfterMilliseconds() {
	_jsii_.InvokeVoid(
		v,
		"resetDownAfterMilliseconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetEnableMonitoring() {
	_jsii_.InvokeVoid(
		v,
		"resetEnableMonitoring",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetFailoverTimeout() {
	_jsii_.InvokeVoid(
		v,
		"resetFailoverTimeout",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetGraphite() {
	_jsii_.InvokeVoid(
		v,
		"resetGraphite",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetLazyfreeLazyEviction() {
	_jsii_.InvokeVoid(
		v,
		"resetLazyfreeLazyEviction",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetLazyfreeLazyExpire() {
	_jsii_.InvokeVoid(
		v,
		"resetLazyfreeLazyExpire",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetLuaTimeLimit() {
	_jsii_.InvokeVoid(
		v,
		"resetLuaTimeLimit",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetMaxclients() {
	_jsii_.InvokeVoid(
		v,
		"resetMaxclients",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetMaxDiskThreshold() {
	_jsii_.InvokeVoid(
		v,
		"resetMaxDiskThreshold",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetMaxmemoryPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetMaxmemoryPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetMaxmemorySamples() {
	_jsii_.InvokeVoid(
		v,
		"resetMaxmemorySamples",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetMetricsFrequency() {
	_jsii_.InvokeVoid(
		v,
		"resetMetricsFrequency",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetMetricsPrefix() {
	_jsii_.InvokeVoid(
		v,
		"resetMetricsPrefix",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetMinReplicasMaxLag() {
	_jsii_.InvokeVoid(
		v,
		"resetMinReplicasMaxLag",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetMinReplicasToWrite() {
	_jsii_.InvokeVoid(
		v,
		"resetMinReplicasToWrite",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetMonitoringInstanceId() {
	_jsii_.InvokeVoid(
		v,
		"resetMonitoringInstanceId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetNotifyKeyspaceEvents() {
	_jsii_.InvokeVoid(
		v,
		"resetNotifyKeyspaceEvents",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetReplBacklogSize() {
	_jsii_.InvokeVoid(
		v,
		"resetReplBacklogSize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetSgwAcl() {
	_jsii_.InvokeVoid(
		v,
		"resetSgwAcl",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetSnapshot() {
	_jsii_.InvokeVoid(
		v,
		"resetSnapshot",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ResetSyslog() {
	_jsii_.InvokeVoid(
		v,
		"resetSyslog",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValkeyInstanceParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

