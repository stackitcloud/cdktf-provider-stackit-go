package loadbalancerobservabilitycredential

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerObservabilityCredentialConfig struct {
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
	// Observability credential name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.38.1/docs/resources/loadbalancer_observability_credential#display_name LoadbalancerObservabilityCredential#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The username for the observability service (e.g. Argus) where the logs/metrics will be pushed into.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.38.1/docs/resources/loadbalancer_observability_credential#password LoadbalancerObservabilityCredential#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// STACKIT project ID to which the load balancer observability credential is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.38.1/docs/resources/loadbalancer_observability_credential#project_id LoadbalancerObservabilityCredential#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The password for the observability service (e.g. Argus) where the logs/metrics will be pushed into.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.38.1/docs/resources/loadbalancer_observability_credential#username LoadbalancerObservabilityCredential#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

