package resourcemanagerproject


type ResourcemanagerProjectMembers struct {
	// The role of the member in the project.
	//
	// Possible values include, but are not limited to: `owner`, `editor`, `reader`. Legacy roles (`project.admin`, `project.auditor`, `project.member`, `project.owner`) are not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.34.0/docs/resources/resourcemanager_project#role ResourcemanagerProject#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Unique identifier of the user, service account or client.
	//
	// This is usually the email address for users or service accounts, and the name in case of clients.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.34.0/docs/resources/resourcemanager_project#subject ResourcemanagerProject#subject}
	Subject *string `field:"required" json:"subject" yaml:"subject"`
}

