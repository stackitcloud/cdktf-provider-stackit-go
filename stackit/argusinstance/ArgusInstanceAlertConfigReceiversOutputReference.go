package argusinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/argusinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ArgusInstanceAlertConfigReceiversOutputReference interface {
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
	EmailConfigs() ArgusInstanceAlertConfigReceiversEmailConfigsList
	EmailConfigsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	OpsgenieConfigs() ArgusInstanceAlertConfigReceiversOpsgenieConfigsList
	OpsgenieConfigsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebhooksConfigs() ArgusInstanceAlertConfigReceiversWebhooksConfigsList
	WebhooksConfigsInput() interface{}
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
	PutEmailConfigs(value interface{})
	PutOpsgenieConfigs(value interface{})
	PutWebhooksConfigs(value interface{})
	ResetEmailConfigs()
	ResetOpsgenieConfigs()
	ResetWebhooksConfigs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ArgusInstanceAlertConfigReceiversOutputReference
type jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) EmailConfigs() ArgusInstanceAlertConfigReceiversEmailConfigsList {
	var returns ArgusInstanceAlertConfigReceiversEmailConfigsList
	_jsii_.Get(
		j,
		"emailConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) EmailConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) OpsgenieConfigs() ArgusInstanceAlertConfigReceiversOpsgenieConfigsList {
	var returns ArgusInstanceAlertConfigReceiversOpsgenieConfigsList
	_jsii_.Get(
		j,
		"opsgenieConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) OpsgenieConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opsgenieConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) WebhooksConfigs() ArgusInstanceAlertConfigReceiversWebhooksConfigsList {
	var returns ArgusInstanceAlertConfigReceiversWebhooksConfigsList
	_jsii_.Get(
		j,
		"webhooksConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) WebhooksConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webhooksConfigsInput",
		&returns,
	)
	return returns
}


func NewArgusInstanceAlertConfigReceiversOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ArgusInstanceAlertConfigReceiversOutputReference {
	_init_.Initialize()

	if err := validateNewArgusInstanceAlertConfigReceiversOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference{}

	_jsii_.Create(
		"stackit.argusInstance.ArgusInstanceAlertConfigReceiversOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewArgusInstanceAlertConfigReceiversOutputReference_Override(a ArgusInstanceAlertConfigReceiversOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.argusInstance.ArgusInstanceAlertConfigReceiversOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) PutEmailConfigs(value interface{}) {
	if err := a.validatePutEmailConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEmailConfigs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) PutOpsgenieConfigs(value interface{}) {
	if err := a.validatePutOpsgenieConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpsgenieConfigs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) PutWebhooksConfigs(value interface{}) {
	if err := a.validatePutWebhooksConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWebhooksConfigs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) ResetEmailConfigs() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailConfigs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) ResetOpsgenieConfigs() {
	_jsii_.InvokeVoid(
		a,
		"resetOpsgenieConfigs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) ResetWebhooksConfigs() {
	_jsii_.InvokeVoid(
		a,
		"resetWebhooksConfigs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ArgusInstanceAlertConfigReceiversOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

