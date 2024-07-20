package student

import "github.com/bhyago/api-crud-go/entities"

func Create(fullName string, age int) (student entities.Studant, err error) {
	student = entities.Newstudent(fullName, age)

	entities.Students = append(entities.Students, student)
	return student, err
}
