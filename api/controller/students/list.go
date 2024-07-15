package students

import (
	"net/http"

	"github.com/bhyago/api-crud-go/entities"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	c.JSON(http.StatusOK, entities.Students)
}
