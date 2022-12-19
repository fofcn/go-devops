package main

import "taskmanager/args"

func main() {
	osargs := args.NewAppArgs()
	controller := NewController(&osargs)
	controller.Init("config/pipeline/setup-demo.yaml")
	controller.Start()
	controller.Shutdown()
}
