package mongodbflexinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MongodbflexInstanceConfig struct {
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
	// The Access Control List (ACL) for the MongoDB Flex instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.7.0/docs/resources/mongodbflex_instance#acl MongodbflexInstance#acl}
	Acl *[]*string `field:"required" json:"acl" yaml:"acl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.7.0/docs/resources/mongodbflex_instance#backup_schedule MongodbflexInstance#backup_schedule}.
	BackupSchedule *string `field:"required" json:"backupSchedule" yaml:"backupSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.7.0/docs/resources/mongodbflex_instance#flavor MongodbflexInstance#flavor}.
	Flavor *MongodbflexInstanceFlavor `field:"required" json:"flavor" yaml:"flavor"`
	// Instance name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.7.0/docs/resources/mongodbflex_instance#name MongodbflexInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.7.0/docs/resources/mongodbflex_instance#options MongodbflexInstance#options}.
	Options *MongodbflexInstanceOptions `field:"required" json:"options" yaml:"options"`
	// STACKIT project ID to which the instance is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.7.0/docs/resources/mongodbflex_instance#project_id MongodbflexInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.7.0/docs/resources/mongodbflex_instance#replicas MongodbflexInstance#replicas}.
	Replicas *float64 `field:"required" json:"replicas" yaml:"replicas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.7.0/docs/resources/mongodbflex_instance#storage MongodbflexInstance#storage}.
	Storage *MongodbflexInstanceStorage `field:"required" json:"storage" yaml:"storage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.7.0/docs/resources/mongodbflex_instance#version MongodbflexInstance#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
}

