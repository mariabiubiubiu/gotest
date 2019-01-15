package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)



func Testhttp()  {
	Client := &http.Client{}
	uri := "http://182.92.96.120:9741/" + "online_bj_ask" + "/" + "bj_ask" + "/" + "_search?pretty"
	dsl := `{"_source":[],"from":0,"query":{"bool":{"must":[{"nested":{"path":"house_info","query":{"bool":{"must":[{"terms":{"house_info.source":["1","2","3","4"]}}],"must_not":null,"should":null}}}},{"nested":{"path":"house_info","query":{"bool":{"must":[{"term":{"house_info.status":"1"}}],"must_not":null,"should":null}}}}],"must_not":null,"should":[{"range":{"min_price":{"gte":"100","lte":"200"}}},{"term":{"house_room":"2"}}]}},"size":1,"sort":[{"house_totalarea":{"order":"desc"}}]}`
	req, _ := http.NewRequest("POST", uri, strings.NewReader(dsl))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	//fmt.Println(r.ESClient().Get("www"))
	start := time.Now()
	a := time.Now().UnixNano()
	resp, err := Client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	b := time.Now().UnixNano()
	fmt.Println("request1", b-a)
	//fmt.Println(resp, err)
	defer resp.Body.Close()
	finish := time.Now()
	buf := make([]byte, 1024)
	res, _ := ioutil.ReadAll(resp.Body)
	c := time.Now().UnixNano()
	fmt.Println("requestIO", c - b)
	resp.Body.Read(buf)
	//res := make([]byte, 1024000)
	//resp.Body.Read(res)
	var rawdata Rawdata
	errs := json.Unmarshal(res, &rawdata)
	fmt.Println(errs)
	d := time.Now().UnixNano()
	fmt.Println("unmarshall", d-c)
	resMap := map[string]interface{}{}
	resList := make([]interface{}, 0)
	for _, v := range rawdata.Hits.Hits {
		resList = append(resList, v.Source)
	}
	//resMap["data"] = map[string]interface{}{"total": rawdata.Hits.Total, "list": resList}
	resMap["data"] = resList
	resMap["page"] = map[string]interface{}{"total": rawdata.Hits.Total}
	resMap["code"] = resp.StatusCode
	resMap["message"] = "success"
	resMap["time"] = finish.Sub(start)
	fmt.Println(resMap)
}

type Rawdata struct {
	Took     int            `json:"took"`
	TimedOut bool           `json:"timed_out"`
	Shards   map[string]int `json:"_shards"`
	Hits     ESdata         `json:"hits"`
}

type ESdata struct {
	Hits     []Hitdata `json:"hits"`
	Total    int       `json:"total"`
	MaxScore float64   `json:"max_score"`
}
type Hitdata struct {
	Index  string                 `json:"_index"`
	Type   string                 `json:"_type"`
	Id     string                 `json:"_id"`
	Score  float64                `json:"_score"`
	Source map[string]interface{} `json:"_source"`
}

func Testfast()  {
	Client := &fasthttp.Client{}
	uri := "http://182.92.96.120:9741/" + "online_bj_ask" + "/" + "bj_ask" + "/" + "_search?pretty"
	dsl := `{"_source":[],"from":0,"query":{"bool":{"must":[{"nested":{"path":"house_info","query":{"bool":{"must":[{"terms":{"house_info.source":["1","2","3","4"]}}],"must_not":null,"should":null}}}},{"nested":{"path":"house_info","query":{"bool":{"must":[{"term":{"house_info.status":"1"}}],"must_not":null,"should":null}}}}],"must_not":null,"should":[{"range":{"min_price":{"gte":"100","lte":"200"}}},{"term":{"house_room":"2"}}]}},"size":1,"sort":[{"house_totalarea":{"order":"desc"}}]}`
	req := &fasthttp.Request{}
	req.SetRequestURI(uri)
	req.Header.SetMethod("POST")
	req.SetBodyString(dsl)
	//req, _ := http.NewRequest("POST", uri, strings.NewReader(dsl))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	//fmt.Println(r.ESClient().Get("www"))
	start := time.Now()
	a := time.Now().UnixNano()
	resp := &fasthttp.Response{}
	err := Client.Do(req, resp)
	if err != nil {
		fmt.Println(err)
	}
	b := time.Now().UnixNano()
	fmt.Println("request1", b-a)
	//fmt.Println(resp, err)
	finish := time.Now()
	//buf := make([]byte, 1024)
	//res, _ := ioutil.ReadAll(resp.Body)
	c := time.Now().UnixNano()
	fmt.Println("requestIO", c - b)
	//resp.Body.Read(buf)
	//res := make([]byte, 1024000)
	//resp.Body.Read(res)
	var rawdata Rawdata
	errs := json.Unmarshal(resp.Body(), &rawdata)
	fmt.Println(errs)
	d := time.Now().UnixNano()
	fmt.Println("unmarshall", d-c)
	resMap := map[string]interface{}{}
	resList := make([]interface{}, 0)
	for _, v := range rawdata.Hits.Hits {
		resList = append(resList, v.Source)
	}
	//resMap["data"] = map[string]interface{}{"total": rawdata.Hits.Total, "list": resList}
	resMap["data"] = resList
	resMap["page"] = map[string]interface{}{"total": rawdata.Hits.Total}
	resMap["code"] = resp.StatusCode
	resMap["message"] = "success"
	resMap["time"] = finish.Sub(start)
	fmt.Println(resMap)
}


func main() {
	Testhttp()
	Testfast()
}