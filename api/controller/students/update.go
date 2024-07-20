package students

import (
	"net/http"

	"github.com/bhyago/api-crud-go/api/controller"
	"github.com/bhyago/api-crud-go/entities"
	"github.com/bhyago/api-crud-go/entities/shared"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var input Input
	var studantFound entities.Studant
	var newStudents []entities.Studant

	err := c.ShouldBindJSON(&input)
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

	for _, studentElement := range entities.Students {
		if studentElement.ID == input.UUID {
			studantFound = studentElement
		}
	}

	if studantFound.ID == shared.GetUuidEmpty() {
		c.JSON(http.StatusNotFound, controller.NewResponseMessageError("Student not found"))
		return
	}

	studantFound.FullName = input.FullName
	studantFound.Age = input.Age

	for _, studentElement := range entities.Students {
		if studantFound.ID == studentElement.ID {
			newStudents = append(newStudents, studantFound)
		} else {
			newStudents = append(newStudents, studentElement)
		}

	}

	entities.Students = newStudents

	c.JSON(http.StatusOK, studantFound)

}
