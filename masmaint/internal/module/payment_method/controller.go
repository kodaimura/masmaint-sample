package payment_method

import (
	"github.com/gin-gonic/gin"
	"masmaint/internal/module"
)

type controller struct {
	service Service
}

func NewController() *controller {
	service := NewService()
	return &controller{service}
}


//GET /payment_method
func (ctr *controller) GetPage(c *gin.Context) {
	c.HTML(200, "payment_method.html", gin.H{})
}


//GET /api/payment_method
func (ctr *controller) Get(c *gin.Context) {
	ret, err := ctr.service.Get()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ret)
}


//POST /api/payment_method
func (ctr *controller) Post(c *gin.Context) {
	var req PostBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(module.NewBindError(err, &req))
		return
	}

	ret, err := ctr.service.Create(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ret)
}


//PUT /api/payment_method
func (ctr *controller) Put(c *gin.Context) {
	var req PutBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(module.NewBindError(err, &req))
		return
	}

	ret, err := ctr.service.Update(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ret)
}


//DELETE /api/payment_method
func (ctr *controller) Delete(c *gin.Context) {
	var req DeleteBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(module.NewBindError(err, &req))
		return
	}

	if err := ctr.service.Delete(req); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{})
}