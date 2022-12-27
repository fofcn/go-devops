package pipeline

import (
	"errors"
	"io/ioutil"
	"log"
	"taskmanager/zlog"

	"github.com/goinggo/mapstructure"
	"gopkg.in/yaml.v2"
)

const variablesKey string = "variables"
const stageKey string = "stage"

var variables map[string]string = make(map[string]string, 256)
var stages []string
var stageSteps []StageStep

type PipelineConfig struct {
	variables  map[string]string
	stages     []string
	stageSteps []StageStep
}

func (pc *PipelineConfig) Init(obj interface{}) error {
	pc.Load(obj.(string))
	return nil
}

func (pc *PipelineConfig) Load(filename string) (interface{}, error) {
	pc.variables = make(map[string]string)
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		zlog.Logger.Fatalf("Read env setup yaml file error, %v", err)
		return nil, err
	}

	var config map[string]interface{}
	err = yaml.Unmarshal([]byte(buf), &config)
	if err != nil {
		zlog.Logger.Fatalf("Unmarshal yaml file error, %v", err)
		return nil, err
	}

	pc.loadVariables(config)
	pc.loadStages(config)
	pc.loadStageSteps(config)

	return config, nil
}

func (pc *PipelineConfig) GetVariable(varName string) (string, error) {
	if _, varVal := pc.variables[varName]; varVal {
		return varName, nil
	}

	return "", errors.New("Variable " + varName + " undefine.")
}

func (pc *PipelineConfig) loadVariables(config map[string]interface{}) {
	// load system environment variables
	// load user define variables
	// merge two variables
	// if user define variables exist in system environment variables, override the variable using user define variables

	// get variables block
	// convert block to string
	// yaml unmarshal to object
	// put the variable into environments

	vars := config[variablesKey]
	if vars == nil {
		return
	}
	for varName, varValue := range vars.(map[interface{}]interface{}) {
		variables[varName.(string)] = varValue.(string)
		pc.variables[varName.(string)] = varValue.(string)
	}

}

func (pc *PipelineConfig) loadStages(config map[string]interface{}) {
	stageArr := config[stageKey]

	stageI := stageArr.([]interface{})
	for _, stageStr := range stageI {
		stages = append(stages, stageStr.(string))
		pc.stages = append(stages, stageStr.(string))
	}
}

func (pc *PipelineConfig) loadStageSteps(config map[string]interface{}) error {
	for _, stage := range stages {
		val := config[stage]
		var stageStep StageStep
		err := mapstructure.Decode(val, &stageStep)
		if err != nil {
			log.Println(err)
			continue
		}

		stageSteps = append(stageSteps, stageStep)
		pc.stageSteps = append(pc.stageSteps, stageStep)
	}

	return nil
}
