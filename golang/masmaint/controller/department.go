package controller

import (
	"github.com/gin-gonic/gin"

	cerror "masmaint/core/error"
	"masmaint/service"
	"masmaint/dto"
)


type DepartmentService interface {
	GetAll() ([]dto.DepartmentDto, error)
	Create(dDto *dto.DepartmentDto) (dto.DepartmentDto, error)
	Update(dDto *dto.DepartmentDto) (dto.DepartmentDto, error)
	Delete(dDto *dto.DepartmentDto) error
}

type departmentController struct {
	dServ DepartmentService
}

func NewDepartmentController() *departmentController {
	dServ := service.NewDepartmentService()
	return &departmentController{dServ}
}


//GET /department
func (ctr *departmentController) GetDepartmentPage(c *gin.Context) {
	c.HTML(200, "department.html", gin.H{})
}


//GET /api/department
func (ctr *departmentController) GetDepartment(c *gin.Context) {
	ret, err := ctr.dServ.GetAll()

	if err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}

	c.JSON(200, ret)
}


//POST /api/department
func (ctr *departmentController) PostDepartment(c *gin.Context) {
	var dDto dto.DepartmentDto

	if err := c.ShouldBindJSON(&dDto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	ret, err := ctr.dServ.Create(&dDto)

	if err != nil {
		if _, ok := err.(*cerror.InvalidArgumentError); ok {
			c.JSON(400, gin.H{})
		} else {
			c.JSON(500, gin.H{})
		}
		c.Abort()
		return
	}

	c.JSON(200, ret)
}


//PUT /api/department
func (ctr *departmentController) PutDepartment(c *gin.Context) {
	var dDto dto.DepartmentDto

	if err := c.ShouldBindJSON(&dDto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	ret, err := ctr.dServ.Update(&dDto)

	if err != nil {
		if _, ok := err.(*cerror.InvalidArgumentError); ok {
			c.JSON(400, gin.H{})
		} else {
			c.JSON(500, gin.H{})
		}
		c.Abort()
		return
	}

	c.JSON(200, ret)
}


//DELETE /api/department
func (ctr *departmentController) DeleteDepartment(c *gin.Context) {
	var dDto dto.DepartmentDto

	if err := c.ShouldBindJSON(&dDto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if err := ctr.dServ.Delete(&dDto); err != nil {
		if _, ok := err.(*cerror.InvalidArgumentError); ok {
			c.JSON(400, gin.H{})
		} else {
			c.JSON(500, gin.H{})
		}
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}
