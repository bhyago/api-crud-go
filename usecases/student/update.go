package student

import (
	"errors"

	"github.com/bhyago/api-crud-go/entities"
	"github.com/bhyago/api-crud-go/entities/shared"
	"github.com/google/uuid"
)

func (su *StudentUseCase) Update(id uuid.UUID, fullName string, age int) (entities.Studant, error) {
	student, err := su.Database.StudentRepository.FindByID(id)
	if err != nil {
		return student, err
	}

	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("Student not found")
	}

	student.FullName = fullName
	student.Age = age

	err = su.Database.StudentRepository.Update(&student)

	return student, err
}
