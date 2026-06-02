package telemetryrouterdestination

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/telemetryrouterdestination/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TelemetryrouterDestinationConfigAOutputReference interface {
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
	SetConfigType(val *string)
	ConfigTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Filter() TelemetryrouterDestinationConfigFilterOutputReference
	FilterInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Opentelemetry() TelemetryrouterDestinationConfigOpentelemetryOutputReference
	OpentelemetryInput() interface{}
	S3() TelemetryrouterDestinationConfigS3OutputReference
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
	PutFilter(value *TelemetryrouterDestinationConfigFilter)
	PutOpentelemetry(value *TelemetryrouterDestinationConfigOpentelemetry)
	PutS3(value *TelemetryrouterDestinationConfigS3)
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

// The jsii proxy struct for TelemetryrouterDestinationConfigAOutputReference
type jsiiProxy_TelemetryrouterDestinationConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) ConfigType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) ConfigTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) Filter() TelemetryrouterDestinationConfigFilterOutputReference {
	var returns TelemetryrouterDestinationConfigFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) Opentelemetry() TelemetryrouterDestinationConfigOpentelemetryOutputReference {
	var returns TelemetryrouterDestinationConfigOpentelemetryOutputReference
	_jsii_.Get(
		j,
		"opentelemetry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) OpentelemetryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opentelemetryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) S3() TelemetryrouterDestinationConfigS3OutputReference {
	var returns TelemetryrouterDestinationConfigS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) S3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTelemetryrouterDestinationConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TelemetryrouterDestinationConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewTelemetryrouterDestinationConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TelemetryrouterDestinationConfigAOutputReference{}

	_jsii_.Create(
		"stackit.telemetryrouterDestination.TelemetryrouterDestinationConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTelemetryrouterDestinationConfigAOutputReference_Override(t TelemetryrouterDestinationConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.telemetryrouterDestination.TelemetryrouterDestinationConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference)SetConfigType(val *string) {
	if err := j.validateSetConfigTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configType",
		val,
	)
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) PutFilter(value *TelemetryrouterDestinationConfigFilter) {
	if err := t.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFilter",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) PutOpentelemetry(value *TelemetryrouterDestinationConfigOpentelemetry) {
	if err := t.validatePutOpentelemetryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOpentelemetry",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) PutS3(value *TelemetryrouterDestinationConfigS3) {
	if err := t.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		t,
		"resetFilter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) ResetOpentelemetry() {
	_jsii_.InvokeVoid(
		t,
		"resetOpentelemetry",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		t,
		"resetS3",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TelemetryrouterDestinationConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

