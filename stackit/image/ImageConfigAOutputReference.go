package image

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/image/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImageConfigAOutputReference interface {
	cdktf.ComplexObject
	BootMenu() interface{}
	SetBootMenu(val interface{})
	BootMenuInput() interface{}
	CdromBus() *string
	SetCdromBus(val *string)
	CdromBusInput() *string
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
	DiskBus() *string
	SetDiskBus(val *string)
	DiskBusInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NicModel() *string
	SetNicModel(val *string)
	NicModelInput() *string
	OperatingSystem() *string
	SetOperatingSystem(val *string)
	OperatingSystemDistro() *string
	SetOperatingSystemDistro(val *string)
	OperatingSystemDistroInput() *string
	OperatingSystemInput() *string
	OperatingSystemVersion() *string
	SetOperatingSystemVersion(val *string)
	OperatingSystemVersionInput() *string
	RescueBus() *string
	SetRescueBus(val *string)
	RescueBusInput() *string
	RescueDevice() *string
	SetRescueDevice(val *string)
	RescueDeviceInput() *string
	SecureBoot() interface{}
	SetSecureBoot(val interface{})
	SecureBootInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uefi() interface{}
	SetUefi(val interface{})
	UefiInput() interface{}
	VideoModel() *string
	SetVideoModel(val *string)
	VideoModelInput() *string
	VirtioScsi() interface{}
	SetVirtioScsi(val interface{})
	VirtioScsiInput() interface{}
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
	ResetBootMenu()
	ResetCdromBus()
	ResetDiskBus()
	ResetNicModel()
	ResetOperatingSystem()
	ResetOperatingSystemDistro()
	ResetOperatingSystemVersion()
	ResetRescueBus()
	ResetRescueDevice()
	ResetSecureBoot()
	ResetUefi()
	ResetVideoModel()
	ResetVirtioScsi()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ImageConfigAOutputReference
type jsiiProxy_ImageConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ImageConfigAOutputReference) BootMenu() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootMenu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) BootMenuInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootMenuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) CdromBus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdromBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) CdromBusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdromBusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) DiskBus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) DiskBusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskBusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) NicModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) NicModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) OperatingSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) OperatingSystemDistro() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystemDistro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) OperatingSystemDistroInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystemDistroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) OperatingSystemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) OperatingSystemVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystemVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) OperatingSystemVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystemVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) RescueBus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rescueBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) RescueBusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rescueBusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) RescueDevice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rescueDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) RescueDeviceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rescueDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) SecureBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secureBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) SecureBootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secureBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) Uefi() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uefi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) UefiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uefiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) VideoModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) VideoModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) VirtioScsi() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"virtioScsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageConfigAOutputReference) VirtioScsiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"virtioScsiInput",
		&returns,
	)
	return returns
}


func NewImageConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ImageConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewImageConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ImageConfigAOutputReference{}

	_jsii_.Create(
		"stackit.image.ImageConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewImageConfigAOutputReference_Override(i ImageConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.image.ImageConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_ImageConfigAOutputReference)SetBootMenu(val interface{}) {
	if err := j.validateSetBootMenuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootMenu",
		val,
	)
}

func (j *jsiiProxy_ImageConfigAOutputReference)SetCdromBus(val *string) {
	if err := j.validateSetCdromBusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdromBus",
		val,
	)
}

func (j *jsiiProxy_ImageConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ImageConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ImageConfigAOutputReference)SetDiskBus(val *string) {
	if err := j.validateSetDiskBusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskBus",
		val,
	)
}

func (j *jsiiProxy_ImageConfigAOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ImageConfigAOutputReference)SetNicModel(val *string) {
	if err := j.validateSetNicModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nicModel",
		val,
	)
}

func (j *jsiiProxy_ImageConfigAOutputReference)SetOperatingSystem(val *string) {
	if err := j.validateSetOperatingSystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operatingSystem",
		val,
	)
}

func (j *jsiiProxy_ImageConfigAOutputReference)SetOperatingSystemDistro(val *string) {
	if err := j.validateSetOperatingSystemDistroParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operatingSystemDistro",
		val,
	)
}

func (j *jsiiProxy_ImageConfigAOutputReference)SetOperatingSystemVersion(val *string) {
	if err := j.validateSetOperatingSystemVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operatingSystemVersion",
		val,
	)
}

func (j *jsiiProxy_ImageConfigAOutputReference)SetRescueBus(val *string) {
	if err := j.validateSetRescueBusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rescueBus",
		val,
	)
}

func (j *jsiiProxy_ImageConfigAOutputReference)SetRescueDevice(val *string) {
	if err := j.validateSetRescueDeviceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rescueDevice",
		val,
	)
}

func (j *jsiiProxy_ImageConfigAOutputReference)SetSecureBoot(val interface{}) {
	if err := j.validateSetSecureBootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secureBoot",
		val,
	)
}

func (j *jsiiProxy_ImageConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ImageConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ImageConfigAOutputReference)SetUefi(val interface{}) {
	if err := j.validateSetUefiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uefi",
		val,
	)
}

func (j *jsiiProxy_ImageConfigAOutputReference)SetVideoModel(val *string) {
	if err := j.validateSetVideoModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"videoModel",
		val,
	)
}

func (j *jsiiProxy_ImageConfigAOutputReference)SetVirtioScsi(val interface{}) {
	if err := j.validateSetVirtioScsiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtioScsi",
		val,
	)
}

func (i *jsiiProxy_ImageConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageConfigAOutputReference) ResetBootMenu() {
	_jsii_.InvokeVoid(
		i,
		"resetBootMenu",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageConfigAOutputReference) ResetCdromBus() {
	_jsii_.InvokeVoid(
		i,
		"resetCdromBus",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageConfigAOutputReference) ResetDiskBus() {
	_jsii_.InvokeVoid(
		i,
		"resetDiskBus",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageConfigAOutputReference) ResetNicModel() {
	_jsii_.InvokeVoid(
		i,
		"resetNicModel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageConfigAOutputReference) ResetOperatingSystem() {
	_jsii_.InvokeVoid(
		i,
		"resetOperatingSystem",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageConfigAOutputReference) ResetOperatingSystemDistro() {
	_jsii_.InvokeVoid(
		i,
		"resetOperatingSystemDistro",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageConfigAOutputReference) ResetOperatingSystemVersion() {
	_jsii_.InvokeVoid(
		i,
		"resetOperatingSystemVersion",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageConfigAOutputReference) ResetRescueBus() {
	_jsii_.InvokeVoid(
		i,
		"resetRescueBus",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageConfigAOutputReference) ResetRescueDevice() {
	_jsii_.InvokeVoid(
		i,
		"resetRescueDevice",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageConfigAOutputReference) ResetSecureBoot() {
	_jsii_.InvokeVoid(
		i,
		"resetSecureBoot",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageConfigAOutputReference) ResetUefi() {
	_jsii_.InvokeVoid(
		i,
		"resetUefi",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageConfigAOutputReference) ResetVideoModel() {
	_jsii_.InvokeVoid(
		i,
		"resetVideoModel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageConfigAOutputReference) ResetVirtioScsi() {
	_jsii_.InvokeVoid(
		i,
		"resetVirtioScsi",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

