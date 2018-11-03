package main

import (
	"fmt"
	"github.com/json-iterator/go"
)

type ESStruct struct {
	Filter map[string]interface{} `json:"filter"`
	From   interface{}            `json:"from"`
	Size   interface{}            `json:"size"`
	Sort   interface{}            `json:"sort"`
	Data   interface{}            `json:"data"`
	Page   interface{}            `json:"page"`
}

func main() {
	b := []byte(`{
  "filter": {
      "id": "1,2,3,4,5",
    "status":"1"
  },
  "from": 0,
  "size": 1,
  "type": "string"
}`)
	var es ESStruct
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(b, &es)
	fmt.Println(es)
	fmt.Print(es.Filter["id"])
}
