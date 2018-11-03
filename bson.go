package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	//args := make(map[string]interface{})
	//args["a"] = "b"
	//out, _ := bson.Marshal(args)
	//fmt.Println(string(out))

	b := []byte(`{"aaa":"bbb"}`)
	var parse map[string]interface{}
	json.Unmarshal(b, &parse)

	bs, _ := bson.Marshal(parse)
	var M bson.M
	bson.Unmarshal(bs, &M)

	fmt.Println(M)
	test(parse)
}

func test(b map[string]interface{}) {
	fmt.Println(b["aaa"])
}
