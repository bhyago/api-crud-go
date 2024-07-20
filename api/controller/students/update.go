package students

import (
	"net/http"

	"github.com/bhyago/api-crud-go/api/controller"
	"github.com/bhyago/api-crud-go/entities/shared"
	student_usecases "github.com/bhyago/api-crud-go/usecases/student"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var input Input
	var err error

	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError((err.Error())))
		return
	}

	input.ID = c.Params.ByName("id")
	input.UUID, err = shared.ParseIDString(input.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	student, err := student_usecases.Update(input.UUID, input.FullName, input.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, student)

}
