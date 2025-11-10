package datastackitobservabilityinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/datastackitobservabilityinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitObservabilityInstanceAlertConfigGlobalOutputReference interface {
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
	InternalValue() *DataStackitObservabilityInstanceAlertConfigGlobal
	SetInternalValue(val *DataStackitObservabilityInstanceAlertConfigGlobal)
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

// The jsii proxy struct for DataStackitObservabilityInstanceAlertConfigGlobalOutputReference
type jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) InternalValue() *DataStackitObservabilityInstanceAlertConfigGlobal {
	var returns *DataStackitObservabilityInstanceAlertConfigGlobal
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) OpsgenieApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsgenieApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) OpsgenieApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsgenieApiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) ResolveTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) SmtpAuthIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpAuthIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) SmtpAuthPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpAuthPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) SmtpAuthUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpAuthUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) SmtpFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) SmtpSmartHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpSmartHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataStackitObservabilityInstanceAlertConfigGlobalOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataStackitObservabilityInstanceAlertConfigGlobalOutputReference {
	_init_.Initialize()

	if err := validateNewDataStackitObservabilityInstanceAlertConfigGlobalOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference{}

	_jsii_.Create(
		"stackit.dataStackitObservabilityInstance.DataStackitObservabilityInstanceAlertConfigGlobalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataStackitObservabilityInstanceAlertConfigGlobalOutputReference_Override(d DataStackitObservabilityInstanceAlertConfigGlobalOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.dataStackitObservabilityInstance.DataStackitObservabilityInstanceAlertConfigGlobalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference)SetInternalValue(val *DataStackitObservabilityInstanceAlertConfigGlobal) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataStackitObservabilityInstanceAlertConfigGlobalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

