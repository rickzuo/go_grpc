package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
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
func requestMain(url string,wg *sync.WaitGroup,cancel chan bool,limit chan bool) {
	defer wg.Done()

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
	<- limit
	for  {
		select {
			case <- cancel:
				fmt.Println(url+":done")
				return

		}
	}

}

func main() {
	var wg sync.WaitGroup
	var cancel = make(chan bool)
	var limit = make(chan bool,30)
	t1 := time.Now()
	requestApi(startApiUrl)
	for _,rd := range resultData{
		wg.Add(1)
		limit <- true
		url := fmt.Sprintf("%s/%s/%s",baseUrl,rd.Slugname,rd.Lang)
		fmt.Println("start parse title:",rd.Title)
		go requestMain(url,&wg,cancel,limit)
	}

	close(cancel)
	wg.Wait()
	fmt.Println("cost time:",time.Since(t1))

}
