package applicationloadbalancer

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/applicationloadbalancer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationLoadBalancerOptionsOutputReference interface {
	cdktf.ComplexObject
	AccessControl() ApplicationLoadBalancerOptionsAccessControlOutputReference
	AccessControlInput() interface{}
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
	EphemeralAddress() interface{}
	SetEphemeralAddress(val interface{})
	EphemeralAddressInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Observability() ApplicationLoadBalancerOptionsObservabilityOutputReference
	ObservabilityInput() interface{}
	PrivateNetworkOnly() interface{}
	SetPrivateNetworkOnly(val interface{})
	PrivateNetworkOnlyInput() interface{}
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
	PutAccessControl(value *ApplicationLoadBalancerOptionsAccessControl)
	PutObservability(value *ApplicationLoadBalancerOptionsObservability)
	ResetAccessControl()
	ResetEphemeralAddress()
	ResetObservability()
	ResetPrivateNetworkOnly()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationLoadBalancerOptionsOutputReference
type jsiiProxy_ApplicationLoadBalancerOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) AccessControl() ApplicationLoadBalancerOptionsAccessControlOutputReference {
	var returns ApplicationLoadBalancerOptionsAccessControlOutputReference
	_jsii_.Get(
		j,
		"accessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) AccessControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) EphemeralAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) EphemeralAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) Observability() ApplicationLoadBalancerOptionsObservabilityOutputReference {
	var returns ApplicationLoadBalancerOptionsObservabilityOutputReference
	_jsii_.Get(
		j,
		"observability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) ObservabilityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"observabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) PrivateNetworkOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateNetworkOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) PrivateNetworkOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateNetworkOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationLoadBalancerOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApplicationLoadBalancerOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationLoadBalancerOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationLoadBalancerOptionsOutputReference{}

	_jsii_.Create(
		"stackit.applicationLoadBalancer.ApplicationLoadBalancerOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationLoadBalancerOptionsOutputReference_Override(a ApplicationLoadBalancerOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.applicationLoadBalancer.ApplicationLoadBalancerOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference)SetEphemeralAddress(val interface{}) {
	if err := j.validateSetEphemeralAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ephemeralAddress",
		val,
	)
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference)SetPrivateNetworkOnly(val interface{}) {
	if err := j.validateSetPrivateNetworkOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateNetworkOnly",
		val,
	)
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) PutAccessControl(value *ApplicationLoadBalancerOptionsAccessControl) {
	if err := a.validatePutAccessControlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessControl",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) PutObservability(value *ApplicationLoadBalancerOptionsObservability) {
	if err := a.validatePutObservabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putObservability",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) ResetAccessControl() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessControl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) ResetEphemeralAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetEphemeralAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) ResetObservability() {
	_jsii_.InvokeVoid(
		a,
		"resetObservability",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) ResetPrivateNetworkOnly() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateNetworkOnly",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApplicationLoadBalancerOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

