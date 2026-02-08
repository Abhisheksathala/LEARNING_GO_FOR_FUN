package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func testhandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "only post is allowed man ",
		})
		return
	}

	defer r.Body.Close()

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

	http.HandleFunc("/test", testhandler)

	err := http.ListenAndServe(":5000", nil)

	fmt.Println(err)

}
