package pipeline

type StageManager struct {
	pipelineconfig *PipelineConfig
}

func (sm *StageManager) Init(pipelineconfig *PipelineConfig) error {
	sm.pipelineconfig = pipelineconfig
	return nil
}

func (sm *StageManager) GetStages() []string {
	return sm.pipelineconfig.stages
}
