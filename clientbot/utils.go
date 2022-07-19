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

func StrParser(request string) *linebot.FlexMessage {
	words := strings.Fields(request)
	response := "請輸入'看魚價'來看今日魚價哦"
	var ret []byte
	switch words[0] {
	case "看魚價":
		response = getallfish(false)
		ret = []byte(getallfish(true))
	}
	container, err := linebot.UnmarshalFlexMessageJSON(ret)
	if err != nil {
		log.Fatalln(err)
	}
	return linebot.NewFlexMessage(response, container)
}

func getallfish(goodlook bool) string {
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

	if goodlook {
		retmsg = "{\"type\": \"carousel\", \"contents\": ["
	}

	for i, fishmp := range fishmpslice {
		if goodlook {
			retmsg += newflexmsg(fishmp)
			if i < len(fishmpslice)-1 {
				retmsg += ","
			}
		} else {
			retmsg += fishmp["name"].(string) + "一" + fishmp["unit"].(string) + fishmp["price"].(string) + "元"
			if i < len(fishmpslice)-1 {
				retmsg += "\n"
			}
		}
	}
	if goodlook {
		retmsg += "]}"
	}
	return retmsg
}

func newflexmsg(fishmp map[string]interface{}) string {
	jsonFile, err := os.Open("flex-msg.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	str := string(byteValue)
	str = strings.Replace(str, "name", fishmp["name"].(string), 1)
	str = strings.Replace(str, "price", fishmp["price"].(string), 1)
	str = strings.Replace(str, "unit", fishmp["unit"].(string), 1)

	return str
}
