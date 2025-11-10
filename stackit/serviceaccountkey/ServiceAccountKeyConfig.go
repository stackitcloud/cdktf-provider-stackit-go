package serviceaccountkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceAccountKeyConfig struct {
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
	// The STACKIT project ID associated with the service account key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.68.0/docs/resources/service_account_key#project_id ServiceAccountKey#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The email address associated with the service account, used for account identification and communication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.68.0/docs/resources/service_account_key#service_account_email ServiceAccountKey#service_account_email}
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// Specifies the public_key (RSA2048 key-pair). If not provided, a certificate from STACKIT will be used to generate a private_key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.68.0/docs/resources/service_account_key#public_key ServiceAccountKey#public_key}
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
	// A map of arbitrary key/value pairs designed to force key recreation when they change, facilitating key rotation based on external factors such as a changing timestamp.
	//
	// Modifying this map triggers the creation of a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.68.0/docs/resources/service_account_key#rotate_when_changed ServiceAccountKey#rotate_when_changed}
	RotateWhenChanged *map[string]*string `field:"optional" json:"rotateWhenChanged" yaml:"rotateWhenChanged"`
	// Specifies the key's validity duration in days. If left unspecified, the key is considered valid until it is deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.68.0/docs/resources/service_account_key#ttl_days ServiceAccountKey#ttl_days}
	TtlDays *float64 `field:"optional" json:"ttlDays" yaml:"ttlDays"`
}

