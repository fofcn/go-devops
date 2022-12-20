package pipeline

type pipeline struct {
	Variables map[interface{}]interface{} `yaml:"variables"`
	Stage     []string                    `yaml:"stage"`
}

type StageTag struct {
	Cluster string
	Node    []string
}

type StageStep struct {
	Stage  string
	Script []string
	Tag    []StageTag
}
