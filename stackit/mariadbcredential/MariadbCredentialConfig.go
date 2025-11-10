package mariadbcredential

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MariadbCredentialConfig struct {
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
	// ID of the MariaDB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.64.0/docs/resources/mariadb_credential#instance_id MariadbCredential#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// STACKIT Project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.64.0/docs/resources/mariadb_credential#project_id MariadbCredential#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

