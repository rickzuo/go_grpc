package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)


type ResultData struct {
	Feature  bool   `json:"feature,omitempty"`
	Slugname string `json:"slugname"`
	Lang     string `json:"lang"`
	Title    string `json:"title"`
}
var resultData = make([]*ResultData,0)
var baseUrl string = "https://www.arealme.com"
var startUrl string = baseUrl + "/roc-famous-woman-quiz/cn/"
var startApiUrl string = baseUrl + "/mental/cn/promo"
var hrefList = make([]string, 0)

func requestApi(url string){
	fmt.Println("url:",url)
	client := &http.Client{}
	method := "GET"
	req, err := http.NewRequest(method, url, bytes.NewBuffer(nil))
	if err != nil {
		errMsg := fmt.Sprintf("[%s] url [%s] new request fail,%s", method, url, err)
		fmt.Println(errMsg)
	}

	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("[%s] url [%s] do request fail,%s", method, url, err)
		fmt.Println(errMsg)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errMsg := fmt.Sprintf("[%s] url [%s] read all fail,%s", method, url, err)
		fmt.Println(errMsg)

	}

	err = json.Unmarshal(body, &resultData)

}
func requestMain(url string) {
	client := &http.Client{}
	method := "GET"
	req, err := http.NewRequest(method, url, bytes.NewBuffer(nil))
	if err != nil {
		errMsg := fmt.Sprintf("[%s] url [%s] new request fail,%s", method, url, err)
		fmt.Println(errMsg)
	}
	// default json
	//req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("[%s] url [%s] do request fail,%s", method, url, err)
		fmt.Println(errMsg)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errMsg := fmt.Sprintf("[%s] url [%s] read all fail,%s", method, url, err)
		fmt.Println(errMsg)

	}
	//fmt.Println(fmt.Sprintf("body:%s", string(body)))
	titleRegex, _ := regexp.Compile(`<title>(.*?)</title>`)
	title := string(titleRegex.Find(body))

	fmt.Println("title:", title)

	//hrefRegex, _ := regexp.Compile(`class="feature" href="(.*?)" `)
	hrefRegex, _ := regexp.Compile(`class="feature" href="(.*?)" `)
	hrefRegList := hrefRegex.FindAllStringSubmatch(string(body), -1)
	for _, href := range hrefRegList {
		hrefList = append(hrefList, baseUrl+href[1])
	}

}

func main() {
	//var wg sync.WaitGroup
	t1 := time.Now().Second()
	requestApi(startApiUrl)
	for _,rd := range resultData[:10]{
		url := fmt.Sprintf("%s/%s/%s",baseUrl,rd.Slugname,rd.Lang)
		fmt.Println("start parse title:",rd.Title)
		requestMain(url)
	}

	fmt.Println("cost time:",time.Now().Second()-t1)

	//for i := 0; i < len(resultData); i++ {
	//	url := hrefList[i]
	//	fmt.Println("request url：", url)
	//	requestMain(url)
	//}
	//fmt.Println("hrefList:", hrefList)
	//wg.Done()
}
