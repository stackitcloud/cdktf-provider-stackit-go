package mongodbflexinstance

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/mongodbflexinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MongodbflexInstanceOptionsOutputReference interface {
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
	DailySnapshotRetentionDays() *float64
	SetDailySnapshotRetentionDays(val *float64)
	DailySnapshotRetentionDaysInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MonthlySnapshotRetentionMonths() *float64
	SetMonthlySnapshotRetentionMonths(val *float64)
	MonthlySnapshotRetentionMonthsInput() *float64
	PointInTimeWindowHours() *float64
	SetPointInTimeWindowHours(val *float64)
	PointInTimeWindowHoursInput() *float64
	SnapshotRetentionDays() *float64
	SetSnapshotRetentionDays(val *float64)
	SnapshotRetentionDaysInput() *float64
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
	WeeklySnapshotRetentionWeeks() *float64
	SetWeeklySnapshotRetentionWeeks(val *float64)
	WeeklySnapshotRetentionWeeksInput() *float64
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
	ResetDailySnapshotRetentionDays()
	ResetMonthlySnapshotRetentionMonths()
	ResetPointInTimeWindowHours()
	ResetSnapshotRetentionDays()
	ResetWeeklySnapshotRetentionWeeks()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MongodbflexInstanceOptionsOutputReference
type jsiiProxy_MongodbflexInstanceOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference) DailySnapshotRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailySnapshotRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference) DailySnapshotRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailySnapshotRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference) MonthlySnapshotRetentionMonths() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monthlySnapshotRetentionMonths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference) MonthlySnapshotRetentionMonthsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monthlySnapshotRetentionMonthsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference) PointInTimeWindowHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pointInTimeWindowHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference) PointInTimeWindowHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pointInTimeWindowHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference) SnapshotRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference) SnapshotRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference) WeeklySnapshotRetentionWeeks() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weeklySnapshotRetentionWeeks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference) WeeklySnapshotRetentionWeeksInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weeklySnapshotRetentionWeeksInput",
		&returns,
	)
	return returns
}


func NewMongodbflexInstanceOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MongodbflexInstanceOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewMongodbflexInstanceOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MongodbflexInstanceOptionsOutputReference{}

	_jsii_.Create(
		"stackit.mongodbflexInstance.MongodbflexInstanceOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMongodbflexInstanceOptionsOutputReference_Override(m MongodbflexInstanceOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.mongodbflexInstance.MongodbflexInstanceOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference)SetDailySnapshotRetentionDays(val *float64) {
	if err := j.validateSetDailySnapshotRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailySnapshotRetentionDays",
		val,
	)
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference)SetMonthlySnapshotRetentionMonths(val *float64) {
	if err := j.validateSetMonthlySnapshotRetentionMonthsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monthlySnapshotRetentionMonths",
		val,
	)
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference)SetPointInTimeWindowHours(val *float64) {
	if err := j.validateSetPointInTimeWindowHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pointInTimeWindowHours",
		val,
	)
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference)SetSnapshotRetentionDays(val *float64) {
	if err := j.validateSetSnapshotRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotRetentionDays",
		val,
	)
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_MongodbflexInstanceOptionsOutputReference)SetWeeklySnapshotRetentionWeeks(val *float64) {
	if err := j.validateSetWeeklySnapshotRetentionWeeksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeklySnapshotRetentionWeeks",
		val,
	)
}

func (m *jsiiProxy_MongodbflexInstanceOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongodbflexInstanceOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongodbflexInstanceOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongodbflexInstanceOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongodbflexInstanceOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongodbflexInstanceOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongodbflexInstanceOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongodbflexInstanceOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongodbflexInstanceOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongodbflexInstanceOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongodbflexInstanceOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongodbflexInstanceOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongodbflexInstanceOptionsOutputReference) ResetDailySnapshotRetentionDays() {
	_jsii_.InvokeVoid(
		m,
		"resetDailySnapshotRetentionDays",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongodbflexInstanceOptionsOutputReference) ResetMonthlySnapshotRetentionMonths() {
	_jsii_.InvokeVoid(
		m,
		"resetMonthlySnapshotRetentionMonths",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongodbflexInstanceOptionsOutputReference) ResetPointInTimeWindowHours() {
	_jsii_.InvokeVoid(
		m,
		"resetPointInTimeWindowHours",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongodbflexInstanceOptionsOutputReference) ResetSnapshotRetentionDays() {
	_jsii_.InvokeVoid(
		m,
		"resetSnapshotRetentionDays",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongodbflexInstanceOptionsOutputReference) ResetWeeklySnapshotRetentionWeeks() {
	_jsii_.InvokeVoid(
		m,
		"resetWeeklySnapshotRetentionWeeks",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongodbflexInstanceOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongodbflexInstanceOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

