package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrderTemplate(ctx *gin.Context) {
	// Рендеринг шаблона в соответствии с полным именем файла и передайте параметры
	ctx.HTML(200, "order.html", gin.H{
		"title": "Hey Google",
	})
}

// func ValidateOrder(ctx *gin.Context) string {
// emailValue := ctx.PostForm("email")
// passwordValue := ctx.PostForm("password")
// fmt.Println(emailValue, passwordValue)
// 	fmt.Println("Yes sir")
// 	return "Yest"
// }

func ValidateOrder(ctx *gin.Context) {
	uuid := ctx.PostForm("uuid")
	fmt.Println(uuid)
	ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func postAlbums(c *gin.Context) {
	var newAlbum []string

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	newAlbum = append(newAlbum, "newAlbum")
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
