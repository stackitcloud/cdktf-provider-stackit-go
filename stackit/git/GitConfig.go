package git

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GitConfig struct {
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
	// Unique name linked to the git instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/git#name Git#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// STACKIT project ID to which the git instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/git#project_id Git#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Restricted ACL for instance access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/git#acl Git#acl}
	Acl *[]*string `field:"optional" json:"acl" yaml:"acl"`
	// Instance flavor.
	//
	// If not provided, defaults to git-100. For a list of available flavors, refer to our API documentation: `https://docs.api.stackit.cloud/documentation/git/version/v1beta`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.81.0/docs/resources/git#flavor Git#flavor}
	Flavor *string `field:"optional" json:"flavor" yaml:"flavor"`
}

