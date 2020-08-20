package main

import (
	"github.com/lguilins/oraculosweb-api/cmd/server"
	"github.com/lguilins/oraculosweb-api/pkg/registry"
)

func main() {
	app := registry.NewContainer()
	server.RunHTTPServer(app)
}
