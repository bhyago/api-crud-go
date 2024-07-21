package students

import (
	"github.com/bhyago/api-crud-go/entities"
	"github.com/bhyago/api-crud-go/entities/shared"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gookit/validate"
)

func getInputBody(c *gin.Context) (input InputBody, err error) {
	err = c.Bind(&input)
	if err != nil {
		return input, err
	}

	validation := validate.Struct(input)
	if !validation.Validate() {
		return input, validation.Errors
	}

	return input, err
}

func getOutputListStudents(students []entities.Studant) (output OutputListStudents, err error) {
	for _, s := range students {
		output.Students = append(output.Students, OutputStudent{
			ID:       s.ID,
			FullName: s.FullName,
			Age:      s.Age,
		})
	}

	return output, err
}

func getInputId(c *gin.Context) (id uuid.UUID, err error) {
	inputID := c.Params.ByName("id")

	id, err = shared.ParseIDString(inputID)

	if err != nil {
		return id, err
	}

	return id, err
}

func getOutputStudent(student entities.Studant) (output OutputStudent, err error) {
	return OutputStudent{
		ID:       student.ID,
		FullName: student.FullName,
		Age:      student.Age,
	}, err

}
