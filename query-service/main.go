package main

import (
	"github.com/AchmadAlli/go-cqrs/query-service/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	routers.RegisterApiRoutes(g)

	g.Run()
}
