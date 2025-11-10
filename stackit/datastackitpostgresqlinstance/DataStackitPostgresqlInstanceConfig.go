package datastackitpostgresqlinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitPostgresqlInstanceConfig struct {
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
	// ID of the PostgreSQL instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.25.1/docs/data-sources/postgresql_instance#instance_id DataStackitPostgresqlInstance#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// STACKIT Project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.25.1/docs/data-sources/postgresql_instance#project_id DataStackitPostgresqlInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

