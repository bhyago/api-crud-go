package students

import (
	"net/http"

	"github.com/bhyago/api-crud-go/api/controller"
	"github.com/gin-gonic/gin"
)

func (sc *StudentController) List(c *gin.Context) {
	students, err := sc.StudentUseCase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, students)
}
