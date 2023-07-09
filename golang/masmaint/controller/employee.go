package controller

import (
	"github.com/gin-gonic/gin"

	"masmaint/service"
	"masmaint/dto"
)


type EmployeeService interface {
	GetAll() ([]dto.EmployeeDto, error)
	Create(eDto *dto.EmployeeDto) (dto.EmployeeDto, error)
	Update(eDto *dto.EmployeeDto) (dto.EmployeeDto, error)
	Delete(eDto *dto.EmployeeDto) error
}

type employeeController struct {
	eServ EmployeeService
}

func NewEmployeeController() *employeeController {
	eServ := service.NewEmployeeService()
	return &employeeController{eServ}
}


//GET /employee
func (ctr *employeeController) GetEmployeePage(c *gin.Context) {
	c.HTML(200, "employee.html", gin.H{})
}


//GET /api/employee
func (ctr *employeeController) GetEmployee(c *gin.Context) {
	ret, err := ctr.eServ.GetAll()

	if err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}

	c.JSON(200, ret)
}


//POST /api/employee
func (ctr *employeeController) PostEmployee(c *gin.Context) {
	var eDto dto.EmployeeDto

	if err := c.ShouldBindJSON(&eDto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	ret, err := ctr.eServ.Create(&eDto)

	if err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}

	c.JSON(200, ret)
}


//PUT /api/employee
func (ctr *employeeController) PutEmployee(c *gin.Context) {
	var eDto dto.EmployeeDto

	if err := c.ShouldBindJSON(&eDto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	ret, err := ctr.eServ.Update(&eDto)

	if err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}

	c.JSON(200, ret)
}


//DELETE /api/employee
func (ctr *employeeController) DeleteEmployee(c *gin.Context) {
	var eDto dto.EmployeeDto

	if err := c.ShouldBindJSON(&eDto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if err := ctr.eServ.Delete(&eDto); err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}
