package datastackitsfssnapshotpolicies

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/datastackitsfssnapshotpolicies/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList
type jsiiProxy_DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList {
	_init_.Initialize()

	if err := validateNewDataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList{}

	_jsii_.Create(
		"stackit.dataStackitSfsSnapshotPolicies.DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList_Override(d DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.dataStackitSfsSnapshotPolicies.DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList) Get(index *float64) DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataStackitSfsSnapshotPoliciesItemsSnapshotSchedulesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

