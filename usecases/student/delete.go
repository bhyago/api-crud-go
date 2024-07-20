package student

import (
	"github.com/bhyago/api-crud-go/entities"
	"github.com/google/uuid"
)

func Delete(id uuid.UUID) (err error) {
	var newStudents []entities.Studant

	for _, studentelement := range entities.Students {
		if studentelement.ID != id {
			newStudents = append(newStudents, studentelement)
		}
	}

	entities.Students = newStudents
	return err
}
