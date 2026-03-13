package datastackitserviceaccounts

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitServiceAccountsConfig struct {
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
	// STACKIT project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/data-sources/service_accounts#project_id DataStackitServiceAccounts#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Optional regular expression to filter service accounts by email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/data-sources/service_accounts#email_regex DataStackitServiceAccounts#email_regex}
	EmailRegex *string `field:"optional" json:"emailRegex" yaml:"emailRegex"`
	// Optional suffix to filter service accounts by email (e.g.,`@sa.stackit.cloud`, `@ske.sa.stackit.cloud`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/data-sources/service_accounts#email_suffix DataStackitServiceAccounts#email_suffix}
	EmailSuffix *string `field:"optional" json:"emailSuffix" yaml:"emailSuffix"`
	// If set to `true`, service accounts are sorted in ascending lexicographical order by email. Defaults to `false` (descending).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.87.0/docs/data-sources/service_accounts#sort_ascending DataStackitServiceAccounts#sort_ascending}
	SortAscending interface{} `field:"optional" json:"sortAscending" yaml:"sortAscending"`
}

