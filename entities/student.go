package entities

import (
	"github.com/bhyago/api-crud-go/entities/shared"
	"github.com/google/uuid"
)

type Studant struct {
	ID       uuid.UUID `json:"id" bson:"_id"`
	FullName string    `json:"full_name" bson:"full_name"`
	Age      int       `json:"age" bson:"age"`
}

func Newstudent(fullName string, age int) Studant {
	return Studant{
		ID:       shared.GetID(),
		FullName: fullName,
		Age:      age,
	}
}

type StudentRepository interface {
	Create(studant *Studant) error
	FindAll() ([]Studant, error)
	FindByID(uuid.UUID) (Studant, error)
	Update(studant *Studant) error
	Delete(uuid.UUID) error
}

var Students = []Studant{
	{ID: shared.GetID(), FullName: "João Eleno", Age: 20},
	{ID: shared.GetID(), FullName: "Maria Elena", Age: 21},
}
