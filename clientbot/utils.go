package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func StrParser(request string) *linebot.TextMessage {
	words := strings.Fields(request)
	response := "請輸入'看魚價'來看今日魚價哦"
	switch words[0] {
	case "看魚價":
		response = getallfish()
	}
	return linebot.NewTextMessage(response)
}

func getallfish() string {
	retmsg := ""
	address := os.Getenv("API_ADDRESS") + "/fish"
	resp, err := http.Get(address)
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
	for i, fishmp := range fishmpslice {
		retmsg += fishmp["name"].(string) + "一" + fishmp["unit"].(string) + fishmp["price"].(string) + "元"
		if i < len(fishmpslice)-1 {
			retmsg += "\n"
		}
	}
	return retmsg
}
