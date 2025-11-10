package argusinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/argusinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ArgusInstanceAlertConfigRouteRoutesOutputReference interface {
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

// The jsii proxy struct for ArgusInstanceAlertConfigRouteRoutesOutputReference
type jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) GroupBy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) GroupByInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) GroupInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) GroupIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) GroupWait() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupWait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) GroupWaitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupWaitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) Match() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) MatchInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) MatchRegex() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"matchRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) MatchRegexInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"matchRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) Receiver() *string {
	var returns *string
	_jsii_.Get(
		j,
		"receiver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) ReceiverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"receiverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) RepeatInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repeatInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) RepeatIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repeatIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewArgusInstanceAlertConfigRouteRoutesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ArgusInstanceAlertConfigRouteRoutesOutputReference {
	_init_.Initialize()

	if err := validateNewArgusInstanceAlertConfigRouteRoutesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference{}

	_jsii_.Create(
		"stackit.argusInstance.ArgusInstanceAlertConfigRouteRoutesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewArgusInstanceAlertConfigRouteRoutesOutputReference_Override(a ArgusInstanceAlertConfigRouteRoutesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.argusInstance.ArgusInstanceAlertConfigRouteRoutesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference)SetGroupBy(val *[]*string) {
	if err := j.validateSetGroupByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupBy",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference)SetGroupInterval(val *string) {
	if err := j.validateSetGroupIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupInterval",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference)SetGroupWait(val *string) {
	if err := j.validateSetGroupWaitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupWait",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference)SetMatch(val *map[string]*string) {
	if err := j.validateSetMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"match",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference)SetMatchRegex(val *map[string]*string) {
	if err := j.validateSetMatchRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchRegex",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference)SetReceiver(val *string) {
	if err := j.validateSetReceiverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"receiver",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference)SetRepeatInterval(val *string) {
	if err := j.validateSetRepeatIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repeatInterval",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) ResetGroupInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) ResetGroupWait() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupWait",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) ResetMatch() {
	_jsii_.InvokeVoid(
		a,
		"resetMatch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) ResetMatchRegex() {
	_jsii_.InvokeVoid(
		a,
		"resetMatchRegex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) ResetRepeatInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetRepeatInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstanceAlertConfigRouteRoutesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

