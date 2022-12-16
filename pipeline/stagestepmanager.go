package pipeline

import "errors"

type StageStepManager struct {
	stageStepTable map[string]StageStep
}

type StepScript struct {
	script string
	cuspos int
}

func (ssm *StageStepManager) Init(pc *PipelineConfig) error {
	ssm.stageStepTable = make(map[string]StageStep)
	for _, stageStep := range pc.stageSteps {
		ssm.stageStepTable[stageStep.Stage] = stageStep
	}

	return nil
}

func (ssm *StageStepManager) GetStageStep(stage string) StageStep {
	return (ssm.stageStepTable[stage])
}

func (ssm *StageStepManager) GetStageStepNextScript(stageStep *StageStep, pos int) (*StepScript, error) {
	var script StepScript
	if len(stageStep.Script) != pos {
		script.script = stageStep.Script[pos]
		script.cuspos = pos + 1

		return &script, nil
	}

	return nil, errors.New("No scripts available.")
}
