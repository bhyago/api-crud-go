package student

import "github.com/bhyago/api-crud-go/entities"

func (su *StudentUseCase) Create(fullName string, age int) (entities.Studant, error) {
	student := entities.Newstudent(fullName, age)

	err := su.Database.StudentRepository.Create(&student)
	return student, err
}
