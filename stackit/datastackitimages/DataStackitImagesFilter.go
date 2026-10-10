package datastackitimages


type DataStackitImagesFilter struct {
	// Experimental: API label selector passed directly to the image-list endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/stackitcloud/stackit/0.119.0/docs/data-sources/images#label_selector DataStackitImages#label_selector}
	LabelSelector *string `field:"optional" json:"labelSelector" yaml:"labelSelector"`
}

