package student

import "github.com/bhyago/api-crud-go/entities"

func (su *StudentUseCase) List() (students []entities.Studant, err error) {
	return su.Database.StudentRepository.FindAll()
}
