package datastackitvalkeyinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/datastackitvalkeyinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitValkeyInstanceParametersOutputReference interface {
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
	EnableMonitoring() cdktf.IResolvable
	FailoverTimeout() *float64
	// Experimental.
	Fqn() *string
	Graphite() *string
	InternalValue() *DataStackitValkeyInstanceParameters
	SetInternalValue(val *DataStackitValkeyInstanceParameters)
	LazyfreeLazyEviction() *string
	LazyfreeLazyExpire() *string
	LuaTimeLimit() *float64
	Maxclients() *float64
	MaxDiskThreshold() *float64
	MaxmemoryPolicy() *string
	MaxmemorySamples() *float64
	MetricsFrequency() *float64
	MetricsPrefix() *string
	MinReplicasMaxLag() *float64
	MinReplicasToWrite() *float64
	MonitoringInstanceId() *string
	NotifyKeyspaceEvents() *string
	ReplBacklogSize() *string
	SgwAcl() *string
	Snapshot() *string
	Syslog() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataStackitValkeyInstanceParametersOutputReference
type jsiiProxy_DataStackitValkeyInstanceParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) DownAfterMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"downAfterMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) EnableMonitoring() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) FailoverTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failoverTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) Graphite() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) InternalValue() *DataStackitValkeyInstanceParameters {
	var returns *DataStackitValkeyInstanceParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) LazyfreeLazyEviction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lazyfreeLazyEviction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) LazyfreeLazyExpire() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lazyfreeLazyExpire",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) LuaTimeLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"luaTimeLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) Maxclients() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxclients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) MaxDiskThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDiskThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) MaxmemoryPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxmemoryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) MaxmemorySamples() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxmemorySamples",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) MetricsFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) MetricsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) MinReplicasMaxLag() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicasMaxLag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) MinReplicasToWrite() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicasToWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) MonitoringInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) NotifyKeyspaceEvents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyKeyspaceEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) ReplBacklogSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replBacklogSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) SgwAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sgwAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) Snapshot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) Syslog() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"syslog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataStackitValkeyInstanceParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataStackitValkeyInstanceParametersOutputReference {
	_init_.Initialize()

	if err := validateNewDataStackitValkeyInstanceParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataStackitValkeyInstanceParametersOutputReference{}

	_jsii_.Create(
		"stackit.dataStackitValkeyInstance.DataStackitValkeyInstanceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataStackitValkeyInstanceParametersOutputReference_Override(d DataStackitValkeyInstanceParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.dataStackitValkeyInstance.DataStackitValkeyInstanceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference)SetInternalValue(val *DataStackitValkeyInstanceParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitValkeyInstanceParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

