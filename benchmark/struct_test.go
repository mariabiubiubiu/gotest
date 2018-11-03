package benchmark

import (
	"encoding/json"
	"testing"
)

type ES struct {
	Filter interface{} `json:"filter"`
	From   interface{} `json:"from"`
	Size   interface{} `json:"size"`
	Sort   interface{} `json:"sort"`
	Data   interface{} `json:"data"`
	Page   interface{} `json:"page"`
}

type ESint struct {
	Filter interface{} `json:"filter"`
	From   interface{} `json:"from"`
	Size   int         `json:"size"`
	Sort   interface{} `json:"sort"`
	Data   interface{} `json:"data"`
	Page   interface{} `json:"page"`
}

func Benchmarkforint(b testing.B) {
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
  "size":10,
}`)
	var es ESint
	for i := 1; i < b.N; i++ {
		json.Unmarshal(buf, &es)
	}
}

func Benchmarkforinterface(b testing.B) {
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
  "size":10,
}`)
	var es ES
	for i := 1; i < b.N; i++ {
		json.Unmarshal(buf, &es)
	}
}
