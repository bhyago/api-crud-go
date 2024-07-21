package students

import (
	"net/http"

	"github.com/bhyago/api-crud-go/api/controller"
	student_usecase "github.com/bhyago/api-crud-go/usecases/student"
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	StudentUseCase *student_usecase.StudentUseCase
}

func NewStudentController(su *student_usecase.StudentUseCase) *StudentController {
	return &StudentController{
		StudentUseCase: su,
	}
}

func (sc *StudentController) Create(c *gin.Context) {
	input, err := getInputBody(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	student, err := sc.StudentUseCase.Create(input.FullName, input.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, student)
}

func (sc *StudentController) List(c *gin.Context) {
	students, err := sc.StudentUseCase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	output, err := getOutputListStudents(students)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, output)
}

func (sc *StudentController) Details(c *gin.Context) {
	id, err := getInputId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Problem parsing ID"))
		return
	}

	studentFound, err := sc.StudentUseCase.SearchById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, controller.NewResponseMessageError(err.Error()))
		return
	}

	output, err := getOutputStudent(studentFound)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, output)
}

func (sc *StudentController) Update(c *gin.Context) {
	id, err := getInputId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Problem parsing ID"))
		return
	}

	input, err := getInputBody(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	student, err := sc.StudentUseCase.Update(id, input.FullName, input.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	output, err := getOutputStudent(student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, output)
}

func (sc *StudentController) Delete(c *gin.Context) {
	id, err := getInputId(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	if err = sc.StudentUseCase.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, controller.NewResponseMessage("Student deleted"))
}
