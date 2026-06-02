package albcertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbCertificateConfig struct {
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
	// Certificate name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/alb_certificate#name AlbCertificate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The PEM encoded private key part.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/alb_certificate#private_key AlbCertificate#private_key}
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// STACKIT project ID to which the certificate is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/alb_certificate#project_id AlbCertificate#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The PEM encoded public key part.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/alb_certificate#public_key AlbCertificate#public_key}
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
	// The resource region (e.g. eu01). If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.98.0/docs/resources/alb_certificate#region AlbCertificate#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

