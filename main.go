package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	generalUrl = "http://api.fanyi.baidu.com/api/trans/vip/translate"
	vipUrl     = "https://fanyi-api.baidu.com/api/trans/vip/fieldtranslate"
)

var (
	appID, appSecret, appUrl, appDomain string
)

type resp struct {
	From        string `json:"from"`
	To          string `json:"to"`
	TransResult []struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	} `json:"trans_result"`
}

type config struct {
	AppId  string `json:"app_id"`
	Secret string `json:"secret"`
	Type   int    `json:"type"`
	Domain string `json:"domain"`
}

func init() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var cfg config
	if err = json.Unmarshal(file, &cfg); err != nil {
		panic(err)
	}
	appID, appSecret, appDomain = cfg.AppId, cfg.Secret, cfg.Domain
	if cfg.Type == 0 {
		appUrl = generalUrl
	} else {
		appUrl = vipUrl
	}
}

func main() {
	salt := strconv.Itoa(int(time.Now().Unix()))
	q := "apple"
	visitUrl := fmt.Sprintf("%s?q=%s&from=en&to=zh&appid=%s&salt=%s&sign=%s&domain=%s",
		appUrl, q, appID, salt, sign(appID, salt, appSecret, q, appDomain), appDomain)
	request, err := http.Get(visitUrl)
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	var r resp
	if err = json.Unmarshal(body, &r); err != nil {
		panic(err)
	}
	fmt.Println(r.TransResult[0].Dst)
}

func sign(appID, salt, secret, q, domain string) string {
	s := appID + q + salt + domain + secret
	fmt.Println(s)
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}
