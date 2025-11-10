package observabilityinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/observabilityinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference interface {
	cdktf.ComplexObject
	AuthIdentity() *string
	SetAuthIdentity(val *string)
	AuthIdentityInput() *string
	AuthPassword() *string
	SetAuthPassword(val *string)
	AuthPasswordInput() *string
	AuthUsername() *string
	SetAuthUsername(val *string)
	AuthUsernameInput() *string
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
	From() *string
	SetFrom(val *string)
	FromInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SmartHost() *string
	SetSmartHost(val *string)
	SmartHostInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	To() *string
	SetTo(val *string)
	ToInput() *string
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
	ResetAuthIdentity()
	ResetAuthPassword()
	ResetAuthUsername()
	ResetFrom()
	ResetSmartHost()
	ResetTo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference
type jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) AuthIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) AuthIdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) AuthPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) AuthPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) AuthUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) AuthUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) From() *string {
	var returns *string
	_jsii_.Get(
		j,
		"from",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) FromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) SmartHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smartHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) SmartHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smartHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) To() *string {
	var returns *string
	_jsii_.Get(
		j,
		"to",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) ToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toInput",
		&returns,
	)
	return returns
}


func NewObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference{}

	_jsii_.Create(
		"stackit.observabilityInstance.ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference_Override(o ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.observabilityInstance.ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference)SetAuthIdentity(val *string) {
	if err := j.validateSetAuthIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authIdentity",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference)SetAuthPassword(val *string) {
	if err := j.validateSetAuthPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authPassword",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference)SetAuthUsername(val *string) {
	if err := j.validateSetAuthUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authUsername",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference)SetFrom(val *string) {
	if err := j.validateSetFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"from",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference)SetSmartHost(val *string) {
	if err := j.validateSetSmartHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smartHost",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference)SetTo(val *string) {
	if err := j.validateSetToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"to",
		val,
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) ResetAuthIdentity() {
	_jsii_.InvokeVoid(
		o,
		"resetAuthIdentity",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) ResetAuthPassword() {
	_jsii_.InvokeVoid(
		o,
		"resetAuthPassword",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) ResetAuthUsername() {
	_jsii_.InvokeVoid(
		o,
		"resetAuthUsername",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) ResetFrom() {
	_jsii_.InvokeVoid(
		o,
		"resetFrom",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) ResetSmartHost() {
	_jsii_.InvokeVoid(
		o,
		"resetSmartHost",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) ResetTo() {
	_jsii_.InvokeVoid(
		o,
		"resetTo",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigReceiversEmailConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

