package datastackitserviceaccountfederatedidentityprovider

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitServiceAccountFederatedIdentityProviderConfig struct {
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
	// The unique identifier for the federated identity provider associated with the service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/data-sources/service_account_federated_identity_provider#federation_id DataStackitServiceAccountFederatedIdentityProvider#federation_id}
	FederationId *string `field:"required" json:"federationId" yaml:"federationId"`
	// The STACKIT project ID associated with the service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/data-sources/service_account_federated_identity_provider#project_id DataStackitServiceAccountFederatedIdentityProvider#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The email address associated with the service account, used for account identification and communication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.100.0/docs/data-sources/service_account_federated_identity_provider#service_account_email DataStackitServiceAccountFederatedIdentityProvider#service_account_email}
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
}

