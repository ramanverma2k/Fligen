package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ramanverma2k/Fligen/routes"
	log "github.com/sirupsen/logrus"
)

type AppConfig struct {
	ginClient  *gin.Engine
	httpClient *http.Client
}

func main() {
	err := godotenv.Load()
	if err != nil {
		// Log an error incase .env file is not found
		log.Error("Error loading .env file")
	}

	val, found := os.LookupEnv("OPENAI_SECRET_KEY")

	if !found || val == "" {
		// Exit the program if the key is not found in
		// environmental path as well.
		//
		// Cannot make any requests to the OpenAI API if
		// we do not have the secret key.
		log.Fatal("DALL-E OpenAI API Secret Key not found.")
	}

	app := &AppConfig{
		ginClient:  ginClient(),
		httpClient: httpClient(),
	}

	routes.Setup(app.httpClient, app.ginClient)

	app.ginClient.Run()
}
