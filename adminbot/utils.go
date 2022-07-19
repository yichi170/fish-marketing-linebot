package main

import (
	"encoding/json"
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
	response += "或輸入「刪除 商品名」\n刪除商品資訊\n\n"
	response += "或輸入「查看」\n查看全部商品\n\n"
	if words[0] == "刪除" {
		response = deletefish(words[1])
	} else if words[0] == "查看" {
		response = getallfish()
	}

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
