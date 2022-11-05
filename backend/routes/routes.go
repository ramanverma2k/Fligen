package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(httpClient *http.Client, gClient *gin.Engine) {
	gClient.GET("api/dall-e/generateImage", generateImageRoute(httpClient))
}

func generateImageRoute(httpClient *http.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"responseCode": 200,
			"status":       "ok",
		})
	}
}
