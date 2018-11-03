package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func TestTemplate(w http.ResponseWriter, r *http.Request) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, _ := client.Get("http://www.qiushibaike.com/img/restful/img?id=9166")
	str, _ := ioutil.ReadAll(resp.Body)

	var jsonObj interface{}
	err := json.Unmarshal(str, &jsonObj)
	if err != nil {
		panic("NewMyJson err" + err.Error())
	}
	fmt.Println(jsonObj)
}
