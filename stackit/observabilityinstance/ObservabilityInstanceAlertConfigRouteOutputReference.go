package observabilityinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/observabilityinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObservabilityInstanceAlertConfigRouteOutputReference interface {
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
	// Experimental.
	Fqn() *string
	GroupBy() *[]*string
	SetGroupBy(val *[]*string)
	GroupByInput() *[]*string
	GroupInterval() *string
	SetGroupInterval(val *string)
	GroupIntervalInput() *string
	GroupWait() *string
	SetGroupWait(val *string)
	GroupWaitInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Receiver() *string
	SetReceiver(val *string)
	ReceiverInput() *string
	RepeatInterval() *string
	SetRepeatInterval(val *string)
	RepeatIntervalInput() *string
	Routes() ObservabilityInstanceAlertConfigRouteRoutesList
	RoutesInput() interface{}
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
	PutRoutes(value interface{})
	ResetGroupBy()
	ResetGroupInterval()
	ResetGroupWait()
	ResetRepeatInterval()
	ResetRoutes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityInstanceAlertConfigRouteOutputReference
type jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) GroupBy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) GroupByInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) GroupInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) GroupIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) GroupWait() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupWait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) GroupWaitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupWaitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) Receiver() *string {
	var returns *string
	_jsii_.Get(
		j,
		"receiver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) ReceiverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"receiverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) RepeatInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repeatInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) RepeatIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repeatIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) Routes() ObservabilityInstanceAlertConfigRouteRoutesList {
	var returns ObservabilityInstanceAlertConfigRouteRoutesList
	_jsii_.Get(
		j,
		"routes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) RoutesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewObservabilityInstanceAlertConfigRouteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ObservabilityInstanceAlertConfigRouteOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityInstanceAlertConfigRouteOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference{}

	_jsii_.Create(
		"stackit.observabilityInstance.ObservabilityInstanceAlertConfigRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewObservabilityInstanceAlertConfigRouteOutputReference_Override(o ObservabilityInstanceAlertConfigRouteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.observabilityInstance.ObservabilityInstanceAlertConfigRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference)SetGroupBy(val *[]*string) {
	if err := j.validateSetGroupByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupBy",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference)SetGroupInterval(val *string) {
	if err := j.validateSetGroupIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupInterval",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference)SetGroupWait(val *string) {
	if err := j.validateSetGroupWaitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupWait",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference)SetReceiver(val *string) {
	if err := j.validateSetReceiverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"receiver",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference)SetRepeatInterval(val *string) {
	if err := j.validateSetRepeatIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repeatInterval",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) PutRoutes(value interface{}) {
	if err := o.validatePutRoutesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRoutes",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		o,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) ResetGroupInterval() {
	_jsii_.InvokeVoid(
		o,
		"resetGroupInterval",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) ResetGroupWait() {
	_jsii_.InvokeVoid(
		o,
		"resetGroupWait",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) ResetRepeatInterval() {
	_jsii_.InvokeVoid(
		o,
		"resetRepeatInterval",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) ResetRoutes() {
	_jsii_.InvokeVoid(
		o,
		"resetRoutes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

