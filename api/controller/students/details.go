package students

import (
	"net/http"

	"github.com/bhyago/api-crud-go/api/controller"
	"github.com/bhyago/api-crud-go/entities"
	"github.com/bhyago/api-crud-go/entities/shared"
	"github.com/gin-gonic/gin"
)

func Details(c *gin.Context) {
	var input Input
	var student entities.Studant
	var err error

	input.ID = c.Params.ByName("id")
	input.UUID, err = shared.ParseIDString(input.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Problem parsing ID"))
		return
	}

	for _, studentelement := range entities.Students {
		if studentelement.ID == input.UUID {
			student = studentelement
		}
	}

	c.JSON(http.StatusOK, student)

}
