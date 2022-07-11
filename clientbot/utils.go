package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func StrParser(request string) *linebot.TextMessage {
	words := strings.Fields(request)
	response := "請輸入'看魚'、'看海鮮'或'看火鍋料'哦"
	switch words[0] {
	case "看魚":
		response = getallfish()
	case "看海鮮":
		response = "蛤蠣一斤120"
	case "看火鍋料":
		response = "火鍋料一斤250"
	}
	return linebot.NewTextMessage(response)
}

func getallfish() string {
	retmsg := ""
	resp, err := http.Get("http://localhost:8080/fish")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var fishmpslice []map[string]interface{}
	if err := json.Unmarshal(body, &fishmpslice); err != nil {
		log.Fatalln(err)
	}
	for _, fishmp := range fishmpslice {
		retmsg += fishmp["name"].(string) + "一" + fishmp["unit"].(string) + fishmp["price"].(string) + "元\n"
	}
	return retmsg
}
