package api

import (
	"fmt"

	"github.com/bhyago/api-crud-go/api/controller/students"
	"github.com/bhyago/api-crud-go/infra/config"
	"github.com/bhyago/api-crud-go/infra/database"
	student_usecase "github.com/bhyago/api-crud-go/usecases/student"
	"github.com/gin-gonic/gin"
)

type Service struct {
	Engine            *gin.Engine
	Database          *database.Database
	StudentController *students.StudentController
}

func NewService(db *database.Database) *Service {
	return &Service{
		Engine:   gin.Default(),
		Database: db,
	}
}

func (s *Service) GetControllers() {
	studentUseCase := student_usecase.NewStudentUseCase(s.Database)
	s.StudentController = students.NewStudentController(studentUseCase)
}

func (s *Service) Start() error {
	s.GetControllers()
	s.GetRoutes()
	return s.Engine.Run(fmt.Sprintf(":%d", config.Env.ApiPort))
}
