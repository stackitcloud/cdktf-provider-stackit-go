package datastackitlogmeinstance

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"stackit.dataStackitLogmeInstance.DataStackitLogmeInstance",
		reflect.TypeOf((*DataStackitLogmeInstance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cfGuid", GoGetter: "CfGuid"},
			_jsii_.MemberProperty{JsiiProperty: "cfOrganizationGuid", GoGetter: "CfOrganizationGuid"},
			_jsii_.MemberProperty{JsiiProperty: "cfSpaceGuid", GoGetter: "CfSpaceGuid"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dashboardUrl", GoGetter: "DashboardUrl"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "imageUrl", GoGetter: "ImageUrl"},
			_jsii_.MemberProperty{JsiiProperty: "instanceId", GoGetter: "InstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "instanceIdInput", GoGetter: "InstanceIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "planId", GoGetter: "PlanId"},
			_jsii_.MemberProperty{JsiiProperty: "planName", GoGetter: "PlanName"},
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "projectIdInput", GoGetter: "ProjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_DataStackitLogmeInstance{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"stackit.dataStackitLogmeInstance.DataStackitLogmeInstanceConfig",
		reflect.TypeOf((*DataStackitLogmeInstanceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"stackit.dataStackitLogmeInstance.DataStackitLogmeInstanceParameters",
		reflect.TypeOf((*DataStackitLogmeInstanceParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"stackit.dataStackitLogmeInstance.DataStackitLogmeInstanceParametersOutputReference",
		reflect.TypeOf((*DataStackitLogmeInstanceParametersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enableMonitoring", GoGetter: "EnableMonitoring"},
			_jsii_.MemberProperty{JsiiProperty: "fluentdTcp", GoGetter: "FluentdTcp"},
			_jsii_.MemberProperty{JsiiProperty: "fluentdTls", GoGetter: "FluentdTls"},
			_jsii_.MemberProperty{JsiiProperty: "fluentdTlsCiphers", GoGetter: "FluentdTlsCiphers"},
			_jsii_.MemberProperty{JsiiProperty: "fluentdTlsMaxVersion", GoGetter: "FluentdTlsMaxVersion"},
			_jsii_.MemberProperty{JsiiProperty: "fluentdTlsMinVersion", GoGetter: "FluentdTlsMinVersion"},
			_jsii_.MemberProperty{JsiiProperty: "fluentdTlsVersion", GoGetter: "FluentdTlsVersion"},
			_jsii_.MemberProperty{JsiiProperty: "fluentdUdp", GoGetter: "FluentdUdp"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "graphite", GoGetter: "Graphite"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ismDeletionAfter", GoGetter: "IsmDeletionAfter"},
			_jsii_.MemberProperty{JsiiProperty: "ismJitter", GoGetter: "IsmJitter"},
			_jsii_.MemberProperty{JsiiProperty: "ismJobInterval", GoGetter: "IsmJobInterval"},
			_jsii_.MemberProperty{JsiiProperty: "javaHeapspace", GoGetter: "JavaHeapspace"},
			_jsii_.MemberProperty{JsiiProperty: "javaMaxmetaspace", GoGetter: "JavaMaxmetaspace"},
			_jsii_.MemberProperty{JsiiProperty: "maxDiskThreshold", GoGetter: "MaxDiskThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "metricsFrequency", GoGetter: "MetricsFrequency"},
			_jsii_.MemberProperty{JsiiProperty: "metricsPrefix", GoGetter: "MetricsPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "monitoringInstanceId", GoGetter: "MonitoringInstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "opensearchTlsCiphers", GoGetter: "OpensearchTlsCiphers"},
			_jsii_.MemberProperty{JsiiProperty: "opensearchTlsProtocols", GoGetter: "OpensearchTlsProtocols"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sgwAcl", GoGetter: "SgwAcl"},
			_jsii_.MemberProperty{JsiiProperty: "syslog", GoGetter: "Syslog"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataStackitLogmeInstanceParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
