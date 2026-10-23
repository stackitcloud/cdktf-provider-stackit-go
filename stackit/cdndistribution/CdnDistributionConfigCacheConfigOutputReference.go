package cdndistribution

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/cdndistribution/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnDistributionConfigCacheConfigOutputReference interface {
	cdktf.ComplexObject
	CacheKeyHeaders() *[]*string
	SetCacheKeyHeaders(val *[]*string)
	CacheKeyHeadersInput() *[]*string
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
	QueryStringVaryEnabled() interface{}
	SetQueryStringVaryEnabled(val interface{})
	QueryStringVaryEnabledInput() interface{}
	QueryStringVaryParameters() *[]*string
	SetQueryStringVaryParameters(val *[]*string)
	QueryStringVaryParametersInput() *[]*string
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
	ResetCacheKeyHeaders()
	ResetQueryStringVaryEnabled()
	ResetQueryStringVaryParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnDistributionConfigCacheConfigOutputReference
type jsiiProxy_CdnDistributionConfigCacheConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) CacheKeyHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheKeyHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) CacheKeyHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheKeyHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) QueryStringVaryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringVaryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) QueryStringVaryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringVaryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) QueryStringVaryParameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringVaryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) QueryStringVaryParametersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringVaryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnDistributionConfigCacheConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnDistributionConfigCacheConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCdnDistributionConfigCacheConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnDistributionConfigCacheConfigOutputReference{}

	_jsii_.Create(
		"stackit.cdnDistribution.CdnDistributionConfigCacheConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnDistributionConfigCacheConfigOutputReference_Override(c CdnDistributionConfigCacheConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.cdnDistribution.CdnDistributionConfigCacheConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference)SetCacheKeyHeaders(val *[]*string) {
	if err := j.validateSetCacheKeyHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheKeyHeaders",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference)SetQueryStringVaryEnabled(val interface{}) {
	if err := j.validateSetQueryStringVaryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringVaryEnabled",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference)SetQueryStringVaryParameters(val *[]*string) {
	if err := j.validateSetQueryStringVaryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringVaryParameters",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) ResetCacheKeyHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheKeyHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) ResetQueryStringVaryEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStringVaryEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) ResetQueryStringVaryParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStringVaryParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionConfigCacheConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

