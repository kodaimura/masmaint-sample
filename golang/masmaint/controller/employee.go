package controller

import (
	"github.com/gin-gonic/gin"

	"masmaint/service"
	"masmaint/dto"
	"masmaint/model/entity"
)

type EmployeeService interface {
	Create(eDto *dto.EmployeeDto) error
	Update(eDto *dto.EmployeeDto) error
	Delete(eDto *dto.EmployeeDto) error
	GetAll() ([]entity.Employee, error)
	GetOne(id int) (entity.Employee, error)
}

type EmployeeController struct {
	eServ *service.EmployeeService
}


func NewEmployeeController() *EmployeeController {
	eServ := service.NewEmployeeService()
	return &EmployeeController{eServ}
}


//GET /employee
func (ctr *EmployeeController) GetEmployeePage(c *gin.Context) {
	c.HTML(200, "employee.html", gin.H{})
}

//GET /api/employee
func (ctr *EmployeeController) GetEmployee(c *gin.Context) {
	es, err := ctr.eServ.GetAll()

	if err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}

	c.JSON(200, es)
}

//POST /api/employee
func (ctr *EmployeeController) PostEmployee(c *gin.Context) {
	var eDto dto.EmployeeDto 

	if err := c.ShouldBindJSON(&eDto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if err := ctr.eServ.Create(&eDto); err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}

//PUT /api/employee
func (ctr *EmployeeController) PutEmployee(c *gin.Context) {
	var eDto dto.EmployeeDto 

	if err := c.ShouldBindJSON(&eDto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if err := ctr.eServ.Update(&eDto); err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}

	e, err := ctr.eServ.GetOne(eDto.Id)

	if err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}
	
	c.JSON(200, e)
}

//DELETE /api/employee
func (ctr *EmployeeController) DeleteEmployee(c *gin.Context) {
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