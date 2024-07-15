package infra

import (
	"net/http"

	"github.com/bhyago/api-crud-go/api/controller"
	"github.com/gin-gonic/gin"
)

func Heart(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, controller.NewResponseMessage("I'm alive!"))
}
