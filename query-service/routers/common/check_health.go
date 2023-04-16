package common

import (
	"net/http"

	"github.com/AchmadAlli/go-cqrs/query-service/utils"
	"github.com/gin-gonic/gin"
)

func RegisterApiCheckHealthRoute(g *gin.RouterGroup) {
	g.GET("/check-health", checkHealth)
}

func checkHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, utils.NewResponse("ok", nil))
}
