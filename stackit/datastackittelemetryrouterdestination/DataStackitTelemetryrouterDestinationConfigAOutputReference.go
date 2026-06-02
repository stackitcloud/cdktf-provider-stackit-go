package datastackittelemetryrouterdestination

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/datastackittelemetryrouterdestination/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitTelemetryrouterDestinationConfigAOutputReference interface {
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
	ConfigType() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Filter() DataStackitTelemetryrouterDestinationConfigFilterOutputReference
	FilterInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DataStackitTelemetryrouterDestinationConfigA
	SetInternalValue(val *DataStackitTelemetryrouterDestinationConfigA)
	Opentelemetry() DataStackitTelemetryrouterDestinationConfigOpentelemetryOutputReference
	OpentelemetryInput() interface{}
	S3() DataStackitTelemetryrouterDestinationConfigS3OutputReference
	S3Input() interface{}
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
	PutFilter(value *DataStackitTelemetryrouterDestinationConfigFilter)
	PutOpentelemetry(value *DataStackitTelemetryrouterDestinationConfigOpentelemetry)
	PutS3(value *DataStackitTelemetryrouterDestinationConfigS3)
	ResetFilter()
	ResetOpentelemetry()
	ResetS3()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataStackitTelemetryrouterDestinationConfigAOutputReference
type jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) ConfigType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) Filter() DataStackitTelemetryrouterDestinationConfigFilterOutputReference {
	var returns DataStackitTelemetryrouterDestinationConfigFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) InternalValue() *DataStackitTelemetryrouterDestinationConfigA {
	var returns *DataStackitTelemetryrouterDestinationConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) Opentelemetry() DataStackitTelemetryrouterDestinationConfigOpentelemetryOutputReference {
	var returns DataStackitTelemetryrouterDestinationConfigOpentelemetryOutputReference
	_jsii_.Get(
		j,
		"opentelemetry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) OpentelemetryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opentelemetryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) S3() DataStackitTelemetryrouterDestinationConfigS3OutputReference {
	var returns DataStackitTelemetryrouterDestinationConfigS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) S3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataStackitTelemetryrouterDestinationConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataStackitTelemetryrouterDestinationConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewDataStackitTelemetryrouterDestinationConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference{}

	_jsii_.Create(
		"stackit.dataStackitTelemetryrouterDestination.DataStackitTelemetryrouterDestinationConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataStackitTelemetryrouterDestinationConfigAOutputReference_Override(d DataStackitTelemetryrouterDestinationConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.dataStackitTelemetryrouterDestination.DataStackitTelemetryrouterDestinationConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference)SetInternalValue(val *DataStackitTelemetryrouterDestinationConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) PutFilter(value *DataStackitTelemetryrouterDestinationConfigFilter) {
	if err := d.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) PutOpentelemetry(value *DataStackitTelemetryrouterDestinationConfigOpentelemetry) {
	if err := d.validatePutOpentelemetryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOpentelemetry",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) PutS3(value *DataStackitTelemetryrouterDestinationConfigS3) {
	if err := d.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putS3",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) ResetOpentelemetry() {
	_jsii_.InvokeVoid(
		d,
		"resetOpentelemetry",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		d,
		"resetS3",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataStackitTelemetryrouterDestinationConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

