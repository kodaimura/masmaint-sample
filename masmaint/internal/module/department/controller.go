package department

import (
	"github.com/gin-gonic/gin"
	"masmaint/internal/core/errs"
)

type controller struct {
	service Service
}

func NewController() *controller {
	service := NewService()
	return &controller{service}
}


//GET /department
func (ctr *controller) GetPage(c *gin.Context) {
	c.HTML(200, "department.html", gin.H{})
}


//GET /api/department
func (ctr *controller) Get(c *gin.Context) {
	ret, err := ctr.service.Get()
	if err != nil {
		c.Error(errs.NewError(err))
		return
	}

	c.JSON(200, ret)
}


//POST /api/department
func (ctr *controller) Post(c *gin.Context) {
	var req PostBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errs.NewError(err))
		return
	}

	ret, err := ctr.service.Create(req)
	if err != nil {
		c.Error(errs.NewError(err))
		return
	}

	c.JSON(200, ret)
}


//PUT /api/department
func (ctr *controller) Put(c *gin.Context) {
	var req PutBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errs.NewError(err))
		return
	}

	ret, err := ctr.service.Update(req)
	if err != nil {
		c.Error(errs.NewError(err))
		return
	}

	c.JSON(200, ret)
}


//DELETE /api/department
func (ctr *controller) Delete(c *gin.Context) {
	var req DeleteBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errs.NewError(err))
		return
	}

	if err := ctr.service.Delete(req); err != nil {
		c.Error(errs.NewError(err))
		return
	}

	c.JSON(200, gin.H{})
}