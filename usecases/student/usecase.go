package student

import "github.com/bhyago/api-crud-go/infra/database"

type StudentUseCase struct {
	Database *database.Database
}

func NewStudentUseCase(db *database.Database) *StudentUseCase {
	return &StudentUseCase{
		Database: db,
	}
}
