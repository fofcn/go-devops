package pipeline

type PipelineManager struct {
	pipelineconfig   PipelineConfig
	variablemanager  VariableManager
	stagemanager     StageManager
	stagestepmanager StageStepManager
}

type ScriptClusterInfo struct {
	ClusterName string
	NodeId      string
	StageStep   StageStep
}

type ScriptContext struct {
	Script      string
	Stage       string
	cusPos      int
	ClusterInfo []ScriptClusterInfo
}

func (pm *PipelineManager) Init(p interface{}) error {
	pm.pipelineconfig = PipelineConfig{}
	pm.variablemanager = VariableManager{}
	pm.stagemanager = StageManager{}
	pm.stagestepmanager = StageStepManager{}

	pm.pipelineconfig.Init("config/pipeline/setup-demo.yaml")
	pm.variablemanager.Init()
	pm.stagemanager.Init(&pm.pipelineconfig)
	pm.stagestepmanager.Init(&pm.pipelineconfig)

	return nil
}

func (pm *PipelineManager) Start() error {
	return nil
}

func (pm *PipelineManager) Shutdown() error {
	return nil
}

func (pm *PipelineManager) GetNextScript(stage string, pos int) (*string, error) {
	stagestep := pm.stagestepmanager.GetStageStep(stage)
	script, err := pm.stagestepmanager.GetStageStepNextScript(&stagestep, pos)
	if err != nil {
		return nil, err
	}

	return &script.script, nil
}

func (pm *PipelineManager) GetStages() []string {
	return pm.stagemanager.GetStages()
}

func (pm *PipelineManager) GetStageStepTags(stage string) ([]StageTag, error) {
	stageStep := pm.stagestepmanager.GetStageStep(stage)
	// if stageStep == nil {
	// 	return nil, errors.New("Stage " + stage + " does not exists.")
	// }

	return stageStep.Tag, nil
}
