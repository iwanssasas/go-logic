package server

import (
	"go-logic/003.dependencyInjection/router"

	"github.com/gin-gonic/gin"
)

func NewServer() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	router.NewRouter(r).Product()
	r.Run()
}
