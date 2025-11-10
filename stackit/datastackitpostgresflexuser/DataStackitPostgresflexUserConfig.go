package datastackitpostgresflexuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitPostgresflexUserConfig struct {
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
	// ID of the PostgresFlex instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.48.0/docs/data-sources/postgresflex_user#instance_id DataStackitPostgresflexUser#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.48.0/docs/data-sources/postgresflex_user#project_id DataStackitPostgresflexUser#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// User ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.48.0/docs/data-sources/postgresflex_user#user_id DataStackitPostgresflexUser#user_id}
	UserId *string `field:"required" json:"userId" yaml:"userId"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.48.0/docs/data-sources/postgresflex_user#region DataStackitPostgresflexUser#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

