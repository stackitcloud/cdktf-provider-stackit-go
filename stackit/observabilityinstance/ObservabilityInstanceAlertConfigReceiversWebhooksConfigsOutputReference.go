package observabilityinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/observabilityinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference interface {
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
	GoogleChat() interface{}
	SetGoogleChat(val interface{})
	GoogleChatInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MsTeams() interface{}
	SetMsTeams(val interface{})
	MsTeamsInput() interface{}
	SendResolved() interface{}
	SetSendResolved(val interface{})
	SendResolvedInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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
	ResetGoogleChat()
	ResetMsTeams()
	ResetSendResolved()
	ResetUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference
type jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) GoogleChat() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"googleChat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) GoogleChatInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"googleChatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) MsTeams() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"msTeams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) MsTeamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"msTeamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) SendResolved() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendResolved",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) SendResolvedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendResolvedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference{}

	_jsii_.Create(
		"stackit.observabilityInstance.ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference_Override(o ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.observabilityInstance.ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference)SetGoogleChat(val interface{}) {
	if err := j.validateSetGoogleChatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleChat",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference)SetMsTeams(val interface{}) {
	if err := j.validateSetMsTeamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"msTeams",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference)SetSendResolved(val interface{}) {
	if err := j.validateSetSendResolvedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendResolved",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) ResetGoogleChat() {
	_jsii_.InvokeVoid(
		o,
		"resetGoogleChat",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) ResetMsTeams() {
	_jsii_.InvokeVoid(
		o,
		"resetMsTeams",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) ResetSendResolved() {
	_jsii_.InvokeVoid(
		o,
		"resetSendResolved",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		o,
		"resetUrl",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversWebhooksConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

