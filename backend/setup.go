package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}

	return client
}

func ginClient() *gin.Engine {
	gClient := gin.Default()

	return gClient
}
