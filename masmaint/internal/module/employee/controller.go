package employee

import (
	"github.com/gin-gonic/gin"
)

type controller struct {
	service Service
}

func NewController() *controller {
	service := NewService()
	return &controller{service}
}


//GET /employee
func (ctr *controller) GetPage(c *gin.Context) {
	c.HTML(200, "employee.html", gin.H{})
}


//GET /api/employee
func (ctr *controller) Get(c *gin.Context) {
	ret, err := ctr.service.Get()

	if err != nil {
		c.JSON(500, gin.H{"error":"予期せぬエラーが発生しました。"})
		c.Abort()
		return
	}

	c.JSON(200, ret)
}


//POST /api/employee
func (ctr *controller) Post(c *gin.Context) {
	var req PostBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error":"不正なリクエストです。"})
		c.Abort()
		return
	}

	ret, err := ctr.service.Create(req)

	if err != nil {
		c.JSON(500, gin.H{"error":"予期せぬエラーが発生しました。"})
		c.Abort()
		return
	}

	c.JSON(200, ret)
}


//PUT /api/employee
func (ctr *controller) Put(c *gin.Context) {
	var req PutBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error":"不正なリクエストです。"})
		c.Abort()
		return
	}

	ret, err := ctr.service.Update(req)

	if err != nil {
		c.JSON(500, gin.H{"error":"予期せぬエラーが発生しました。"})
		c.Abort()
		return
	}

	c.JSON(200, ret)
}


//DELETE /api/employee
func (ctr *controller) Delete(c *gin.Context) {
	var req DeleteBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error":"不正なリクエストです。"})
		c.Abort()
		return
	}

	if err := ctr.service.Delete(req); err != nil {
		c.JSON(500, gin.H{"error":"予期せぬエラーが発生しました。"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}