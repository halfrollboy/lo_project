package handler

import (
	"github.com/gin-gonic/gin"
)

type router struct {
	Router *gin.Engine
}

func NewRouter() *router {
	return &router{gin.Default()}
}

func (r *router) baseConfigureRouter() {
	r.Router.LoadHTMLGlob("internal/templates/*")
	r.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Router.POST("/order", ValidateOrder)
	r.Router.GET("/order", GetOrderTemplate)
}
