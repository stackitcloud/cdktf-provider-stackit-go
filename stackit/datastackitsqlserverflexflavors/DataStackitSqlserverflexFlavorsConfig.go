package datastackitsqlserverflexflavors

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitSqlserverflexFlavorsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.113.0/docs/data-sources/sqlserverflex_flavors#project_id DataStackitSqlserverflexFlavors#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// SqlserverFlex flavors data source region. If undefined the providers region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.113.0/docs/data-sources/sqlserverflex_flavors#region DataStackitSqlserverflexFlavors#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.113.0/docs/data-sources/sqlserverflex_flavors#timeouts DataStackitSqlserverflexFlavors#timeouts}.
	Timeouts *DataStackitSqlserverflexFlavorsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

