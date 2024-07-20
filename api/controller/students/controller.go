package students

import (
	student_usecase "github.com/bhyago/api-crud-go/usecases/student"
)

type StudentController struct {
	StudentUseCase *student_usecase.StudentUseCase
}

func NewStudentController(su *student_usecase.StudentUseCase) *StudentController {
	return &StudentController{
		StudentUseCase: su,
	}
}
