package main

import (
	"taskmanager/web/http"
	"taskmanager/web/http/handler"

	"go-micro.dev/v4"
)

func main() {
	starthttpserver()
	// starthttpserver()
	// startdemoservice()
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

func starthttpserver() {
	httpconfig := http.HttpServerConfig{
		IP:   "localhost",
		Port: 8080,
	}

	httpserver := http.NewHttpServer()
	httpserver.Config(httpconfig)
	httpserver.Init()
	httpserver.RegisterHandler(handler.NewIndexHandler())
	handler, err := handler.NewLoginHandler()
	if err != nil {
		return
	}
	httpserver.RegisterHandler(handler)
	httpserver.Start()
}
