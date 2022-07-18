package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetOrderTemplate(ctx *gin.Context) {
	// Рендеринг шаблона в соответствии с полным именем файла и передайте параметры
	ctx.HTML(200, "order.html", gin.H{
		"title": "Вы два большого дурака",
	})
}

func ValidateOrder(ctx *gin.Context) {
	emailValue := ctx.PostForm("email")
	passwordValue := ctx.PostForm("password")
	fmt.Println(emailValue, passwordValue)
}
