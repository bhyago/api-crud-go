package students

import (
	"net/http"

	"github.com/bhyago/api-crud-go/api/controller"
	"github.com/bhyago/api-crud-go/entities"
	"github.com/bhyago/api-crud-go/entities/shared"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var input Input
	var newStudents []entities.Studant
	var err error

	input.ID = c.Params.ByName("id")
	input.UUID, err = shared.ParseIDString(input.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Problem parsing ID"))
		return
	}

	for _, studentelement := range entities.Students {
		if studentelement.ID != input.UUID {
			newStudents = append(newStudents, studentelement)
		}
	}

	entities.Students = newStudents

	c.JSON(http.StatusOK, controller.NewResponseMessage("Student deleted"))

}
