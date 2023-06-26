package controller

import (
	"github.com/gin-gonic/gin"

	"masmaint/core/auth"
)


func SetRouter(r *gin.Engine) {
	employeeController := NewEmployeeController()
	departmentController := NewDepartmentController()

	rm := r.Group("/mastertables", auth.NoopAuthMiddleware())
	{
		rm.GET("/", func(c *gin.Context) {
			c.HTML(200, "index.html", gin.H{})
		})

		rm.GET("/employee", employeeController.GetEmployeePage)
		rm.GET("/api/employee", employeeController.GetEmployee)
		rm.POST("/api/employee", employeeController.PostEmployee)
		rm.PUT("/api/employee", employeeController.PutEmployee)
		rm.DELETE("/api/employee", employeeController.DeleteEmployee)

		rm.GET("/department", departmentController.GetDepartmentPage)
		rm.GET("/api/department", departmentController.GetDepartment)
		rm.POST("/api/department", departmentController.PostDepartment)
		rm.PUT("/api/department", departmentController.PutDepartment)
		rm.DELETE("/api/department", departmentController.DeleteDepartment)
	}
}