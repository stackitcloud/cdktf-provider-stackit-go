package serverserviceaccountattach

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServerServiceAccountAttachConfig struct {
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
	// STACKIT project ID to which the service account attachment is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.68.0/docs/resources/server_service_account_attach#project_id ServerServiceAccountAttach#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The server ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.68.0/docs/resources/server_service_account_attach#server_id ServerServiceAccountAttach#server_id}
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// The service account email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.68.0/docs/resources/server_service_account_attach#service_account_email ServerServiceAccountAttach#service_account_email}
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
}

