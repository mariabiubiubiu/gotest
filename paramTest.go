package main

import (
	"fmt"
	"github.com/json-iterator/go"
	"strings"
)

type ES struct {
	Filter interface{} `json:"filter"`
	From   interface{} `json:"from"`
	Size   interface{} `json:"size"`
	Sort   interface{} `json:"sort"`
	Data   interface{} `json:"data"`
	Page   interface{} `json:"page"`
}

func main() {
	var param ES
	buf := []byte(`{
  "filter": {
    "status": 1,
	"content":"sdniu",
	"not-uid":323
"is_check":"sdniu",
"create_time":"1-2",
"content":"sdndcdaiu"
  }
  "from":0,
  "size":10			
}`)
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(buf, &param)
	parseMap := param.Filter.(map[string]interface{})
	terms := []string{"uid", "qid", "status", "is_check", "not-uid"}
	likes := []string{"title", "class", "type", "content"}
	ranges := []string{"create_time", "answer_num"}
	ret := new(Bool)
	ret = keyInSlice(parseMap, ranges, "range", ret)
	ret = keyInSlice(parseMap, terms, "terms", ret)
	ret = keyInSlice(parseMap, likes, "likes", ret)

	p := make(map[string]interface{})
	p["bool"] = ret
	a, _ := json.Marshal(p)
	fmt.Println(string(a))
}

func keyInSlice(destMap map[string]interface{}, slices []string, name string, ret *Bool) *Bool {
	for _, value := range slices {
		if mapValue, ok := destMap[value]; ok {
			if strings.HasPrefix(value, "not-") {
				array := strings.Split(value, "-")
				term := make(map[string]interface{})
				if name != "range" {
					term[array[1]] = mapValue
				} else {
					temp := make(map[string]interface{})
					ranges := strings.Split(mapValue.(string), "-")
					temp["gt"] = ranges[0]
					temp["lt"] = ranges[1]
					term[array[1]] = temp
				}
				ret.MustNot = append(ret.MustNot, map[string]map[string]interface{}{name: term})
			} else {
				term := make(map[string]interface{})
				if name != "range" {
					term[value] = mapValue
				} else {
					temp := make(map[string]interface{})
					ranges := strings.Split(mapValue.(string), "-")
					temp["gt"] = ranges[0]
					temp["lt"] = ranges[1]
					term[value] = temp
				}
				ret.Must = append(ret.Must, map[string]map[string]interface{}{name: term})
			}
		}
	}
	return ret
}

type Bool struct {
	Must    []map[string]map[string]interface{} `json:"must"`
	MustNot []map[string]map[string]interface{} `json:"must_not"`
}
