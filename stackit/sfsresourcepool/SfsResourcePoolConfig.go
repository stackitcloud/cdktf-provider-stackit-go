package sfsresourcepool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SfsResourcePoolConfig struct {
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
	// Availability zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/sfs_resource_pool#availability_zone SfsResourcePool#availability_zone}
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// List of IPs that can mount the resource pool in read-only;
	//
	// IPs must have a subnet mask (e.g. "172.16.0.0/24" for a range of IPs, or "172.16.0.250/32" for a specific IP).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/sfs_resource_pool#ip_acl SfsResourcePool#ip_acl}
	IpAcl *[]*string `field:"required" json:"ipAcl" yaml:"ipAcl"`
	// Name of the resource pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/sfs_resource_pool#name SfsResourcePool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the performance class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/sfs_resource_pool#performance_class SfsResourcePool#performance_class}
	PerformanceClass *string `field:"required" json:"performanceClass" yaml:"performanceClass"`
	// STACKIT project ID to which the resource pool is associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/sfs_resource_pool#project_id SfsResourcePool#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Size of the resource pool (unit: gigabytes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/sfs_resource_pool#size_gigabytes SfsResourcePool#size_gigabytes}
	SizeGigabytes *float64 `field:"required" json:"sizeGigabytes" yaml:"sizeGigabytes"`
	// The resource region. If not defined, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/sfs_resource_pool#region SfsResourcePool#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// If set to true, snapshots are visible and accessible to users. (default: false).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.79.0/docs/resources/sfs_resource_pool#snapshots_are_visible SfsResourcePool#snapshots_are_visible}
	SnapshotsAreVisible interface{} `field:"optional" json:"snapshotsAreVisible" yaml:"snapshotsAreVisible"`
}

