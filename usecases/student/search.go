package student

import (
	"errors"

	"github.com/bhyago/api-crud-go/entities"
	"github.com/bhyago/api-crud-go/entities/shared"
	"github.com/google/uuid"
)

func SearchById(id uuid.UUID) (student entities.Studant, err error) {

	for _, studentelement := range entities.Students {
		if studentelement.ID == id {
			student = studentelement
		}
	}

	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("Student not found")
	}
	return student, err
}
