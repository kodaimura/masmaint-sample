package controller

import (
	"github.com/gin-gonic/gin"

	"masmaint/service"
	"masmaint/dto"
	"masmaint/model/entity"
)

type DepartmentService interface {
	Create(dDto *dto.DepartmentDto) error
	Update(dDto *dto.DepartmentDto) error
	Delete(dDto *dto.DepartmentDto) error
	GetAll() ([]entity.Department, error)
	GetOne(id int) (entity.Department, error)
}

type DepartmentController struct {
	dServ *service.DepartmentService
}


func NewDepartmentController() *DepartmentController {
	dServ := service.NewDepartmentService()
	return &DepartmentController{dServ}
}


//GET /employee
func (ctr *DepartmentController) GetDepartmentPage(c *gin.Context) {
	c.HTML(200, "department.html", gin.H{})
}

//GET /api/employee
func (ctr *DepartmentController) GetDepartment(c *gin.Context) {
	ds, err := ctr.dServ.GetAll()

	if err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}

	c.JSON(200, ds)
}

//POST /api/employee
func (ctr *DepartmentController) PostDepartment(c *gin.Context) {
	var dDto dto.DepartmentDto 

	if err := c.ShouldBindJSON(&dDto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if err := ctr.dServ.Create(&dDto); err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}

//PUT /api/employee
func (ctr *DepartmentController) PutDepartment(c *gin.Context) {
	var dDto dto.DepartmentDto 

	if err := c.ShouldBindJSON(&dDto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if err := ctr.dServ.Update(&dDto); err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}

	d, err := ctr.dServ.GetOne(dDto.Id)

	if err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}
	
	c.JSON(200, d)
}

//DELETE /api/employee
func (ctr *DepartmentController) DeleteDepartment(c *gin.Context) {
	var dDto dto.DepartmentDto 

	if err := c.ShouldBindJSON(&dDto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if err := ctr.dServ.Delete(&dDto); err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}
	
	c.JSON(200, gin.H{})
}