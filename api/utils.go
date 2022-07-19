package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/api/iterator"
)

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

func getallfish(c *gin.Context) {
	client, err := connect()
	ctx := context.Background()
	if err != nil {
		log.Fatalln("failed to connect Cloud Firestore @getallfish")
	}
	var mpslice []map[string]interface{}
	iter := client.Collection("fish").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln("iterate failed @getallfish")
		}
		mpslice = append(mpslice, doc.Data())
	}
	defer client.Close()
	c.IndentedJSON(http.StatusOK, mpslice)
}

func postfish(c *gin.Context) {
	client, err := connect()
	ctx := context.Background()
	if err != nil {
		log.Fatalln("failed to connect Cloud Firestore @POST")
	}

	name := c.PostForm("name")
	price := c.PostForm("price")
	unit := c.PostForm("unit")

	m := map[string]interface{}{
		"name":  name,
		"price": price,
		"unit":  unit,
	}

	ret := "更新「" + name + "」的價格為一" + unit + price + "元"

	update := false
	iter := client.Collection("fish").Where("name", "==", name).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		doc.Ref.Set(ctx, m)
		update = true
	}
	if update == false {
		_, _, err = client.Collection("fish").Add(ctx, m)
		if err != nil {
			ret = "更新" + name + "價格時發生問題"
			log.Fatalln(err)
		}
	}

	defer client.Close()
	c.String(http.StatusOK, ret)
}
