package main

import (
	"os"
	"taskmanager/args"
	"taskmanager/zlog"
)

func main() {
	zlog.InitLogger()
	defer zlog.Logger.Sync()
	osargs := args.NewAppArgs()
	controller := NewController(&osargs)
	err := controller.Init()
	if err != nil {
		zlog.Logger.Fatal(err)
		os.Exit(1)
	}
	controller.Start()
	controller.Shutdown()
}
