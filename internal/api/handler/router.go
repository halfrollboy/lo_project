package handler

import (
	"lo_project/internal/config"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type router struct {
	Router *gin.Engine
	Config *config.AppConfig
}

func NewRouter(config *config.AppConfig) *router {
	r := &router{
		Router: gin.Default(),
		Config: config,
	}
	r.Router.Use(cors.Default())
	r.baseConfigureRouter()
	return r
}

func (r *router) baseConfigureRouter() {
	r.Router.LoadHTMLGlob("internal/templates/*")
	r.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Router.POST("/order", func(ctx *gin.Context) {
		uuid := ctx.PostForm("uuid")

		order, err := r.Config.OrderService.GetOrder(uuid)
		if err != nil {
			log.Println(err)
		}
		// b, err := json.Marshal(order)
		// if err != nil {
		// 	log.Println(err)
		// }
		ctx.JSON(http.StatusOK, gin.H{"order": order})
	})
	r.Router.GET("/order", GetOrderTemplate)
}
