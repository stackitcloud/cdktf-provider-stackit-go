package volume

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/volume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VolumeEncryptionParametersOutputReference interface {
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
	KekKeyId() *string
	SetKekKeyId(val *string)
	KekKeyIdInput() *string
	KekKeyringId() *string
	SetKekKeyringId(val *string)
	KekKeyringIdInput() *string
	KekKeyVersion() *float64
	SetKekKeyVersion(val *float64)
	KekKeyVersionInput() *float64
	KeyPayloadBase64() *string
	SetKeyPayloadBase64(val *string)
	KeyPayloadBase64Input() *string
	KeyPayloadBase64Wo() *string
	SetKeyPayloadBase64Wo(val *string)
	KeyPayloadBase64WoInput() *string
	KeyPayloadBase64WoVersion() *float64
	SetKeyPayloadBase64WoVersion(val *float64)
	KeyPayloadBase64WoVersionInput() *float64
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
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
	ResetKeyPayloadBase64()
	ResetKeyPayloadBase64Wo()
	ResetKeyPayloadBase64WoVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VolumeEncryptionParametersOutputReference
type jsiiProxy_VolumeEncryptionParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) KekKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kekKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) KekKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kekKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) KekKeyringId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kekKeyringId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) KekKeyringIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kekKeyringIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) KekKeyVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"kekKeyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) KekKeyVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"kekKeyVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) KeyPayloadBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPayloadBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) KeyPayloadBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPayloadBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) KeyPayloadBase64Wo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPayloadBase64Wo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) KeyPayloadBase64WoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPayloadBase64WoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) KeyPayloadBase64WoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyPayloadBase64WoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) KeyPayloadBase64WoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyPayloadBase64WoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVolumeEncryptionParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VolumeEncryptionParametersOutputReference {
	_init_.Initialize()

	if err := validateNewVolumeEncryptionParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VolumeEncryptionParametersOutputReference{}

	_jsii_.Create(
		"stackit.volume.VolumeEncryptionParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVolumeEncryptionParametersOutputReference_Override(v VolumeEncryptionParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.volume.VolumeEncryptionParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference)SetKekKeyId(val *string) {
	if err := j.validateSetKekKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kekKeyId",
		val,
	)
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference)SetKekKeyringId(val *string) {
	if err := j.validateSetKekKeyringIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kekKeyringId",
		val,
	)
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference)SetKekKeyVersion(val *float64) {
	if err := j.validateSetKekKeyVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kekKeyVersion",
		val,
	)
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference)SetKeyPayloadBase64(val *string) {
	if err := j.validateSetKeyPayloadBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPayloadBase64",
		val,
	)
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference)SetKeyPayloadBase64Wo(val *string) {
	if err := j.validateSetKeyPayloadBase64WoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPayloadBase64Wo",
		val,
	)
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference)SetKeyPayloadBase64WoVersion(val *float64) {
	if err := j.validateSetKeyPayloadBase64WoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPayloadBase64WoVersion",
		val,
	)
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VolumeEncryptionParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VolumeEncryptionParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VolumeEncryptionParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VolumeEncryptionParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VolumeEncryptionParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VolumeEncryptionParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VolumeEncryptionParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VolumeEncryptionParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VolumeEncryptionParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VolumeEncryptionParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VolumeEncryptionParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VolumeEncryptionParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VolumeEncryptionParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VolumeEncryptionParametersOutputReference) ResetKeyPayloadBase64() {
	_jsii_.InvokeVoid(
		v,
		"resetKeyPayloadBase64",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VolumeEncryptionParametersOutputReference) ResetKeyPayloadBase64Wo() {
	_jsii_.InvokeVoid(
		v,
		"resetKeyPayloadBase64Wo",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VolumeEncryptionParametersOutputReference) ResetKeyPayloadBase64WoVersion() {
	_jsii_.InvokeVoid(
		v,
		"resetKeyPayloadBase64WoVersion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VolumeEncryptionParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VolumeEncryptionParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

