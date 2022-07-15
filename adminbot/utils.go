package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	_ "github.com/joho/godotenv/autoload"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func StrParser(request string) *linebot.TextMessage {
	words := strings.Fields(request)
	response := "請輸入「商品名 價錢 單位」並以空白隔開"
	if len(words) < 3 {
		return linebot.NewTextMessage(response)

	}
	if _, err := strconv.Atoi(words[2]); err != nil {
		response = "價錢需為數字"
	} else {
		response = postfish(words)
	}
	return linebot.NewTextMessage(response)
}

func postfish(words []string) string {
	address := os.Getenv("API_ADDRESS") + "/fish"

	data := url.Values{
		"name":  {words[0]},
		"price": {words[1]},
		"unit":  {words[2]},
	}

	retmsg := ""
	resp, err := http.PostForm(address, data)
	if err != nil {
		retmsg = "更新資料時發生問題"
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	tmp, err := ioutil.ReadAll(resp.Body)
	retmsg = string(tmp)

	if err != nil {
		retmsg = "更新資料時發生問題"
		log.Fatalln(err)
	}
	return retmsg
}
