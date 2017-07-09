package app

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"thrifty/gen-go/personservice"
	"thrifty/controllers"
)

type App struct{}

func (app *App) Startup() {
	fmt.Println("Application starting...")

	addr := "localhost:9090"

	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTSimpleJSONProtocolFactory()

	controller := &controllers.EmployeeController{}
	processor := personservice.NewPersonServiceProcessor(controller)
	transport, _ := thrift.NewTServerSocket(addr)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	server.Serve()
}

func (app *App) Shutdown() {
	fmt.Println("Application shutting down...")
}
