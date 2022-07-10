package main

import (
	"strings"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func StrParser(request string) *linebot.TextMessage {
	words := strings.Fields(request)
	response := "請輸入'看魚'、'看海鮮'或'看火鍋料'哦"
	switch words[0] {
	case "看魚":
		response = "白鯧一斤1200"
	case "看海鮮":
		response = "蛤蠣一斤120"
	case "看火鍋料":
		response = "火鍋料一斤250"
	}
	return linebot.NewTextMessage(response)
}
