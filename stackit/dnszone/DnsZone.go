package dnszone

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/dnszone/internal"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.23.0/docs/resources/dns_zone stackit_dns_zone}.
type DnsZone interface {
	cdktf.TerraformResource
	Acl() *string
	SetAcl(val *string)
	AclInput() *string
	Active() interface{}
	SetActive(val interface{})
	ActiveInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContactEmail() *string
	SetContactEmail(val *string)
	ContactEmailInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DefaultTtl() *float64
	SetDefaultTtl(val *float64)
	DefaultTtlInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DnsName() *string
	SetDnsName(val *string)
	DnsNameInput() *string
	ExpireTime() *float64
	SetExpireTime(val *float64)
	ExpireTimeInput() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IsReverseZone() interface{}
	SetIsReverseZone(val interface{})
	IsReverseZoneInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NegativeCache() *float64
	SetNegativeCache(val *float64)
	NegativeCacheInput() *float64
	// The tree node.
	Node() constructs.Node
	Primaries() *[]*string
	SetPrimaries(val *[]*string)
	PrimariesInput() *[]*string
	PrimaryNameServer() *string
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RecordCount() *float64
	RefreshTime() *float64
	SetRefreshTime(val *float64)
	RefreshTimeInput() *float64
	RetryTime() *float64
	SetRetryTime(val *float64)
	RetryTimeInput() *float64
	SerialNumber() *float64
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Visibility() *string
	ZoneId() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAcl()
	ResetActive()
	ResetContactEmail()
	ResetDefaultTtl()
	ResetDescription()
	ResetExpireTime()
	ResetIsReverseZone()
	ResetNegativeCache()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrimaries()
	ResetRefreshTime()
	ResetRetryTime()
	ResetType()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DnsZone
type jsiiProxy_DnsZone struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DnsZone) Acl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) AclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) Active() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"active",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) ActiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) ContactEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) ContactEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) DefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) DefaultTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) DnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) ExpireTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expireTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) ExpireTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expireTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) IsReverseZone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isReverseZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) IsReverseZoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isReverseZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) NegativeCache() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"negativeCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) NegativeCacheInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"negativeCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) Primaries() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"primaries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) PrimariesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"primariesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) PrimaryNameServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryNameServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) RecordCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recordCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) RefreshTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) RefreshTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) RetryTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) RetryTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) SerialNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serialNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) Visibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZone) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.23.0/docs/resources/dns_zone stackit_dns_zone} Resource.
func NewDnsZone(scope constructs.Construct, id *string, config *DnsZoneConfig) DnsZone {
	_init_.Initialize()

	if err := validateNewDnsZoneParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DnsZone{}

	_jsii_.Create(
		"stackit.dnsZone.DnsZone",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.23.0/docs/resources/dns_zone stackit_dns_zone} Resource.
func NewDnsZone_Override(d DnsZone, scope constructs.Construct, id *string, config *DnsZoneConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.dnsZone.DnsZone",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DnsZone)SetAcl(val *string) {
	if err := j.validateSetAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acl",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetActive(val interface{}) {
	if err := j.validateSetActiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"active",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetContactEmail(val *string) {
	if err := j.validateSetContactEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactEmail",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetDefaultTtl(val *float64) {
	if err := j.validateSetDefaultTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTtl",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetDnsName(val *string) {
	if err := j.validateSetDnsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsName",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetExpireTime(val *float64) {
	if err := j.validateSetExpireTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expireTime",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetIsReverseZone(val interface{}) {
	if err := j.validateSetIsReverseZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isReverseZone",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetNegativeCache(val *float64) {
	if err := j.validateSetNegativeCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negativeCache",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetPrimaries(val *[]*string) {
	if err := j.validateSetPrimariesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaries",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetRefreshTime(val *float64) {
	if err := j.validateSetRefreshTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshTime",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetRetryTime(val *float64) {
	if err := j.validateSetRetryTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryTime",
		val,
	)
}

func (j *jsiiProxy_DnsZone)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a DnsZone resource upon running "cdktf plan <stack-name>".
func DnsZone_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDnsZone_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"stackit.dnsZone.DnsZone",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DnsZone_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDnsZone_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.dnsZone.DnsZone",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DnsZone_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDnsZone_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.dnsZone.DnsZone",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DnsZone_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDnsZone_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.dnsZone.DnsZone",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DnsZone_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"stackit.dnsZone.DnsZone",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DnsZone) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DnsZone) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DnsZone) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DnsZone) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DnsZone) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DnsZone) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DnsZone) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DnsZone) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DnsZone) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DnsZone) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DnsZone) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DnsZone) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZone) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DnsZone) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZone) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DnsZone) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DnsZone) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DnsZone) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DnsZone) ResetAcl() {
	_jsii_.InvokeVoid(
		d,
		"resetAcl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZone) ResetActive() {
	_jsii_.InvokeVoid(
		d,
		"resetActive",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZone) ResetContactEmail() {
	_jsii_.InvokeVoid(
		d,
		"resetContactEmail",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZone) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZone) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZone) ResetExpireTime() {
	_jsii_.InvokeVoid(
		d,
		"resetExpireTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZone) ResetIsReverseZone() {
	_jsii_.InvokeVoid(
		d,
		"resetIsReverseZone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZone) ResetNegativeCache() {
	_jsii_.InvokeVoid(
		d,
		"resetNegativeCache",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZone) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZone) ResetPrimaries() {
	_jsii_.InvokeVoid(
		d,
		"resetPrimaries",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZone) ResetRefreshTime() {
	_jsii_.InvokeVoid(
		d,
		"resetRefreshTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZone) ResetRetryTime() {
	_jsii_.InvokeVoid(
		d,
		"resetRetryTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZone) ResetType() {
	_jsii_.InvokeVoid(
		d,
		"resetType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZone) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZone) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZone) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZone) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZone) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZone) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

