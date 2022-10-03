package server

import (
	"assignment-2/app/interface/container"
	"assignment-2/app/shared/config"
	"assignment-2/app/transport"

	"github.com/gin-gonic/gin"
)

func SetupServer(container container.Container) {
	app := gin.Default()
	transport := transport.SetupTransport(container)
	SetupRouter(transport, app)

	app.Run(config.Server.PORTHTTP)
}