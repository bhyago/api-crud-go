package students

import (
	"net/http"

	"github.com/bhyago/api-crud-go/api/controller"
	"github.com/bhyago/api-crud-go/entities"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	student := entities.Newstudent(
		input.FullName,
		input.Age,
	)
	entities.Students = append(entities.Students, student)

	c.JSON(http.StatusAccepted, student)
}
