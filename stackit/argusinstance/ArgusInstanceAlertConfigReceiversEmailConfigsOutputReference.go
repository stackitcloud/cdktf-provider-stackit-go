package argusinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/argusinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference interface {
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

// The jsii proxy struct for ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference
type jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) AuthIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) AuthIdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) AuthPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) AuthPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) AuthUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) AuthUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) From() *string {
	var returns *string
	_jsii_.Get(
		j,
		"from",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) FromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) SmartHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smartHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) SmartHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smartHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) To() *string {
	var returns *string
	_jsii_.Get(
		j,
		"to",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) ToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toInput",
		&returns,
	)
	return returns
}


func NewArgusInstanceAlertConfigReceiversEmailConfigsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewArgusInstanceAlertConfigReceiversEmailConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference{}

	_jsii_.Create(
		"stackit.argusInstance.ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewArgusInstanceAlertConfigReceiversEmailConfigsOutputReference_Override(a ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.argusInstance.ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference)SetAuthIdentity(val *string) {
	if err := j.validateSetAuthIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authIdentity",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference)SetAuthPassword(val *string) {
	if err := j.validateSetAuthPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authPassword",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference)SetAuthUsername(val *string) {
	if err := j.validateSetAuthUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authUsername",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference)SetFrom(val *string) {
	if err := j.validateSetFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"from",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference)SetSmartHost(val *string) {
	if err := j.validateSetSmartHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smartHost",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference)SetTo(val *string) {
	if err := j.validateSetToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"to",
		val,
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) ResetAuthIdentity() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthIdentity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) ResetAuthPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) ResetAuthUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) ResetFrom() {
	_jsii_.InvokeVoid(
		a,
		"resetFrom",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) ResetSmartHost() {
	_jsii_.InvokeVoid(
		a,
		"resetSmartHost",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) ResetTo() {
	_jsii_.InvokeVoid(
		a,
		"resetTo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversEmailConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

