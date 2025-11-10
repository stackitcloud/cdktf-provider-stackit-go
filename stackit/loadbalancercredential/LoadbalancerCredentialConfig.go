package loadbalancercredential

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerCredentialConfig struct {
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
	// Credential name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.2/docs/resources/loadbalancer_credential#display_name LoadbalancerCredential#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The password used for the ARGUS instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.2/docs/resources/loadbalancer_credential#password LoadbalancerCredential#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// STACKIT project ID to which the load balancer credential is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.2/docs/resources/loadbalancer_credential#project_id LoadbalancerCredential#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The username used for the ARGUS instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.33.2/docs/resources/loadbalancer_credential#username LoadbalancerCredential#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

