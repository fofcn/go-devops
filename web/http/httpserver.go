package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"taskmanager/web/http/handler"
)

type HttpServerConfig struct {
	IP   string
	Port int
}

type HttpServer interface {
	Config(HttpServerConfig)
	Init()
	Start()
	Shutdown()
	RegisterHandler(handler.HttpHandler)
}

type httpserver struct {
	mux    *http.ServeMux
	config HttpServerConfig
}

func NewHttpServer() HttpServer {
	return &httpserver{}
}

func (server *httpserver) Config(config HttpServerConfig) {
	server.config.IP = config.IP
	server.config.Port = config.Port
}

func (server *httpserver) Init() {
	server.mux = http.NewServeMux()
}

func (server *httpserver) Start() {
	addr := server.buildAddr()

	srv := &http.Server{
		Addr:    addr,
		Handler: server.mux,
	}

	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-done

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Fatal("Shutdown server:", err)
		}
	}()

	err := srv.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}
}

func (server *httpserver) RegisterHandler(handler handler.HttpHandler) {
	server.mux.Handle(handler.Pattern(), handler.Handler())
}

func (server *httpserver) Shutdown() {
}

func (server httpserver) buildAddr() string {
	return server.config.IP + ":" + fmt.Sprintf("%v", server.config.Port)
}
