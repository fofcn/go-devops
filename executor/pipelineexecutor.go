package executor

import (
	"io"
	"taskmanager/cluster"
	executor "taskmanager/executor/param"
	"taskmanager/pipeline"
)

type Executor interface {
	Exec(Script Script) error
}

type IOStream interface {
	GetStdout() io.Writer
	GetStderr() io.Writer
}

type Script struct {
	Script  string
	Cluster string
	NodeId  string
	IO      IOStream
}

type PipelineExecutor struct {
	parambinder executor.ScriptParamBinder
}

func (pe *PipelineExecutor) Init(variableManager *pipeline.VariableManager) error {
	pe.parambinder = executor.NewParamBinder(variableManager)
	return nil
}

func (pe PipelineExecutor) Shutdown() error {
	return nil
}

func (pe PipelineExecutor) Exec(session cluster.ClusterSessionManager, script Script) error {
	execScript, err := pe.parambinder.Bind(&script.Script)
	if err != nil {
		return err
	}
	return session.RunCmd(script.Cluster, script.NodeId, execScript, script.IO.GetStdout(), script.IO.GetStderr())
}
