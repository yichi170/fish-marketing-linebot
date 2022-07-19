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
	response := "請按照以下格式\n「商品名 單位 (斤/盤) 價錢」\n並以空白隔開\n例：鮭魚 斤 750\n\n"
	response += "或輸入「刪除 商品名」\n刪除商品資訊"
	if len(words) < 3 {
		return linebot.NewTextMessage(response)
	}
	if _, err := strconv.Atoi(words[2]); err != nil {
		response = "價錢需為數字"
	} else if words[0] != "刪除" {
		response = postfish(words)
	} else {
		response = deletefish(words[1])
	}
	return linebot.NewTextMessage(response)
}

func postfish(words []string) string {
	address := os.Getenv("API_ADDRESS") + "/fish"

	data := url.Values{
		"name":  {words[0]},
		"unit":  {words[1]},
		"price": {words[2]},
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

func deletefish(fish string) string {
	address := os.Getenv("API_ADDRESS") + "/fish/delete"

	data := url.Values{
		"name": {fish},
	}
	retmsg := ""
	resp, err := http.PostForm(address, data)
	if err != nil {
		retmsg = "刪除資料時發生問題"
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
