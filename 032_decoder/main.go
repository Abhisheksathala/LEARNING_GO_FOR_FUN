package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func writeJSON(w http.ResponseWriter, r *http.Request) {

}

func testhandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	res := map[string]any{
		"ok":          true,
		"message":     "JSON encode successfully",
		"dateandtime": time.Now().UTC(),
	}

	_ = json.NewEncoder().Encode()

}

func main() {

	http.HandleFunc("/test", testhandler)

	err := http.ListenAndServe(":5000", nil)

	fmt.Println(err)

}
