//go:build no_runtime_type_checking

package applicationloadbalancer

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationLoadBalancerErrorsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancerErrorsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancerErrorsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancerErrorsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancerErrorsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancerErrorsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApplicationLoadBalancerErrorsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

