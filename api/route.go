package api

import (
	controller_infra "github.com/bhyago/api-crud-go/api/controller/infra"
	controller_student "github.com/bhyago/api-crud-go/api/controller/students"
)

func (s *Service) GetRoutes() {
	s.Engine.GET("/heart", controller_infra.Heart)
	students := s.Engine.Group("/students")
	{
		students.GET("/", controller_student.List)
		students.POST("/", controller_student.Create)
		// students.GET("/:id", controller_infra.GetStudent)
		// students.PUT("/:id", controller_infra.PutStudent)
		// students.DELETE("/:id", controller_infra.DeleteStudent)
	}

}
