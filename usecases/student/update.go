package student

import (
	"errors"

	"github.com/bhyago/api-crud-go/entities"
	"github.com/bhyago/api-crud-go/entities/shared"
	"github.com/google/uuid"
)

func Update(id uuid.UUID, fullName string, age int) (student entities.Studant, err error) {
	var newStudents []entities.Studant

	for _, studentElement := range entities.Students {
		if studentElement.ID == id {
			student = studentElement
		}
	}

	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("Student not found")
	}

	student.FullName = fullName
	student.Age = age

	for _, studentElement := range entities.Students {
		if student.ID == studentElement.ID {
			newStudents = append(newStudents, student)
		} else {
			newStudents = append(newStudents, studentElement)
		}
	}

	entities.Students = newStudents

	return student, err
}
