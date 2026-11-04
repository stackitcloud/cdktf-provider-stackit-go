package albwafmanagedruleset

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/albwafmanagedruleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbWafManagedRuleSetConfiguredGroupsRulesMap interface {
	cdktf.ComplexMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	Get(key *string) AlbWafManagedRuleSetConfiguredGroupsRulesOutputReference
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

// The jsii proxy struct for AlbWafManagedRuleSetConfiguredGroupsRulesMap
type jsiiProxy_AlbWafManagedRuleSetConfiguredGroupsRulesMap struct {
	internal.Type__cdktfComplexMap
}

func (j *jsiiProxy_AlbWafManagedRuleSetConfiguredGroupsRulesMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbWafManagedRuleSetConfiguredGroupsRulesMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbWafManagedRuleSetConfiguredGroupsRulesMap) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbWafManagedRuleSetConfiguredGroupsRulesMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbWafManagedRuleSetConfiguredGroupsRulesMap) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlbWafManagedRuleSetConfiguredGroupsRulesMap(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbWafManagedRuleSetConfiguredGroupsRulesMap {
	_init_.Initialize()

	if err := validateNewAlbWafManagedRuleSetConfiguredGroupsRulesMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlbWafManagedRuleSetConfiguredGroupsRulesMap{}

	_jsii_.Create(
		"stackit.albWafManagedRuleSet.AlbWafManagedRuleSetConfiguredGroupsRulesMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbWafManagedRuleSetConfiguredGroupsRulesMap_Override(a AlbWafManagedRuleSetConfiguredGroupsRulesMap, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.albWafManagedRuleSet.AlbWafManagedRuleSetConfiguredGroupsRulesMap",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbWafManagedRuleSetConfiguredGroupsRulesMap)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbWafManagedRuleSetConfiguredGroupsRulesMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbWafManagedRuleSetConfiguredGroupsRulesMap)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbWafManagedRuleSetConfiguredGroupsRulesMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbWafManagedRuleSetConfiguredGroupsRulesMap) Get(key *string) AlbWafManagedRuleSetConfiguredGroupsRulesOutputReference {
	if err := a.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns AlbWafManagedRuleSetConfiguredGroupsRulesOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbWafManagedRuleSetConfiguredGroupsRulesMap) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlbWafManagedRuleSetConfiguredGroupsRulesMap) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AlbWafManagedRuleSetConfiguredGroupsRulesMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

