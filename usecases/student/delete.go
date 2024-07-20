package student

import (
	"errors"

	"github.com/bhyago/api-crud-go/entities/shared"
	"github.com/google/uuid"
)

func (su *StudentUseCase) Delete(id uuid.UUID) (err error) {
	student, err := su.Database.StudentRepository.FindByID(id)
	if err != nil {
		return err
	}

	if student.ID == shared.GetUuidEmpty() {
		return errors.New("Student not found")
	}

	err = su.Database.StudentRepository.Delete(id)

	return err
}
