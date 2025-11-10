package observabilityinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/observabilityinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObservabilityInstanceAlertConfigGlobalOutputReference interface {
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

// The jsii proxy struct for ObservabilityInstanceAlertConfigGlobalOutputReference
type jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) OpsgenieApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsgenieApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) OpsgenieApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsgenieApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) OpsgenieApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsgenieApiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) OpsgenieApiUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsgenieApiUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) ResolveTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) ResolveTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolveTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) SmtpAuthIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpAuthIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) SmtpAuthIdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpAuthIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) SmtpAuthPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpAuthPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) SmtpAuthPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpAuthPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) SmtpAuthUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpAuthUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) SmtpAuthUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpAuthUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) SmtpFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) SmtpFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) SmtpSmartHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpSmartHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) SmtpSmartHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpSmartHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewObservabilityInstanceAlertConfigGlobalOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ObservabilityInstanceAlertConfigGlobalOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityInstanceAlertConfigGlobalOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference{}

	_jsii_.Create(
		"stackit.observabilityInstance.ObservabilityInstanceAlertConfigGlobalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewObservabilityInstanceAlertConfigGlobalOutputReference_Override(o ObservabilityInstanceAlertConfigGlobalOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.observabilityInstance.ObservabilityInstanceAlertConfigGlobalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference)SetOpsgenieApiKey(val *string) {
	if err := j.validateSetOpsgenieApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opsgenieApiKey",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference)SetOpsgenieApiUrl(val *string) {
	if err := j.validateSetOpsgenieApiUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opsgenieApiUrl",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference)SetResolveTimeout(val *string) {
	if err := j.validateSetResolveTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resolveTimeout",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference)SetSmtpAuthIdentity(val *string) {
	if err := j.validateSetSmtpAuthIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smtpAuthIdentity",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference)SetSmtpAuthPassword(val *string) {
	if err := j.validateSetSmtpAuthPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smtpAuthPassword",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference)SetSmtpAuthUsername(val *string) {
	if err := j.validateSetSmtpAuthUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smtpAuthUsername",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference)SetSmtpFrom(val *string) {
	if err := j.validateSetSmtpFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smtpFrom",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference)SetSmtpSmartHost(val *string) {
	if err := j.validateSetSmtpSmartHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smtpSmartHost",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) ResetOpsgenieApiKey() {
	_jsii_.InvokeVoid(
		o,
		"resetOpsgenieApiKey",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) ResetOpsgenieApiUrl() {
	_jsii_.InvokeVoid(
		o,
		"resetOpsgenieApiUrl",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) ResetResolveTimeout() {
	_jsii_.InvokeVoid(
		o,
		"resetResolveTimeout",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) ResetSmtpAuthIdentity() {
	_jsii_.InvokeVoid(
		o,
		"resetSmtpAuthIdentity",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) ResetSmtpAuthPassword() {
	_jsii_.InvokeVoid(
		o,
		"resetSmtpAuthPassword",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) ResetSmtpAuthUsername() {
	_jsii_.InvokeVoid(
		o,
		"resetSmtpAuthUsername",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) ResetSmtpFrom() {
	_jsii_.InvokeVoid(
		o,
		"resetSmtpFrom",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) ResetSmtpSmartHost() {
	_jsii_.InvokeVoid(
		o,
		"resetSmtpSmartHost",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityInstanceAlertConfigGlobalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

