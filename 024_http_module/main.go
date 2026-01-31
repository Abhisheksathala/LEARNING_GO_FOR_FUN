package main

import (
	"fmt"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "only get is allowed", http.StatusMethodNotAllowed)
	}
	_, _ = res.Write([]byte("hello from GO net/http server"))
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("try going to 4000 port ")
	err := http.ListenAndServe(":5000", nil)
	fmt.Println(err)
}
