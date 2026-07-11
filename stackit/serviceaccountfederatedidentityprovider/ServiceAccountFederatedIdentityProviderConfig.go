package serviceaccountfederatedidentityprovider

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceAccountFederatedIdentityProviderConfig struct {
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
	// The assertions for the federated identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/service_account_federated_identity_provider#assertions ServiceAccountFederatedIdentityProvider#assertions}
	Assertions interface{} `field:"required" json:"assertions" yaml:"assertions"`
	// The issuer URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/service_account_federated_identity_provider#issuer ServiceAccountFederatedIdentityProvider#issuer}
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// The name of the federated identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/service_account_federated_identity_provider#name ServiceAccountFederatedIdentityProvider#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The STACKIT project ID associated with the service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/service_account_federated_identity_provider#project_id ServiceAccountFederatedIdentityProvider#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The email address associated with the service account, used for account identification and communication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/resources/service_account_federated_identity_provider#service_account_email ServiceAccountFederatedIdentityProvider#service_account_email}
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
}

