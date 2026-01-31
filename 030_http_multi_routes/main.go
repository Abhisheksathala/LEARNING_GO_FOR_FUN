package main

import (
	"fmt"
	"net/http"
)

func roothandler(res http.ResponseWriter, req *http.Request) {

	_, _ = res.Write([]byte("wellcome try to /hello?name= sangam "))

}

func hellohandler(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	if name == "" {
		name = "gust"
	}

	_, _ = res.Write([]byte(name))

}

func main() {

	http.HandleFunc("/", roothandler)
	http.HandleFunc("/hello", hellohandler)

	err := http.ListenAndServe(":3000", nil)

	fmt.Println(err)
}
