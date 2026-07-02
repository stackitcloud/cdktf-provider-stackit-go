package vpnconnection

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/vpnconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnConnectionTunnel1OutputReference interface {
	cdktf.ComplexObject
	Bgp() VpnConnectionTunnel1BgpOutputReference
	BgpInput() interface{}
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
	Peering() VpnConnectionTunnel1PeeringOutputReference
	PeeringInput() interface{}
	Phase1() VpnConnectionTunnel1Phase1OutputReference
	Phase1Input() interface{}
	Phase2() VpnConnectionTunnel1Phase2OutputReference
	Phase2Input() interface{}
	PreSharedKey() *string
	SetPreSharedKey(val *string)
	PreSharedKeyInput() *string
	PreSharedKeyWo() *string
	SetPreSharedKeyWo(val *string)
	PreSharedKeyWoInput() *string
	PreSharedKeyWoVersion() *float64
	SetPreSharedKeyWoVersion(val *float64)
	PreSharedKeyWoVersionInput() *float64
	RemoteAddress() *string
	SetRemoteAddress(val *string)
	RemoteAddressInput() *string
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
	PutBgp(value *VpnConnectionTunnel1Bgp)
	PutPeering(value *VpnConnectionTunnel1Peering)
	PutPhase1(value *VpnConnectionTunnel1Phase1)
	PutPhase2(value *VpnConnectionTunnel1Phase2)
	ResetBgp()
	ResetPeering()
	ResetPreSharedKey()
	ResetPreSharedKeyWo()
	ResetPreSharedKeyWoVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpnConnectionTunnel1OutputReference
type jsiiProxy_VpnConnectionTunnel1OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) Bgp() VpnConnectionTunnel1BgpOutputReference {
	var returns VpnConnectionTunnel1BgpOutputReference
	_jsii_.Get(
		j,
		"bgp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) BgpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bgpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) Peering() VpnConnectionTunnel1PeeringOutputReference {
	var returns VpnConnectionTunnel1PeeringOutputReference
	_jsii_.Get(
		j,
		"peering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) PeeringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"peeringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) Phase1() VpnConnectionTunnel1Phase1OutputReference {
	var returns VpnConnectionTunnel1Phase1OutputReference
	_jsii_.Get(
		j,
		"phase1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) Phase1Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"phase1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) Phase2() VpnConnectionTunnel1Phase2OutputReference {
	var returns VpnConnectionTunnel1Phase2OutputReference
	_jsii_.Get(
		j,
		"phase2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) Phase2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"phase2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) PreSharedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSharedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) PreSharedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSharedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) PreSharedKeyWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSharedKeyWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) PreSharedKeyWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSharedKeyWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) PreSharedKeyWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preSharedKeyWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) PreSharedKeyWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preSharedKeyWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) RemoteAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) RemoteAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpnConnectionTunnel1OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VpnConnectionTunnel1OutputReference {
	_init_.Initialize()

	if err := validateNewVpnConnectionTunnel1OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpnConnectionTunnel1OutputReference{}

	_jsii_.Create(
		"stackit.vpnConnection.VpnConnectionTunnel1OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVpnConnectionTunnel1OutputReference_Override(v VpnConnectionTunnel1OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.vpnConnection.VpnConnectionTunnel1OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference)SetPreSharedKey(val *string) {
	if err := j.validateSetPreSharedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preSharedKey",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference)SetPreSharedKeyWo(val *string) {
	if err := j.validateSetPreSharedKeyWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preSharedKeyWo",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference)SetPreSharedKeyWoVersion(val *float64) {
	if err := j.validateSetPreSharedKeyWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preSharedKeyWoVersion",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference)SetRemoteAddress(val *string) {
	if err := j.validateSetRemoteAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteAddress",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnConnectionTunnel1OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) PutBgp(value *VpnConnectionTunnel1Bgp) {
	if err := v.validatePutBgpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putBgp",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) PutPeering(value *VpnConnectionTunnel1Peering) {
	if err := v.validatePutPeeringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPeering",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) PutPhase1(value *VpnConnectionTunnel1Phase1) {
	if err := v.validatePutPhase1Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPhase1",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) PutPhase2(value *VpnConnectionTunnel1Phase2) {
	if err := v.validatePutPhase2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPhase2",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) ResetBgp() {
	_jsii_.InvokeVoid(
		v,
		"resetBgp",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) ResetPeering() {
	_jsii_.InvokeVoid(
		v,
		"resetPeering",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) ResetPreSharedKey() {
	_jsii_.InvokeVoid(
		v,
		"resetPreSharedKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) ResetPreSharedKeyWo() {
	_jsii_.InvokeVoid(
		v,
		"resetPreSharedKeyWo",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) ResetPreSharedKeyWoVersion() {
	_jsii_.InvokeVoid(
		v,
		"resetPreSharedKeyWoVersion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VpnConnectionTunnel1OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

