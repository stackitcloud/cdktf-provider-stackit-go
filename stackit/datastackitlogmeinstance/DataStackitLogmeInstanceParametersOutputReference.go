package datastackitlogmeinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/datastackitlogmeinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitLogmeInstanceParametersOutputReference interface {
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
	EnableMonitoring() cdktf.IResolvable
	FluentdTcp() *float64
	FluentdTls() *float64
	FluentdTlsCiphers() *string
	FluentdTlsMaxVersion() *string
	FluentdTlsMinVersion() *string
	FluentdTlsVersion() *string
	FluentdUdp() *float64
	// Experimental.
	Fqn() *string
	Graphite() *string
	InternalValue() *DataStackitLogmeInstanceParameters
	SetInternalValue(val *DataStackitLogmeInstanceParameters)
	IsmDeletionAfter() *string
	IsmJitter() *float64
	IsmJobInterval() *float64
	JavaHeapspace() *float64
	JavaMaxmetaspace() *float64
	MaxDiskThreshold() *float64
	MetricsFrequency() *float64
	MetricsPrefix() *string
	MonitoringInstanceId() *string
	OpensearchTlsCiphers() *[]*string
	OpensearchTlsProtocols() *[]*string
	SgwAcl() *string
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

// The jsii proxy struct for DataStackitLogmeInstanceParametersOutputReference
type jsiiProxy_DataStackitLogmeInstanceParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) EnableMonitoring() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) FluentdTcp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fluentdTcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) FluentdTls() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fluentdTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) FluentdTlsCiphers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fluentdTlsCiphers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) FluentdTlsMaxVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fluentdTlsMaxVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) FluentdTlsMinVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fluentdTlsMinVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) FluentdTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fluentdTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) FluentdUdp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fluentdUdp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) Graphite() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) InternalValue() *DataStackitLogmeInstanceParameters {
	var returns *DataStackitLogmeInstanceParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) IsmDeletionAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ismDeletionAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) IsmJitter() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismJitter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) IsmJobInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismJobInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) JavaHeapspace() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"javaHeapspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) JavaMaxmetaspace() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"javaMaxmetaspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) MaxDiskThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDiskThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) MetricsFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) MetricsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) MonitoringInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) OpensearchTlsCiphers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"opensearchTlsCiphers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) OpensearchTlsProtocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"opensearchTlsProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) SgwAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sgwAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) Syslog() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"syslog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataStackitLogmeInstanceParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataStackitLogmeInstanceParametersOutputReference {
	_init_.Initialize()

	if err := validateNewDataStackitLogmeInstanceParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataStackitLogmeInstanceParametersOutputReference{}

	_jsii_.Create(
		"stackit.dataStackitLogmeInstance.DataStackitLogmeInstanceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataStackitLogmeInstanceParametersOutputReference_Override(d DataStackitLogmeInstanceParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.dataStackitLogmeInstance.DataStackitLogmeInstanceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference)SetInternalValue(val *DataStackitLogmeInstanceParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataStackitLogmeInstanceParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

