//go:build no_runtime_type_checking

package server

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Server) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_Server) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_Server) validatePutBootVolumeParameters(value *ServerBootVolume) error {
	return nil
}

func validateServer_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateServer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateServer_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateServer_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetAffinityGroupParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetAvailabilityZoneParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetDesiredStatusParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetImageIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetKeypairNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetLabelsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetMachineTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetProjectIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetUserDataParameters(val *string) error {
	return nil
}

func validateNewServerParameters(scope constructs.Construct, id *string, config *ServerConfig) error {
	return nil
}

