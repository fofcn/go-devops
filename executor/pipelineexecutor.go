package executor

import (
	"io"
	"taskmanager/cluster"
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
}

func (pe PipelineExecutor) Init() error {
	return nil
}

func (pe PipelineExecutor) Shutdown() error {
	return nil
}

func (pe PipelineExecutor) Exec(session cluster.ClusterSessionManager, script Script) error {
	return session.RunCmd(script.Cluster, script.NodeId, script.Script, script.IO.GetStdout(), script.IO.GetStderr())

}
