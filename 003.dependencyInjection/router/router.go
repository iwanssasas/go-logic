package router

import (
	"go-logic/003.dependencyInjection/handler"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Product()
}

type router struct {
	router  *gin.Engine
	handler handler.Handler
}

func NewRouter(r *gin.Engine) Router {
	return &router{
		router:  r,
		handler: handler.NewHandler(),
	}
}

func (r *router) Product() {
	r.router.GET("/ping", r.handler.TestPing)
	r.router.POST("/students", r.handler.AddStudents)
	r.router.GET("/students", r.handler.GetAllStudents)
	r.router.DELETE("/students/:id", r.handler.DeleteStudents)
	r.router.PUT("/students/:id", r.handler.UpdateStudents)
}
