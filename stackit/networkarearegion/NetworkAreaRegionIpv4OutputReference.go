package networkarearegion

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/networkarearegion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkAreaRegionIpv4OutputReference interface {
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
	DefaultNameservers() *[]*string
	SetDefaultNameservers(val *[]*string)
	DefaultNameserversInput() *[]*string
	DefaultPrefixLength() *float64
	SetDefaultPrefixLength(val *float64)
	DefaultPrefixLengthInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxPrefixLength() *float64
	SetMaxPrefixLength(val *float64)
	MaxPrefixLengthInput() *float64
	MinPrefixLength() *float64
	SetMinPrefixLength(val *float64)
	MinPrefixLengthInput() *float64
	NetworkRanges() NetworkAreaRegionIpv4NetworkRangesList
	NetworkRangesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransferNetwork() *string
	SetTransferNetwork(val *string)
	TransferNetworkInput() *string
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
	PutNetworkRanges(value interface{})
	ResetDefaultNameservers()
	ResetDefaultPrefixLength()
	ResetMaxPrefixLength()
	ResetMinPrefixLength()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkAreaRegionIpv4OutputReference
type jsiiProxy_NetworkAreaRegionIpv4OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference) DefaultNameservers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultNameservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference) DefaultNameserversInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultNameserversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference) DefaultPrefixLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPrefixLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference) DefaultPrefixLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPrefixLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference) MaxPrefixLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPrefixLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference) MaxPrefixLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPrefixLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference) MinPrefixLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minPrefixLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference) MinPrefixLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minPrefixLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference) NetworkRanges() NetworkAreaRegionIpv4NetworkRangesList {
	var returns NetworkAreaRegionIpv4NetworkRangesList
	_jsii_.Get(
		j,
		"networkRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference) NetworkRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference) TransferNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transferNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference) TransferNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transferNetworkInput",
		&returns,
	)
	return returns
}


func NewNetworkAreaRegionIpv4OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkAreaRegionIpv4OutputReference {
	_init_.Initialize()

	if err := validateNewNetworkAreaRegionIpv4OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkAreaRegionIpv4OutputReference{}

	_jsii_.Create(
		"stackit.networkAreaRegion.NetworkAreaRegionIpv4OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkAreaRegionIpv4OutputReference_Override(n NetworkAreaRegionIpv4OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.networkAreaRegion.NetworkAreaRegionIpv4OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference)SetDefaultNameservers(val *[]*string) {
	if err := j.validateSetDefaultNameserversParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultNameservers",
		val,
	)
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference)SetDefaultPrefixLength(val *float64) {
	if err := j.validateSetDefaultPrefixLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultPrefixLength",
		val,
	)
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference)SetMaxPrefixLength(val *float64) {
	if err := j.validateSetMaxPrefixLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPrefixLength",
		val,
	)
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference)SetMinPrefixLength(val *float64) {
	if err := j.validateSetMinPrefixLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minPrefixLength",
		val,
	)
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetworkAreaRegionIpv4OutputReference)SetTransferNetwork(val *string) {
	if err := j.validateSetTransferNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transferNetwork",
		val,
	)
}

func (n *jsiiProxy_NetworkAreaRegionIpv4OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAreaRegionIpv4OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAreaRegionIpv4OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAreaRegionIpv4OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAreaRegionIpv4OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAreaRegionIpv4OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAreaRegionIpv4OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAreaRegionIpv4OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAreaRegionIpv4OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAreaRegionIpv4OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAreaRegionIpv4OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAreaRegionIpv4OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAreaRegionIpv4OutputReference) PutNetworkRanges(value interface{}) {
	if err := n.validatePutNetworkRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putNetworkRanges",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkAreaRegionIpv4OutputReference) ResetDefaultNameservers() {
	_jsii_.InvokeVoid(
		n,
		"resetDefaultNameservers",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkAreaRegionIpv4OutputReference) ResetDefaultPrefixLength() {
	_jsii_.InvokeVoid(
		n,
		"resetDefaultPrefixLength",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkAreaRegionIpv4OutputReference) ResetMaxPrefixLength() {
	_jsii_.InvokeVoid(
		n,
		"resetMaxPrefixLength",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkAreaRegionIpv4OutputReference) ResetMinPrefixLength() {
	_jsii_.InvokeVoid(
		n,
		"resetMinPrefixLength",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkAreaRegionIpv4OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAreaRegionIpv4OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

