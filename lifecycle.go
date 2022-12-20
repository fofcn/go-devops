package main

type LifeCycle interface {
	Init(obj interface{}) error
	Start() error
	Shutdown() error
}
