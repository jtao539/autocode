package router

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	DepartmentRouter
}

var router Router

func Register(g *gin.Engine) {
	router.InitDepartmentRouter(g)
}
