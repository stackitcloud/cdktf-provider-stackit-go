//go:build no_runtime_type_checking

package applicationloadbalancer

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationLoadBalancerTargetPoolsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancerTargetPoolsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancerTargetPoolsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancerTargetPoolsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancerTargetPoolsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancerTargetPoolsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancerTargetPoolsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApplicationLoadBalancerTargetPoolsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

