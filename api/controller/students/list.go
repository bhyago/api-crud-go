package students

import (
	"net/http"

	"github.com/bhyago/api-crud-go/api/controller"
	student_usecases "github.com/bhyago/api-crud-go/usecases/student"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	students, err := student_usecases.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, students)
}
