package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server...")
	router := gin.Default()

	// ルーティングの設定
	router.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	// エラーが発生した場合にログを出力する
	log.Fatal(router.Run(":8000"))
}
