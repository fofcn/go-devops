package main

import (
	"errors"
	"io"
	"log"
	"os"
	"taskmanager/args"
	"taskmanager/cluster"
	"taskmanager/executor"
	"taskmanager/pipeline"
	"taskmanager/zlog"
)

type ApplicationController struct {
	executor executor.PipelineExecutor
	cluster  cluster.ClusterManager
	pipeline pipeline.PipelineManager
	session  cluster.ClusterSessionManager
	appargs  *args.ApplicationArgs
}

type ContextRecord struct {
	clusterName string
	nodeName    string
	stageName   string
	scriptPos   int
}

var contextRecordTable map[string]ContextRecord = map[string]ContextRecord{}

func NewController(appargs *args.ApplicationArgs) ApplicationController {
	return ApplicationController{
		executor: executor.PipelineExecutor{},
		cluster:  cluster.ClusterManager{},
		pipeline: pipeline.PipelineManager{},
		session:  cluster.ClusterSessionManager{},
		appargs:  appargs,
	}
}

func (ac *ApplicationController) Init(obj interface{}) error {
	err := ac.cluster.Init(*ac.appargs)
	if err != nil {
		return err
	}

	ac.pipeline.Init(*ac.appargs)
	ac.executor.Init(ac.pipeline.GetVariableManager())
	ac.session.Init(&ac.cluster)

	return nil
}

func (ac *ApplicationController) Start() error {
	// get the order of the stages
	stages := ac.pipeline.GetStages()
	if len(stages) == 0 {
		return errors.New("no stage to run, please check your configuration")
	}

	// get specific stage
	for _, stage := range stages {

		var scriptPos int = 0

		// get the tag of the stage
		stageStepTags, err := ac.pipeline.GetStageStepTags(stage)
		if err != nil {
			return err
		}

		// get the script of the stage
		script, err := ac.pipeline.GetNextScript(stage, scriptPos)
		for err == nil {
			// replace the variable if essential

			// run command in difference cluster and node
			var scriptExec executor.Script
			scriptExec.Script = *script
			scriptExec.IO = ac

			for _, tag := range stageStepTags {
				scriptExec.Cluster = tag.Cluster

				for _, node := range tag.Node {
					scriptExec.NodeId = node
					zlog.Logger.Infof("Execute shell command: %v on cluster: %v Node: %v", *script,
						scriptExec.Cluster, scriptExec.NodeId)
					err := ac.executor.Exec(ac.session, scriptExec)
					if err != nil {
						log.Fatal(err)
						return err
					}
				}

			}
			scriptPos++
			script, err = ac.pipeline.GetNextScript(stage, scriptPos)
		}

	}

	return nil
}

func (ac *ApplicationController) Shutdown() error {
	err := ac.cluster.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
	err = ac.pipeline.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
	err = ac.session.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
	err = ac.executor.Shutdown()
	return err
}

func (ac *ApplicationController) GetStdout() io.Writer {
	return os.Stdout
}

func (ac *ApplicationController) GetStderr() io.Writer {
	return os.Stderr
}
