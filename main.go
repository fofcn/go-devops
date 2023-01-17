package main

import (
	"go-micro.dev/v4"
)

func main() {
	startdemoservice()
	// zlog.InitLogger()
	// defer zlog.Logger.Sync()

	// osargs := args.NewAppArgs()
	// controller := NewController(&osargs)
	// err := controller.Init()
	// if err != nil {
	// 	zlog.Logger.Fatal(err)
	// 	os.Exit(1)
	// }
	// controller.Start()
	// controller.Shutdown()
}

func startdemoservice() {
	service := micro.NewService(
		micro.Name("helloworld"),
	)
	service.Init()
	service.Run()
}
