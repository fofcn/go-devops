package main

import (
	"taskmanager/args"
	"taskmanager/zlog"
)

func main() {
	zlog.InitLogger()
	defer zlog.Logger.Sync()
	osargs := args.NewAppArgs()
	controller := NewController(&osargs)
	controller.Init("config/pipeline/setup-demo.yaml")
	controller.Start()
	controller.Shutdown()
}
