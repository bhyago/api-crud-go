package students

import (
	"net/http"

	"github.com/bhyago/api-crud-go/api/controller"
	"github.com/bhyago/api-crud-go/entities/shared"
	"github.com/gin-gonic/gin"
)

func (sc *StudentController) Delete(c *gin.Context) {
	var input Input
	var err error

	input.ID = c.Params.ByName("id")
	input.UUID, err = shared.ParseIDString(input.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Problem parsing ID"))
		return
	}

	if err = sc.StudentUseCase.Delete(input.UUID); err != nil {
		c.JSON(http.StatusNotFound, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, controller.NewResponseMessage("Student deleted"))
}
