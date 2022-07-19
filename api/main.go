package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/api/option"
)

func connect() (*firestore.Client, error) {
	ctx := context.Background()
	// sak := os.Getenv("SERVICE_ACCOUNT_KEY")
	sak := os.Getenv("SERVICE_ACCOUNT_KEY_JSON")
	config := &firebase.Config{ProjectID: "fish63-485"}
	// opt := option.WithCredentialsFile(sak)
	opt := option.WithCredentialsJSON([]byte(sak))
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return client, err
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/fish", getallfish)
	r.POST("/fish", postfish)

	return r
}
