package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("127.0.0.1:12345")
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		s := fmt.Sprintf("helloworld")
		fmt.Fprintf(w, "%v\n", s)
	})
	if err := http.ListenAndServe(":12345", nil); err != nil{
		fmt.Println("start")
	}
}