package api

import (
	controller_infra "github.com/bhyago/api-crud-go/api/controller/infra"
)

func (s *Service) GetRoutes() {
	s.Engine.GET("/heart", controller_infra.Heart)
	students := s.Engine.Group("/students")
	{
		students.GET("/", s.StudentController.List)
		students.POST("/", s.StudentController.Create)
		students.PUT("/:id", s.StudentController.Update)
		students.DELETE("/:id", s.StudentController.Delete)
		students.GET("/:id", s.StudentController.Details)
	}
}
