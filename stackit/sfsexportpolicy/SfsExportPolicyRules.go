package sfsexportpolicy


type SfsExportPolicyRules struct {
	// IP access control list;
	//
	// IPs must have a subnet mask (e.g. "172.16.0.0/24" for a range of IPs, or "172.16.0.250/32" for a specific IP).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.76.0/docs/resources/sfs_export_policy#ip_acl SfsExportPolicy#ip_acl}
	IpAcl *[]*string `field:"required" json:"ipAcl" yaml:"ipAcl"`
	// Order of the rule within a Share Export Policy.
	//
	// The order is used so that when a client IP matches multiple rules, the first rule is applied
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.76.0/docs/resources/sfs_export_policy#order SfsExportPolicy#order}
	Order *float64 `field:"required" json:"order" yaml:"order"`
	// Description of the Rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.76.0/docs/resources/sfs_export_policy#description SfsExportPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Flag to indicate if client IPs matching this rule can only mount the share in read only mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.76.0/docs/resources/sfs_export_policy#read_only SfsExportPolicy#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Flag to honor set UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.76.0/docs/resources/sfs_export_policy#set_uuid SfsExportPolicy#set_uuid}
	SetUuid interface{} `field:"optional" json:"setUuid" yaml:"setUuid"`
	// Flag to indicate if client IPs matching this rule have root access on the Share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.76.0/docs/resources/sfs_export_policy#super_user SfsExportPolicy#super_user}
	SuperUser interface{} `field:"optional" json:"superUser" yaml:"superUser"`
}

