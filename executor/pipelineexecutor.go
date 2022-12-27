package executor

import (
	"io"
	"strings"
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
	err = session.RunCmd(script.Cluster, script.NodeId, execScript, script.IO.GetStdout(), script.IO.GetStderr())
	if err != nil {
		if ignoreErr := pe.ingoreError(script); ignoreErr {
			return nil
		}
	}

	return err
}

func (pe PipelineExecutor) ingoreError(script Script) bool {
	var ignoreErrCmds = [128]string{
		"setenforce",
		"https://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | apt-key add -",
	}

	for _, cmd := range ignoreErrCmds {
		if strings.Contains(
			strings.ToLower(script.Script),
			strings.ToLower(cmd),
		) {
			return true
		}
	}

	return false
}
