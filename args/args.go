package args

import (
	"flag"
)

type ApplicationArgs interface {
	GetCluster() string

	GetPipeline() string
}

type applicationargs struct {
	cluster  string
	pipeline string
}

func NewAppArgs() ApplicationArgs {
	var cluster string
	var pipeline string
	flag.StringVar(&cluster, "cluster", "config/cluster", "The cluster yaml file must not be empty")
	flag.StringVar(&pipeline, "pipeline", "", "The pipeline yaml file must not be empty")
	return &applicationargs{
		cluster:  cluster,
		pipeline: pipeline,
	}
}

func (osargs applicationargs) GetCluster() string {
	return osargs.cluster
}

func (osargs applicationargs) GetPipeline() string {
	return osargs.pipeline
}
