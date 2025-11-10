package network

import (
	_init_ "github.com/stackitcloud/cdktf-provider-stackit-go/stackit/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/stackitcloud/cdktf-provider-stackit-go/stackit/network/internal"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/network stackit_network}.
type Network interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	Ipv4Gateway() *string
	SetIpv4Gateway(val *string)
	Ipv4GatewayInput() *string
	Ipv4Nameservers() *[]*string
	SetIpv4Nameservers(val *[]*string)
	Ipv4NameserversInput() *[]*string
	Ipv4Prefix() *string
	SetIpv4Prefix(val *string)
	Ipv4Prefixes() *[]*string
	Ipv4PrefixInput() *string
	Ipv4PrefixLength() *float64
	SetIpv4PrefixLength(val *float64)
	Ipv4PrefixLengthInput() *float64
	Ipv6Gateway() *string
	SetIpv6Gateway(val *string)
	Ipv6GatewayInput() *string
	Ipv6Nameservers() *[]*string
	SetIpv6Nameservers(val *[]*string)
	Ipv6NameserversInput() *[]*string
	Ipv6Prefix() *string
	SetIpv6Prefix(val *string)
	Ipv6Prefixes() *[]*string
	Ipv6PrefixInput() *string
	Ipv6PrefixLength() *float64
	SetIpv6PrefixLength(val *float64)
	Ipv6PrefixLengthInput() *float64
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nameservers() *[]*string
	SetNameservers(val *[]*string)
	NameserversInput() *[]*string
	NetworkId() *string
	// The tree node.
	Node() constructs.Node
	NoIpv4Gateway() interface{}
	SetNoIpv4Gateway(val interface{})
	NoIpv4GatewayInput() interface{}
	NoIpv6Gateway() interface{}
	SetNoIpv6Gateway(val interface{})
	NoIpv6GatewayInput() interface{}
	Prefixes() *[]*string
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
	PublicIp() *string
	// Experimental.
	RawOverrides() interface{}
	Routed() interface{}
	SetRouted(val interface{})
	RoutedInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetIpv4Gateway()
	ResetIpv4Nameservers()
	ResetIpv4Prefix()
	ResetIpv4PrefixLength()
	ResetIpv6Gateway()
	ResetIpv6Nameservers()
	ResetIpv6Prefix()
	ResetIpv6PrefixLength()
	ResetLabels()
	ResetNameservers()
	ResetNoIpv4Gateway()
	ResetNoIpv6Gateway()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRouted()
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

// The jsii proxy struct for Network
type jsiiProxy_Network struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Network) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Ipv4Gateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4Gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Ipv4GatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4GatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Ipv4Nameservers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv4Nameservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Ipv4NameserversInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv4NameserversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Ipv4Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4Prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Ipv4Prefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv4Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Ipv4PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4PrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Ipv4PrefixLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4PrefixLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Ipv4PrefixLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4PrefixLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Ipv6Gateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6Gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Ipv6GatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6GatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Ipv6Nameservers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6Nameservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Ipv6NameserversInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6NameserversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Ipv6Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6Prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Ipv6Prefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Ipv6PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6PrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Ipv6PrefixLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6PrefixLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Ipv6PrefixLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6PrefixLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Nameservers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nameservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) NameserversInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nameserversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) NoIpv4Gateway() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noIpv4Gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) NoIpv4GatewayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noIpv4GatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) NoIpv6Gateway() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noIpv6Gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) NoIpv6GatewayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noIpv6GatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Prefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) PublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) Routed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) RoutedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Network) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/network stackit_network} Resource.
func NewNetwork(scope constructs.Construct, id *string, config *NetworkConfig) Network {
	_init_.Initialize()

	if err := validateNewNetworkParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Network{}

	_jsii_.Create(
		"stackit.network.Network",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.49.1/docs/resources/network stackit_network} Resource.
func NewNetwork_Override(n Network, scope constructs.Construct, id *string, config *NetworkConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"stackit.network.Network",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_Network)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Network)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Network)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Network)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Network)SetIpv4Gateway(val *string) {
	if err := j.validateSetIpv4GatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Gateway",
		val,
	)
}

func (j *jsiiProxy_Network)SetIpv4Nameservers(val *[]*string) {
	if err := j.validateSetIpv4NameserversParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Nameservers",
		val,
	)
}

func (j *jsiiProxy_Network)SetIpv4Prefix(val *string) {
	if err := j.validateSetIpv4PrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Prefix",
		val,
	)
}

func (j *jsiiProxy_Network)SetIpv4PrefixLength(val *float64) {
	if err := j.validateSetIpv4PrefixLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4PrefixLength",
		val,
	)
}

