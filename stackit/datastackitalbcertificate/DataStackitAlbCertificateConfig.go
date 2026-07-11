package datastackitalbcertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitAlbCertificateConfig struct {
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
	// The ID of the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/data-sources/alb_certificate#cert_id DataStackitAlbCertificate#cert_id}
	CertId *string `field:"required" json:"certId" yaml:"certId"`
	// STACKIT project ID to which the certificate is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.102.0/docs/data-sources/alb_certificate#project_id DataStackitAlbCertificate#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

