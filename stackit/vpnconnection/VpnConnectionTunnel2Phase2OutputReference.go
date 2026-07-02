package vpnconnection

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/vpnconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnConnectionTunnel2Phase2OutputReference interface {
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
	DhGroups() *[]*string
	SetDhGroups(val *[]*string)
	DhGroupsInput() *[]*string
	DpdAction() *string
	SetDpdAction(val *string)
	DpdActionInput() *string
	EncryptionAlgorithms() *[]*string
	SetEncryptionAlgorithms(val *[]*string)
	EncryptionAlgorithmsInput() *[]*string
	// Experimental.
	Fqn() *string
	IntegrityAlgorithms() *[]*string
	SetIntegrityAlgorithms(val *[]*string)
	IntegrityAlgorithmsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RekeyTime() *float64
	SetRekeyTime(val *float64)
	RekeyTimeInput() *float64
	StartAction() *string
	SetStartAction(val *string)
	StartActionInput() *string
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
	ResetDhGroups()
	ResetDpdAction()
	ResetRekeyTime()
	ResetStartAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpnConnectionTunnel2Phase2OutputReference
type jsiiProxy_VpnConnectionTunnel2Phase2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) DhGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dhGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) DhGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dhGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) DpdAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dpdAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) DpdActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dpdActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) EncryptionAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"encryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) EncryptionAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"encryptionAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) IntegrityAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"integrityAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) IntegrityAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"integrityAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) RekeyTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rekeyTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) RekeyTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rekeyTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) StartAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) StartActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpnConnectionTunnel2Phase2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VpnConnectionTunnel2Phase2OutputReference {
	_init_.Initialize()

	if err := validateNewVpnConnectionTunnel2Phase2OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpnConnectionTunnel2Phase2OutputReference{}

	_jsii_.Create(
		"stackit.vpnConnection.VpnConnectionTunnel2Phase2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVpnConnectionTunnel2Phase2OutputReference_Override(v VpnConnectionTunnel2Phase2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.vpnConnection.VpnConnectionTunnel2Phase2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference)SetDhGroups(val *[]*string) {
	if err := j.validateSetDhGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dhGroups",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference)SetDpdAction(val *string) {
	if err := j.validateSetDpdActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dpdAction",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference)SetEncryptionAlgorithms(val *[]*string) {
	if err := j.validateSetEncryptionAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionAlgorithms",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference)SetIntegrityAlgorithms(val *[]*string) {
	if err := j.validateSetIntegrityAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrityAlgorithms",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference)SetRekeyTime(val *float64) {
	if err := j.validateSetRekeyTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rekeyTime",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference)SetStartAction(val *string) {
	if err := j.validateSetStartActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startAction",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) ResetDhGroups() {
	_jsii_.InvokeVoid(
		v,
		"resetDhGroups",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) ResetDpdAction() {
	_jsii_.InvokeVoid(
		v,
		"resetDpdAction",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) ResetRekeyTime() {
	_jsii_.InvokeVoid(
		v,
		"resetRekeyTime",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) ResetStartAction() {
	_jsii_.InvokeVoid(
		v,
		"resetStartAction",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel2Phase2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

