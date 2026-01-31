package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func successhandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	res := map[string]any{
		"ok":          true,
		"message":     "JSON encode successfully",
		"dateandtime": time.Now().UTC(),
	}

	_ = json.NewEncoder(w).Encode(res)

}

func main() {

	http.HandleFunc("/ok", successhandler)

	err := http.ListenAndServe(":5000", nil)

	fmt.Println(err)

}
