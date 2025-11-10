package datastackitmariadbcredential

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitMariadbCredentialConfig struct {
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
	// The credential's ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.39.0/docs/data-sources/mariadb_credential#credential_id DataStackitMariadbCredential#credential_id}
	CredentialId *string `field:"required" json:"credentialId" yaml:"credentialId"`
	// ID of the MariaDB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.39.0/docs/data-sources/mariadb_credential#instance_id DataStackitMariadbCredential#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.39.0/docs/data-sources/mariadb_credential#project_id DataStackitMariadbCredential#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

