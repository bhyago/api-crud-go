package students

import (
	"net/http"

	"github.com/bhyago/api-crud-go/api/controller"
	student_usecases "github.com/bhyago/api-crud-go/usecases/student"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	student, err := student_usecases.Create(input.FullName, input.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusAccepted, student)
}
