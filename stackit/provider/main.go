package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"stackit.provider.StackitProvider",
		reflect.TypeOf((*StackitProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "argusCustomEndpoint", GoGetter: "ArgusCustomEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "argusCustomEndpointInput", GoGetter: "ArgusCustomEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsPath", GoGetter: "CredentialsPath"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsPathInput", GoGetter: "CredentialsPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "dnsCustomEndpoint", GoGetter: "DnsCustomEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "dnsCustomEndpointInput", GoGetter: "DnsCustomEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "logmeCustomEndpoint", GoGetter: "LogmeCustomEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "logmeCustomEndpointInput", GoGetter: "LogmeCustomEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "mariadbCustomEndpoint", GoGetter: "MariadbCustomEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "mariadbCustomEndpointInput", GoGetter: "MariadbCustomEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "opensearchCustomEndpoint", GoGetter: "OpensearchCustomEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "opensearchCustomEndpointInput", GoGetter: "OpensearchCustomEndpointInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "postgresflexCustomEndpoint", GoGetter: "PostgresflexCustomEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "postgresflexCustomEndpointInput", GoGetter: "PostgresflexCustomEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "postgresqlCustomEndpoint", GoGetter: "PostgresqlCustomEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "postgresqlCustomEndpointInput", GoGetter: "PostgresqlCustomEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "rabbitmqCustomEndpoint", GoGetter: "RabbitmqCustomEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "rabbitmqCustomEndpointInput", GoGetter: "RabbitmqCustomEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "redisCustomEndpoint", GoGetter: "RedisCustomEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "redisCustomEndpointInput", GoGetter: "RedisCustomEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetArgusCustomEndpoint", GoMethod: "ResetArgusCustomEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetCredentialsPath", GoMethod: "ResetCredentialsPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetDnsCustomEndpoint", GoMethod: "ResetDnsCustomEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogmeCustomEndpoint", GoMethod: "ResetLogmeCustomEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetMariadbCustomEndpoint", GoMethod: "ResetMariadbCustomEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetOpensearchCustomEndpoint", GoMethod: "ResetOpensearchCustomEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPostgresflexCustomEndpoint", GoMethod: "ResetPostgresflexCustomEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetPostgresqlCustomEndpoint", GoMethod: "ResetPostgresqlCustomEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetRabbitmqCustomEndpoint", GoMethod: "ResetRabbitmqCustomEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedisCustomEndpoint", GoMethod: "ResetRedisCustomEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourcemanagerCustomEndpoint", GoMethod: "ResetResourcemanagerCustomEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceAccountEmail", GoMethod: "ResetServiceAccountEmail"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceAccountToken", GoMethod: "ResetServiceAccountToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkeCustomEndpoint", GoMethod: "ResetSkeCustomEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "resourcemanagerCustomEndpoint", GoGetter: "ResourcemanagerCustomEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "resourcemanagerCustomEndpointInput", GoGetter: "ResourcemanagerCustomEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountEmail", GoGetter: "ServiceAccountEmail"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountEmailInput", GoGetter: "ServiceAccountEmailInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountToken", GoGetter: "ServiceAccountToken"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountTokenInput", GoGetter: "ServiceAccountTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "skeCustomEndpoint", GoGetter: "SkeCustomEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "skeCustomEndpointInput", GoGetter: "SkeCustomEndpointInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_StackitProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"stackit.provider.StackitProviderConfig",
		reflect.TypeOf((*StackitProviderConfig)(nil)).Elem(),
	)
}