func (j *jsiiProxy_Network)SetIpv6Gateway(val *string) {
	if err := j.validateSetIpv6GatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Gateway",
		val,
	)
}

func (j *jsiiProxy_Network)SetIpv6Nameservers(val *[]*string) {
	if err := j.validateSetIpv6NameserversParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Nameservers",
		val,
	)
}

func (j *jsiiProxy_Network)SetIpv6Prefix(val *string) {
	if err := j.validateSetIpv6PrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Prefix",
		val,
	)
}

func (j *jsiiProxy_Network)SetIpv6PrefixLength(val *float64) {
	if err := j.validateSetIpv6PrefixLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6PrefixLength",
		val,
	)
}

func (j *jsiiProxy_Network)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_Network)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Network)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Network)SetNameservers(val *[]*string) {
	if err := j.validateSetNameserversParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameservers",
		val,
	)
}

func (j *jsiiProxy_Network)SetNoIpv4Gateway(val interface{}) {
	if err := j.validateSetNoIpv4GatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noIpv4Gateway",
		val,
	)
}

func (j *jsiiProxy_Network)SetNoIpv6Gateway(val interface{}) {
	if err := j.validateSetNoIpv6GatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noIpv6Gateway",
		val,
	)
}

func (j *jsiiProxy_Network)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_Network)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Network)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Network)SetRouted(val interface{}) {
	if err := j.validateSetRoutedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routed",
		val,
	)
}

// Generates CDKTF code for importing a Network resource upon running "cdktf plan <stack-name>".
func Network_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNetwork_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"stackit.network.Network",
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
func Network_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetwork_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.network.Network",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Network_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetwork_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.network.Network",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Network_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetwork_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"stackit.network.Network",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Network_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"stackit.network.Network",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_Network) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_Network) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_Network) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Network) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Network) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Network) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Network) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Network) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Network) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Network) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Network) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Network) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Network) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_Network) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Network) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_Network) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_Network) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_Network) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_Network) ResetIpv4Gateway() {
	_jsii_.InvokeVoid(
		n,
		"resetIpv4Gateway",
		nil, // no parameters
	)
}

func (n *jsiiProxy_Network) ResetIpv4Nameservers() {
	_jsii_.InvokeVoid(
		n,
		"resetIpv4Nameservers",
		nil, // no parameters
	)
}

func (n *jsiiProxy_Network) ResetIpv4Prefix() {
	_jsii_.InvokeVoid(
		n,
		"resetIpv4Prefix",
		nil, // no parameters
	)
}

func (n *jsiiProxy_Network) ResetIpv4PrefixLength() {
	_jsii_.InvokeVoid(
		n,
		"resetIpv4PrefixLength",
		nil, // no parameters
	)
}

func (n *jsiiProxy_Network) ResetIpv6Gateway() {
	_jsii_.InvokeVoid(
		n,
		"resetIpv6Gateway",
		nil, // no parameters
	)
}

func (n *jsiiProxy_Network) ResetIpv6Nameservers() {
	_jsii_.InvokeVoid(
		n,
		"resetIpv6Nameservers",
		nil, // no parameters
	)
}

func (n *jsiiProxy_Network) ResetIpv6Prefix() {
	_jsii_.InvokeVoid(
		n,
		"resetIpv6Prefix",
		nil, // no parameters
	)
}

func (n *jsiiProxy_Network) ResetIpv6PrefixLength() {
	_jsii_.InvokeVoid(
		n,
		"resetIpv6PrefixLength",
		nil, // no parameters
	)
}

func (n *jsiiProxy_Network) ResetLabels() {
	_jsii_.InvokeVoid(
		n,
		"resetLabels",
		nil, // no parameters
	)
}

func (n *jsiiProxy_Network) ResetNameservers() {
	_jsii_.InvokeVoid(
		n,
		"resetNameservers",
		nil, // no parameters
	)
}

func (n *jsiiProxy_Network) ResetNoIpv4Gateway() {
	_jsii_.InvokeVoid(
		n,
		"resetNoIpv4Gateway",
		nil, // no parameters
	)
}

func (n *jsiiProxy_Network) ResetNoIpv6Gateway() {
	_jsii_.InvokeVoid(
		n,
		"resetNoIpv6Gateway",
		nil, // no parameters
	)
}

func (n *jsiiProxy_Network) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_Network) ResetRouted() {
	_jsii_.InvokeVoid(
		n,
		"resetRouted",
		nil, // no parameters
	)
}

func (n *jsiiProxy_Network) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Network) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Network) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Network) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Network) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Network) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

