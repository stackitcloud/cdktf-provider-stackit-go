//go:build no_runtime_type_checking

package applicationloadbalancer

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationLoadBalancerNetworksList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancerNetworksList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancerNetworksList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancerNetworksList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancerNetworksList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancerNetworksList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApplicationLoadBalancerNetworksList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApplicationLoadBalancerNetworksListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

