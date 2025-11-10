package datastackitargusinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/datastackitargusinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitArgusInstanceAlertConfigGlobalOutputReference interface {
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
	InternalValue() *DataStackitArgusInstanceAlertConfigGlobal
	SetInternalValue(val *DataStackitArgusInstanceAlertConfigGlobal)
	OpsgenieApiKey() *string
	OpsgenieApiUrl() *string
	ResolveTimeout() *string
	SmtpAuthIdentity() *string
	SmtpAuthPassword() *string
	SmtpAuthUsername() *string
	SmtpFrom() *string
	SmtpSmartHost() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataStackitArgusInstanceAlertConfigGlobalOutputReference
type jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) InternalValue() *DataStackitArgusInstanceAlertConfigGlobal {
	var returns *DataStackitArgusInstanceAlertConfigGlobal
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) OpsgenieApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsgenieApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) OpsgenieApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsgenieApiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) ResolveTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) SmtpAuthIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpAuthIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) SmtpAuthPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpAuthPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) SmtpAuthUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpAuthUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) SmtpFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) SmtpSmartHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpSmartHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataStackitArgusInstanceAlertConfigGlobalOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataStackitArgusInstanceAlertConfigGlobalOutputReference {
	_init_.Initialize()

	if err := validateNewDataStackitArgusInstanceAlertConfigGlobalOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference{}

	_jsii_.Create(
		"stackit.dataStackitArgusInstance.DataStackitArgusInstanceAlertConfigGlobalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataStackitArgusInstanceAlertConfigGlobalOutputReference_Override(d DataStackitArgusInstanceAlertConfigGlobalOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.dataStackitArgusInstance.DataStackitArgusInstanceAlertConfigGlobalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference)SetInternalValue(val *DataStackitArgusInstanceAlertConfigGlobal) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitArgusInstanceAlertConfigGlobalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

