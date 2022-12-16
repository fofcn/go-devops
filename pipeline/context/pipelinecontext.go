package pipeline

import "taskmanager/pipeline"

// type PipelineContext struct {
// 	// 集群
// 	Cluster string
// 	// 节点
// 	Node string
// 	// 阶段
// 	Stage string
// 	// 脚本
// 	Script string
// 	// 脚本位置（索引）
// 	pos int
// }

/*
Step Job 脚本上下文
*/
type PipelineContext struct {
	pipelinemanager pipeline.PipelineManager
}

type scriptpos struct {
	// 阶段
	Stage string
	// 脚本
	Script string
	// 脚本位置（索引）
	pos int
}

var scriptPosTable map[string]scriptpos = make(map[string]scriptpos)

func CreateScriptContext() *PipelineContext {
	var context *PipelineContext
	return context
}

/*
获取脚本
*/
func (sc PipelineContext) GetScript(stage string, pos int) (*string, error) {

	stepscript, err := sc.pipelinemanager.GetNextScript(stage, pos)
	if err != nil {
		return nil, err
	}

	return stepscript, nil
}
