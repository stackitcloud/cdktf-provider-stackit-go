package cdndistribution

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/cdndistribution/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnDistributionConfigWafOutputReference interface {
	cdktf.ComplexObject
	AllowedHttpMethods() *[]*string
	SetAllowedHttpMethods(val *[]*string)
	AllowedHttpMethodsInput() *[]*string
	AllowedHttpVersions() *[]*string
	SetAllowedHttpVersions(val *[]*string)
	AllowedHttpVersionsInput() *[]*string
	AllowedRequestContentTypes() *[]*string
	SetAllowedRequestContentTypes(val *[]*string)
	AllowedRequestContentTypesInput() *[]*string
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
	DisabledRuleCollectionIds() *[]*string
	SetDisabledRuleCollectionIds(val *[]*string)
	DisabledRuleCollectionIdsInput() *[]*string
	DisabledRuleGroupIds() *[]*string
	SetDisabledRuleGroupIds(val *[]*string)
	DisabledRuleGroupIdsInput() *[]*string
	DisabledRuleIds() *[]*string
	SetDisabledRuleIds(val *[]*string)
	DisabledRuleIdsInput() *[]*string
	EnabledRuleCollectionIds() *[]*string
	SetEnabledRuleCollectionIds(val *[]*string)
	EnabledRuleCollectionIdsInput() *[]*string
	EnabledRuleGroupIds() *[]*string
	SetEnabledRuleGroupIds(val *[]*string)
	EnabledRuleGroupIdsInput() *[]*string
	EnabledRuleIds() *[]*string
	SetEnabledRuleIds(val *[]*string)
	EnabledRuleIdsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogOnlyRuleCollectionIds() *[]*string
	SetLogOnlyRuleCollectionIds(val *[]*string)
	LogOnlyRuleCollectionIdsInput() *[]*string
	LogOnlyRuleGroupIds() *[]*string
	SetLogOnlyRuleGroupIds(val *[]*string)
	LogOnlyRuleGroupIdsInput() *[]*string
	LogOnlyRuleIds() *[]*string
	SetLogOnlyRuleIds(val *[]*string)
	LogOnlyRuleIdsInput() *[]*string
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	ParanoiaLevel() *string
	SetParanoiaLevel(val *string)
	ParanoiaLevelInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetAllowedHttpMethods()
	ResetAllowedHttpVersions()
	ResetAllowedRequestContentTypes()
	ResetDisabledRuleCollectionIds()
	ResetDisabledRuleGroupIds()
	ResetDisabledRuleIds()
	ResetEnabledRuleCollectionIds()
	ResetEnabledRuleGroupIds()
	ResetEnabledRuleIds()
	ResetLogOnlyRuleCollectionIds()
	ResetLogOnlyRuleGroupIds()
	ResetLogOnlyRuleIds()
	ResetMode()
	ResetParanoiaLevel()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnDistributionConfigWafOutputReference
type jsiiProxy_CdnDistributionConfigWafOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) AllowedHttpMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHttpMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) AllowedHttpMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHttpMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) AllowedHttpVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHttpVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) AllowedHttpVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHttpVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) AllowedRequestContentTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRequestContentTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) AllowedRequestContentTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRequestContentTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) DisabledRuleCollectionIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledRuleCollectionIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) DisabledRuleCollectionIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledRuleCollectionIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) DisabledRuleGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledRuleGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) DisabledRuleGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledRuleGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) DisabledRuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledRuleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) DisabledRuleIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledRuleIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) EnabledRuleCollectionIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledRuleCollectionIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) EnabledRuleCollectionIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledRuleCollectionIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) EnabledRuleGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledRuleGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) EnabledRuleGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledRuleGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) EnabledRuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledRuleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) EnabledRuleIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledRuleIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) LogOnlyRuleCollectionIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logOnlyRuleCollectionIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) LogOnlyRuleCollectionIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logOnlyRuleCollectionIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) LogOnlyRuleGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logOnlyRuleGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) LogOnlyRuleGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logOnlyRuleGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) LogOnlyRuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logOnlyRuleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) LogOnlyRuleIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logOnlyRuleIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) ParanoiaLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paranoiaLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) ParanoiaLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paranoiaLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewCdnDistributionConfigWafOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnDistributionConfigWafOutputReference {
	_init_.Initialize()

	if err := validateNewCdnDistributionConfigWafOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnDistributionConfigWafOutputReference{}

	_jsii_.Create(
		"stackit.cdnDistribution.CdnDistributionConfigWafOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnDistributionConfigWafOutputReference_Override(c CdnDistributionConfigWafOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.cdnDistribution.CdnDistributionConfigWafOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetAllowedHttpMethods(val *[]*string) {
	if err := j.validateSetAllowedHttpMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedHttpMethods",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetAllowedHttpVersions(val *[]*string) {
	if err := j.validateSetAllowedHttpVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedHttpVersions",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetAllowedRequestContentTypes(val *[]*string) {
	if err := j.validateSetAllowedRequestContentTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedRequestContentTypes",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetDisabledRuleCollectionIds(val *[]*string) {
	if err := j.validateSetDisabledRuleCollectionIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabledRuleCollectionIds",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetDisabledRuleGroupIds(val *[]*string) {
	if err := j.validateSetDisabledRuleGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabledRuleGroupIds",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetDisabledRuleIds(val *[]*string) {
	if err := j.validateSetDisabledRuleIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabledRuleIds",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetEnabledRuleCollectionIds(val *[]*string) {
	if err := j.validateSetEnabledRuleCollectionIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledRuleCollectionIds",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetEnabledRuleGroupIds(val *[]*string) {
	if err := j.validateSetEnabledRuleGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledRuleGroupIds",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetEnabledRuleIds(val *[]*string) {
	if err := j.validateSetEnabledRuleIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledRuleIds",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetLogOnlyRuleCollectionIds(val *[]*string) {
	if err := j.validateSetLogOnlyRuleCollectionIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logOnlyRuleCollectionIds",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetLogOnlyRuleGroupIds(val *[]*string) {
	if err := j.validateSetLogOnlyRuleGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logOnlyRuleGroupIds",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetLogOnlyRuleIds(val *[]*string) {
	if err := j.validateSetLogOnlyRuleIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logOnlyRuleIds",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetParanoiaLevel(val *string) {
	if err := j.validateSetParanoiaLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paranoiaLevel",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnDistributionConfigWafOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) ResetAllowedHttpMethods() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedHttpMethods",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) ResetAllowedHttpVersions() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedHttpVersions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) ResetAllowedRequestContentTypes() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedRequestContentTypes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) ResetDisabledRuleCollectionIds() {
	_jsii_.InvokeVoid(
		c,
		"resetDisabledRuleCollectionIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) ResetDisabledRuleGroupIds() {
	_jsii_.InvokeVoid(
		c,
		"resetDisabledRuleGroupIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) ResetDisabledRuleIds() {
	_jsii_.InvokeVoid(
		c,
		"resetDisabledRuleIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) ResetEnabledRuleCollectionIds() {
	_jsii_.InvokeVoid(
		c,
		"resetEnabledRuleCollectionIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) ResetEnabledRuleGroupIds() {
	_jsii_.InvokeVoid(
		c,
		"resetEnabledRuleGroupIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) ResetEnabledRuleIds() {
	_jsii_.InvokeVoid(
		c,
		"resetEnabledRuleIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) ResetLogOnlyRuleCollectionIds() {
	_jsii_.InvokeVoid(
		c,
		"resetLogOnlyRuleCollectionIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) ResetLogOnlyRuleGroupIds() {
	_jsii_.InvokeVoid(
		c,
		"resetLogOnlyRuleGroupIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) ResetLogOnlyRuleIds() {
	_jsii_.InvokeVoid(
		c,
		"resetLogOnlyRuleIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		c,
		"resetMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) ResetParanoiaLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetParanoiaLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		c,
		"resetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CdnDistributionConfigWafOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

