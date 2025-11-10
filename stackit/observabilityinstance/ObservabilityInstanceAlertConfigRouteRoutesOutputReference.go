package observabilityinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/observabilityinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObservabilityInstanceAlertConfigRouteRoutesOutputReference interface {
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
	Match() *map[string]*string
	SetMatch(val *map[string]*string)
	MatchInput() *map[string]*string
	MatchRegex() *map[string]*string
	SetMatchRegex(val *map[string]*string)
	MatchRegexInput() *map[string]*string
	Receiver() *string
	SetReceiver(val *string)
	ReceiverInput() *string
	RepeatInterval() *string
	SetRepeatInterval(val *string)
	RepeatIntervalInput() *string
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
	ResetGroupBy()
	ResetGroupInterval()
	ResetGroupWait()
	ResetMatch()
	ResetMatchRegex()
	ResetRepeatInterval()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityInstanceAlertConfigRouteRoutesOutputReference
type jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) GroupBy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) GroupByInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) GroupInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) GroupIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) GroupWait() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupWait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) GroupWaitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupWaitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) Match() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) MatchInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) MatchRegex() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"matchRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) MatchRegexInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"matchRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) Receiver() *string {
	var returns *string
	_jsii_.Get(
		j,
		"receiver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) ReceiverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"receiverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) RepeatInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repeatInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) RepeatIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repeatIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewObservabilityInstanceAlertConfigRouteRoutesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ObservabilityInstanceAlertConfigRouteRoutesOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityInstanceAlertConfigRouteRoutesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference{}

	_jsii_.Create(
		"stackit.observabilityInstance.ObservabilityInstanceAlertConfigRouteRoutesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewObservabilityInstanceAlertConfigRouteRoutesOutputReference_Override(o ObservabilityInstanceAlertConfigRouteRoutesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.observabilityInstance.ObservabilityInstanceAlertConfigRouteRoutesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference)SetGroupBy(val *[]*string) {
	if err := j.validateSetGroupByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupBy",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference)SetGroupInterval(val *string) {
	if err := j.validateSetGroupIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupInterval",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference)SetGroupWait(val *string) {
	if err := j.validateSetGroupWaitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupWait",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference)SetMatch(val *map[string]*string) {
	if err := j.validateSetMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"match",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference)SetMatchRegex(val *map[string]*string) {
	if err := j.validateSetMatchRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchRegex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference)SetReceiver(val *string) {
	if err := j.validateSetReceiverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"receiver",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference)SetRepeatInterval(val *string) {
	if err := j.validateSetRepeatIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repeatInterval",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		o,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) ResetGroupInterval() {
	_jsii_.InvokeVoid(
		o,
		"resetGroupInterval",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) ResetGroupWait() {
	_jsii_.InvokeVoid(
		o,
		"resetGroupWait",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) ResetMatch() {
	_jsii_.InvokeVoid(
		o,
		"resetMatch",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) ResetMatchRegex() {
	_jsii_.InvokeVoid(
		o,
		"resetMatchRegex",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) ResetRepeatInterval() {
	_jsii_.InvokeVoid(
		o,
		"resetRepeatInterval",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigRouteRoutesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

