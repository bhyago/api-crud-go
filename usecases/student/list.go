package student

import "github.com/bhyago/api-crud-go/entities"

func List() (students []entities.Studant, err error) {
	return entities.Students, err
}
