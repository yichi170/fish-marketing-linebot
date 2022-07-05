package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.POST("/callback", func(c *gin.Context) {
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				c.Writer.WriteHeader(400)
			} else {
				c.Writer.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					response := linebot.NewTextMessage(message.Text)
					if _, err = bot.ReplyMessage(event.ReplyToken, response).Do(); err != nil {
						log.Print(err)
					}
				default:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Hello World")).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})

	router.Run()
}
