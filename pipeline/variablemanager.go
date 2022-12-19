package pipeline

import (
	"os"
)

type VariableManager struct {
	pm *PipelineManager
}

func New(pm *PipelineManager) *VariableManager {
	var vm *VariableManager
	vm.pm = pm
	return vm
}

func (vm *VariableManager) Init() error {
	return nil
}

func (vm *VariableManager) GetVariable(varName string) string {
	variable, err := vm.pm.pipelineconfig.GetVariable(varName)
	if err != nil {
		variable = os.Getenv(varName)
	}

	return variable
}
