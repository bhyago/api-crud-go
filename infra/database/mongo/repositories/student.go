package repositories

import (
	"context"

	"github.com/bhyago/api-crud-go/entities"
	"github.com/bhyago/api-crud-go/infra/database/mongo"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	mongo_drive "go.mongodb.org/mongo-driver/mongo"
)

const (
	StudentCollection = "students"
)

type StudentRepository struct {
	Ctx        context.Context
	Collection *mongo_drive.Collection
}

func NewStudentRepository(ctx context.Context, client *mongo_drive.Client) *StudentRepository {
	collection, err := mongo.GetCollection(ctx, client, StudentCollection)
	if err != nil {
		return nil
	}

	return &StudentRepository{
		Ctx:        ctx,
		Collection: collection,
	}
}

func (sr StudentRepository) Create(student *entities.Studant) error {
	_, err := sr.Collection.InsertOne(sr.Ctx, student)
	return err
}

func (sr StudentRepository) FindAll() (students []entities.Studant, err error) {
	cursor, err := sr.Collection.Find(sr.Ctx, bson.M{})
	if err != nil {
		return students, err
	}

	defer cursor.Close(sr.Ctx)

	for cursor.Next(sr.Ctx) {
		var student entities.Studant
		if err = cursor.Decode(&student); err != nil {
			return students, err
		}

		students = append(students, student)
	}

	return students, nil
}

func (sr StudentRepository) FindByID(id uuid.UUID) (student entities.Studant, err error) {
	err = sr.Collection.FindOne(sr.Ctx, bson.M{"_id": id}).Decode(&student)
	return student, err
}

func (sr StudentRepository) Update(student *entities.Studant) error {
	_, err := sr.Collection.UpdateOne(
		sr.Ctx,
		bson.M{"_id": student.ID},
		bson.M{"$set": student})

	return err
}

func (sr StudentRepository) Delete(id uuid.UUID) error {
	_, err := sr.Collection.DeleteOne(sr.Ctx, bson.M{"_id": id})
	return err
}
