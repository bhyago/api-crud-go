package student

import (
	"github.com/bhyago/api-crud-go/entities"
	"github.com/google/uuid"
)

func (su *StudentUseCase) SearchById(id uuid.UUID) (student entities.Studant, err error) {
	student, err = su.Database.StudentRepository.FindByID(id)
	return student, err
}
