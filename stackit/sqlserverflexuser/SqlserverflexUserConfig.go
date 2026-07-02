package sqlserverflexuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SqlserverflexUserConfig struct {
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
	// ID of the SQLServer Flex instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/sqlserverflex_user#instance_id SqlserverflexUser#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/sqlserverflex_user#project_id SqlserverflexUser#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Database access levels for the user. The values for the default roles are: `##STACKIT_DatabaseManager##`, `##STACKIT_LoginManager##`, `##STACKIT_ProcessManager##`, `##STACKIT_ServerManager##`, `##STACKIT_SQLAgentManager##`, `##STACKIT_SQLAgentUser##`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/sqlserverflex_user#roles SqlserverflexUser#roles}
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
	// Username of the SQLServer Flex instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/sqlserverflex_user#username SqlserverflexUser#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/sqlserverflex_user#region SqlserverflexUser#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// A map of arbitrary key/value pairs that will force recreation of the resource when they change, enabling resource rotation based on external conditions such as a rotating timestamp.
	//
	// Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.101.0/docs/resources/sqlserverflex_user#rotate_when_changed SqlserverflexUser#rotate_when_changed}
	RotateWhenChanged *map[string]*string `field:"optional" json:"rotateWhenChanged" yaml:"rotateWhenChanged"`
}

