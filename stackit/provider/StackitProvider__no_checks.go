//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StackitProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_StackitProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateStackitProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateStackitProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStackitProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateStackitProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_StackitProvider) validateSetEnableBetaResourcesParameters(val interface{}) error {
	return nil
}

func validateNewStackitProviderParameters(scope constructs.Construct, id *string, config *StackitProviderConfig) error {
	return nil
}

