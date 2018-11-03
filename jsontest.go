package main

import "fmt"
import "encoding/json"

//import "reflect"

var jsonStr = `{
	"age":111,
    "agg_source": {
        "filter": {
            "bool": {
                "must": [
                    {
                        "term": {
                            "status": 1
                        }
                    }
                ],
                "must_not": [
                    {
                        "term": {
                            "borough_id": 9999999
                        }
                    }
                ]
            }
        }
        
    },
    "aggs": {
            "aggs_borough": "aaaa"
     }
}
`

type Inter struct {
	Age        int         `json:"age"`
	Agg        Agg         `json:"aggs"`
	Agg_source interface{} `json:"agg_source"`
}

type Agg struct {
	A string `json:"aggs_borough"`
}

func main() {
	var inter Inter
	err := json.Unmarshal([]byte(jsonStr), &inter)
	fmt.Println(err)
	fmt.Println(inter)

	b, err := json.Marshal(inter)
	if err != nil {
		fmt.Println()
	}
	fmt.Println(string(b))
	// var f interface{}
	// b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	// json.Unmarshal(b, &f)
	// fmt.Println(f)
}
