package routers

import (
	"github.com/AchmadAlli/go-cqrs/query-service/routers/common"
	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(e *gin.Engine) {
	g := e.Group("/api")

	common.RegisterApiCheckHealthRoute(g)
}
