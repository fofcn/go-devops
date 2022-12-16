package main

import (
	"errors"
	"io"
	"log"
	"os"
	"taskmanager/cluster"
	"taskmanager/executor"
	"taskmanager/pipeline"
)

type ApplicationContext struct {
	executor executor.PipelineExecutor
	cluster  cluster.ClusterManager
	pipeline pipeline.PipelineManager
	session  cluster.ClusterSessionManager
}

type ContextRecord struct {
	clusterName string
	nodeName    string
	stageName   string
	scriptPos   int
}

var contextRecordTable map[string]ContextRecord = map[string]ContextRecord{}

func (ac *ApplicationContext) Init(obj interface{}) error {
	ac.executor = executor.PipelineExecutor{}
	ac.cluster = cluster.ClusterManager{}
	ac.pipeline = pipeline.PipelineManager{}
	ac.session = cluster.ClusterSessionManager{}

	err := ac.cluster.Init(nil)
	if err != nil {
		return err
	}
	ac.pipeline.Init(obj.(string))
	ac.executor.Init()
	ac.session.Init(&ac.cluster)

	return nil
}

func (ac *ApplicationContext) Start() error {
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
			log.Println(script)

			// run command in difference cluster and node
			var scriptExec executor.Script
			scriptExec.Script = *script
			scriptExec.IO = ac

			for _, tag := range stageStepTags {
				scriptExec.Cluster = tag.Cluster

				for _, node := range tag.Node {
					scriptExec.NodeId = node

					err := ac.executor.Exec(ac.session, scriptExec)
					if err != nil {
						log.Println(err)
						continue
						// return err
					}
				}

			}
			scriptPos++
			script, err = ac.pipeline.GetNextScript(stage, scriptPos)
		}

	}

	return nil
}

func (ac *ApplicationContext) Shutdown() error {
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

func (ac *ApplicationContext) GetStdout() io.Writer {
	return os.Stdout
}

func (ac *ApplicationContext) GetStderr() io.Writer {
	return os.Stderr
}
