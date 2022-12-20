package args

import (
	"flag"
	"os"
)

type ApplicationArgs interface {
	GetCluster() string

	GetPipeline() string
}

type applicationargs struct {
	cluster  string
	pipeline string
}

var requiredparameters = [...]string{"cluster", "pipeline"}

func NewAppArgs() ApplicationArgs {
	var cluster string
	var pipeline string
	flag.StringVar(&cluster, "cluster", "config/cluster", "Cluster configuration, see:")
	flag.StringVar(&pipeline, "pipeline", "", "Pipeline configuration, see:")
	flag.Parse()

	if cluster == "" || pipeline == "" {
		flag.Usage()
		os.Exit(1)
	}

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
