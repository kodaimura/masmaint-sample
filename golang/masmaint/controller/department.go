package controller

import (
	"github.com/gin-gonic/gin"

	"masmaint/service"
	"masmaint/dto"
)

type DepartmentService interface {
	Create(dDto *dto.DepartmentDto) (dto.DepartmentDto, error)
	Update(dDto *dto.DepartmentDto) (dto.DepartmentDto, error)
	Delete(dDto *dto.DepartmentDto) error
	GetAll() ([]dto.DepartmentDto, error)
	GetOne(id int64) ([]dto.DepartmentDto, error)
}

type DepartmentController struct {
	dServ *service.DepartmentService
}


func NewDepartmentController() *DepartmentController {
	dServ := service.NewDepartmentService()
	return &DepartmentController{dServ}
}


//GET /department
func (ctr *DepartmentController) GetDepartmentPage(c *gin.Context) {
	c.HTML(200, "department.html", gin.H{})
}

//GET /api/department
func (ctr *DepartmentController) GetDepartment(c *gin.Context) {
	ds, err := ctr.dServ.GetAll()

	if err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}

	c.JSON(200, ds)
}

//POST /api/department
func (ctr *DepartmentController) PostDepartment(c *gin.Context) {
	var dDto dto.DepartmentDto 

	if err := c.ShouldBindJSON(&dDto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	ret, err := ctr.dServ.Create(&dDto)

	if err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}

	c.JSON(200, ret)
}

//PUT /api/department
func (ctr *DepartmentController) PutDepartment(c *gin.Context) {
	var dDto dto.DepartmentDto 

	if err := c.ShouldBindJSON(&dDto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	ret, err := ctr.dServ.Update(&dDto)

	if  err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}
	
	c.JSON(200, ret)
}

//DELETE /api/department
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