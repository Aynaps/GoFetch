package main

import (
	"net/http"
	"src/main/src/receiptService"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.GET("/:id/points", receiptService.HandleGetRewardPoints)
	r.POST("/receipt", receiptService.HandleProcessReceipt)

	r.Run()
}
