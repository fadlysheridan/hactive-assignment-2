package main

import (
	"assignment-2/app/interface/container"
	"assignment-2/app/interface/server"
)


func main() {
	container := container.SetupContainer()
	server.SetupServer(container)
}