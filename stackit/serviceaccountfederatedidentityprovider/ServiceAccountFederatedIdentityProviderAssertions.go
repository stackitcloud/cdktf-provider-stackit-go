package serviceaccountfederatedidentityprovider


type ServiceAccountFederatedIdentityProviderAssertions struct {
	// The assertion claim. At least one assertion with the claim "aud" is required for security reasons.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.99.0/docs/resources/service_account_federated_identity_provider#item ServiceAccountFederatedIdentityProvider#item}
	Item *string `field:"required" json:"item" yaml:"item"`
	// The assertion operator. Currently, the only supported operator is "equals".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.99.0/docs/resources/service_account_federated_identity_provider#operator ServiceAccountFederatedIdentityProvider#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The assertion value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.99.0/docs/resources/service_account_federated_identity_provider#value ServiceAccountFederatedIdentityProvider#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

