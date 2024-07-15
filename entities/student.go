package entities

import (
	"github.com/bhyago/api-crud-go/entities/shared"
	"github.com/google/uuid"
)

type Studant struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Age      int       `json:"age"`
}

func Newstudent(fullName string, age int) Studant {
	return Studant{
		ID:       shared.GetID(),
		FullName: fullName,
		Age:      age,
	}
}

var Students = []Studant{
	{ID: shared.GetID(), FullName: "João Eleno", Age: 20},
	{ID: shared.GetID(), FullName: "Maria Elena", Age: 21},
}
