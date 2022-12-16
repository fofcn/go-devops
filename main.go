package main

func main() {
	var context ApplicationContext
	context.Init("config/pipeline/setup-demo.yaml")
	context.Start()
	context.Shutdown()
}
