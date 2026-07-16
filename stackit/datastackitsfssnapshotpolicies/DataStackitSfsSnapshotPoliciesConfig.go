package datastackitsfssnapshotpolicies

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataStackitSfsSnapshotPoliciesConfig struct {
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
	// STACKIT project ID to which the snapshot policy is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/data-sources/sfs_snapshot_policies#project_id DataStackitSfsSnapshotPolicies#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Filter snapshot policies by immutability.
	//
	// Possible values are: `all`, `immutable-only`, `mutable-only`. Defaults to `all`. This attribute is in beta, may have breaking changes in the future.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.103.0/docs/data-sources/sfs_snapshot_policies#immutable DataStackitSfsSnapshotPolicies#immutable}
	Immutable *string `field:"optional" json:"immutable" yaml:"immutable"`
}

