package config

type ConfigManager interface {

	/*
		加载配置到类型
	*/
	Load(file string) (interface{}, error)
}
