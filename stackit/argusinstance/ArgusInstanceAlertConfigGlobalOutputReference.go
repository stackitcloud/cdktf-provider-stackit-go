package argusinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/argusinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ArgusInstanceAlertConfigGlobalOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OpsgenieApiKey() *string
	SetOpsgenieApiKey(val *string)
	OpsgenieApiKeyInput() *string
	OpsgenieApiUrl() *string
	SetOpsgenieApiUrl(val *string)
	OpsgenieApiUrlInput() *string
	ResolveTimeout() *string
	SetResolveTimeout(val *string)
	ResolveTimeoutInput() *string
	SmtpAuthIdentity() *string
	SetSmtpAuthIdentity(val *string)
	SmtpAuthIdentityInput() *string
	SmtpAuthPassword() *string
	SetSmtpAuthPassword(val *string)
	SmtpAuthPasswordInput() *string
	SmtpAuthUsername() *string
	SetSmtpAuthUsername(val *string)
	SmtpAuthUsernameInput() *string
	SmtpFrom() *string
	SetSmtpFrom(val *string)
	SmtpFromInput() *string
	SmtpSmartHost() *string
	SetSmtpSmartHost(val *string)
	SmtpSmartHostInput() *string
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
	ResetOpsgenieApiKey()
	ResetOpsgenieApiUrl()
	ResetResolveTimeout()
	ResetSmtpAuthIdentity()
	ResetSmtpAuthPassword()
	ResetSmtpAuthUsername()
	ResetSmtpFrom()
	ResetSmtpSmartHost()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ArgusInstanceAlertConfigGlobalOutputReference
type jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) OpsgenieApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsgenieApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) OpsgenieApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsgenieApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) OpsgenieApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsgenieApiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) OpsgenieApiUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsgenieApiUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) ResolveTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) ResolveTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolveTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) SmtpAuthIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpAuthIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) SmtpAuthIdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpAuthIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) SmtpAuthPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpAuthPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) SmtpAuthPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpAuthPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) SmtpAuthUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpAuthUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) SmtpAuthUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpAuthUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) SmtpFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) SmtpFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) SmtpSmartHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpSmartHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) SmtpSmartHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpSmartHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewArgusInstanceAlertConfigGlobalOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ArgusInstanceAlertConfigGlobalOutputReference {
	_init_.Initialize()

	if err := validateNewArgusInstanceAlertConfigGlobalOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference{}

	_jsii_.Create(
		"stackit.argusInstance.ArgusInstanceAlertConfigGlobalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewArgusInstanceAlertConfigGlobalOutputReference_Override(a ArgusInstanceAlertConfigGlobalOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.argusInstance.ArgusInstanceAlertConfigGlobalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference)SetOpsgenieApiKey(val *string) {
	if err := j.validateSetOpsgenieApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opsgenieApiKey",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference)SetOpsgenieApiUrl(val *string) {
	if err := j.validateSetOpsgenieApiUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opsgenieApiUrl",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference)SetResolveTimeout(val *string) {
	if err := j.validateSetResolveTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resolveTimeout",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference)SetSmtpAuthIdentity(val *string) {
	if err := j.validateSetSmtpAuthIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smtpAuthIdentity",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference)SetSmtpAuthPassword(val *string) {
	if err := j.validateSetSmtpAuthPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smtpAuthPassword",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference)SetSmtpAuthUsername(val *string) {
	if err := j.validateSetSmtpAuthUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smtpAuthUsername",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference)SetSmtpFrom(val *string) {
	if err := j.validateSetSmtpFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smtpFrom",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference)SetSmtpSmartHost(val *string) {
	if err := j.validateSetSmtpSmartHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smtpSmartHost",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) ResetOpsgenieApiKey() {
	_jsii_.InvokeVoid(
		a,
		"resetOpsgenieApiKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) ResetOpsgenieApiUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetOpsgenieApiUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) ResetResolveTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetResolveTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) ResetSmtpAuthIdentity() {
	_jsii_.InvokeVoid(
		a,
		"resetSmtpAuthIdentity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) ResetSmtpAuthPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetSmtpAuthPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) ResetSmtpAuthUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetSmtpAuthUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) ResetSmtpFrom() {
	_jsii_.InvokeVoid(
		a,
		"resetSmtpFrom",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) ResetSmtpSmartHost() {
	_jsii_.InvokeVoid(
		a,
		"resetSmtpSmartHost",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigGlobalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

