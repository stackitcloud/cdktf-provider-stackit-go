//go:build no_runtime_type_checking

package skecluster

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SkeClusterNodePoolsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SkeClusterNodePoolsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SkeClusterNodePoolsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SkeClusterNodePoolsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SkeClusterNodePoolsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SkeClusterNodePoolsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SkeClusterNodePoolsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSkeClusterNodePoolsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

