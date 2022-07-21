package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jtao539/autocode/template/api"
)

type DepartmentRouter struct {
	webApi api.DepartmentApi
}

func (d *DepartmentRouter) InitDepartmentRouter(g *gin.Engine) {
	deRouter := g.Group("department")
	{
		deRouter.POST("list", d.webApi.GetDepartmentList)
		deRouter.GET(":id", d.webApi.GetDepartmentById)
		deRouter.POST("add", d.webApi.AddDepartment)
		deRouter.POST("delete", d.webApi.DeleteDepartment)
		deRouter.POST("update", d.webApi.UpdateDepartment)
	}
}
