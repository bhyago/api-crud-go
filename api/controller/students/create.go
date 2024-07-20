package students

import (
	"net/http"

	"github.com/bhyago/api-crud-go/api/controller"
	"github.com/gin-gonic/gin"
)

func (sc *StudentController) Create(c *gin.Context) {
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	student, err := sc.StudentUseCase.Create(input.FullName, input.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusAccepted, student)
}
