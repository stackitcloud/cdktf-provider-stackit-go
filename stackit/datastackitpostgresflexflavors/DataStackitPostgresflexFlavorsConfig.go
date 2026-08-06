package datastackitpostgresflexflavors

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitPostgresflexFlavorsConfig struct {
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
	// STACKIT project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.109.0/docs/data-sources/postgresflex_flavors#project_id DataStackitPostgresflexFlavors#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Postgres Flex flavors data source region. If undefined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.109.0/docs/data-sources/postgresflex_flavors#region DataStackitPostgresflexFlavors#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.109.0/docs/data-sources/postgresflex_flavors#timeouts DataStackitPostgresflexFlavors#timeouts}.
	Timeouts *DataStackitPostgresflexFlavorsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

