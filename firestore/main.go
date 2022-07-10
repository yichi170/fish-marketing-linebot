package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/api/option"
)

func connect() (*firestore.Client, error) {
	sak := os.Getenv("SEVICE_ACCOUNT_KEY")
	ctx := context.Background()
	opt := option.WithCredentialsFile(sak)
	app, err := firebase.NewApp(ctx, nil, opt)
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

func getfish(c *gin.Context) {
	client, err := connect()
	ctx := context.Background()
	if err != nil {
		log.Fatalln("failed to connect Cloud Firestore @GET")
	}
	fishname := c.Param("fishname")
	dsnap, err := client.Collection("fish").Doc(fishname).Get(ctx)

	if err != nil {
		log.Fatalln(err)
	}
	m := dsnap.Data()
	defer client.Close()
	c.IndentedJSON(http.StatusOK, m)
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/fish/:fishname", getfish)
	r.Run()
}
