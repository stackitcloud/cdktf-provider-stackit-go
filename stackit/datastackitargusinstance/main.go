package datastackitargusinstance

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"stackit.dataStackitArgusInstance.DataStackitArgusInstance",
		reflect.TypeOf((*DataStackitArgusInstance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acl", GoGetter: "Acl"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alertingUrl", GoGetter: "AlertingUrl"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "grafanaInitialAdminPassword", GoGetter: "GrafanaInitialAdminPassword"},
			_jsii_.MemberProperty{JsiiProperty: "grafanaInitialAdminUser", GoGetter: "GrafanaInitialAdminUser"},
			_jsii_.MemberProperty{JsiiProperty: "grafanaPublicReadAccess", GoGetter: "GrafanaPublicReadAccess"},
			_jsii_.MemberProperty{JsiiProperty: "grafanaUrl", GoGetter: "GrafanaUrl"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "instanceId", GoGetter: "InstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "instanceIdInput", GoGetter: "InstanceIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isUpdatable", GoGetter: "IsUpdatable"},
			_jsii_.MemberProperty{JsiiProperty: "jaegerTracesUrl", GoGetter: "JaegerTracesUrl"},
			_jsii_.MemberProperty{JsiiProperty: "jaegerUiUrl", GoGetter: "JaegerUiUrl"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "logsPushUrl", GoGetter: "LogsPushUrl"},
			_jsii_.MemberProperty{JsiiProperty: "logsUrl", GoGetter: "LogsUrl"},
			_jsii_.MemberProperty{JsiiProperty: "metricsPushUrl", GoGetter: "MetricsPushUrl"},
			_jsii_.MemberProperty{JsiiProperty: "metricsRetentionDays", GoGetter: "MetricsRetentionDays"},
			_jsii_.MemberProperty{JsiiProperty: "metricsRetentionDays1HDownsampling", GoGetter: "MetricsRetentionDays1HDownsampling"},
			_jsii_.MemberProperty{JsiiProperty: "metricsRetentionDays5MDownsampling", GoGetter: "MetricsRetentionDays5MDownsampling"},
			_jsii_.MemberProperty{JsiiProperty: "metricsUrl", GoGetter: "MetricsUrl"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "otlpTracesUrl", GoGetter: "OtlpTracesUrl"},
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
			_jsii_.MemberProperty{JsiiProperty: "targetsUrl", GoGetter: "TargetsUrl"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "zipkinSpansUrl", GoGetter: "ZipkinSpansUrl"},
		},
		func() interface{} {
			j := jsiiProxy_DataStackitArgusInstance{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"stackit.dataStackitArgusInstance.DataStackitArgusInstanceConfig",
		reflect.TypeOf((*DataStackitArgusInstanceConfig)(nil)).Elem(),
	)
}
